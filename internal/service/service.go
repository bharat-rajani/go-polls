package service

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"
)

type Service struct {
	server *Server
	Log    *zerolog.Logger
	config *ServiceConfig
}

type Server struct {
	*http.Server
	mux *http.ServeMux
}

type ServiceConfig struct {
	Name   string
	Server struct {
		Address string
		Port    int
		Timeout struct {
			Read       int
			Write      int
			ReadHeader int
			Idle       int
		}
	}
	Log struct {
		BuildInfo bool
		Level     string
	}
	JWK struct {
	}
}

func (s *Service) GetConfig() *ServiceConfig {
	return s.config
}

func (s *Service) Run() {

	// OPTIONAL --- configuring shutdown hooks ---
	s.server.RegisterOnShutdown(func() {
		// initiate a shutdown process of another object but should not wait for it to complete
		go func() {

			// just a simulation
			s.Log.Info().Msg("Shutdown process triggered for foreign object")
		}()
	})

	s.Log.Info().Msgf("Starting server on %s:%d", s.GetConfig().Server.Address, s.GetConfig().Server.Port)

	go func() {
		err := s.server.ListenAndServe()
		if err != nil {
			s.Log.Fatal().Err(err).Send()
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	a := <-ch
	s.Log.Info().Str("signal received", a.String()).Send()
	s.shutdownServer()
}

func (s *Service) shutdownServer() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	s.Log.Info().Msg("shutting down the server")
	err := s.server.Shutdown(ctx)
	if err != nil {
		s.Log.Fatal().Err(err).Send()
	}
}

func NewServiceFromConfig(ctx context.Context, config *ServiceConfig) *Service {
	return &Service{
		server: NewHttpServerFromConfig(config),
		Log:    NewLogger(config),
		config: config,
	}
}

func NewHttpServerFromConfig(config *ServiceConfig) *Server {

	addr := fmt.Sprintf("%s:%d", config.Server.Address, config.Server.Port)
	mux := http.NewServeMux()
	return &Server{
		Server: &http.Server{
			Addr:              addr,
			Handler:           mux,
			ReadTimeout:       time.Duration(config.Server.Timeout.Read) * time.Second,
			ReadHeaderTimeout: time.Duration(config.Server.Timeout.ReadHeader) * time.Second,
			WriteTimeout:      time.Duration(config.Server.Timeout.Write) * time.Second,
			IdleTimeout:       time.Duration(config.Server.Timeout.Idle) * time.Minute,
			ConnState: func(conn net.Conn, state http.ConnState) {
				log.Info().Str("conn", conn.RemoteAddr().String()).Str("state", state.String()).Send()
			},
		},
		mux: mux,
	}
}

func NewLogger(config *ServiceConfig) *zerolog.Logger {
	logLevel, err := zerolog.ParseLevel(config.Log.Level)
	if err != nil {
		panic(err)
	}
	logObj := zerolog.New(os.Stdout).
		Level(logLevel).
		With().
		Timestamp().
		Caller().
		Logger()

	if config.Log.Level == "debug" {
		buildInfo, _ := debug.ReadBuildInfo()
		logObj = logObj.Output(zerolog.ConsoleWriter{Out: os.Stdout}).With().Int("pid", os.Getpid()).Str("go_version", buildInfo.GoVersion).Logger()
	}

	return &logObj
}

type Registrar func(service *Service) http.HandlerFunc

func (s *Service) RegisterRoute(pattern string, r Registrar) {
	s.server.mux.HandleFunc(pattern, r(s))
}
