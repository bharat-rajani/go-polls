package api

import (
	"context"
	"fmt"
	"github.com/bharat-rajani/go-polls/internal/api/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"
)

var ServiceConfig *config.ServiceConfig

func StartService() {
	ServiceConfig = loadConfig()
	configureGlobalLogger()

	addr := fmt.Sprintf("%s:%d", ServiceConfig.Address, ServiceConfig.Port)
	mux := http.NewServeMux()
	mux = RegisterRootRoutes(mux)
	srv := http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadTimeout:       time.Duration(ServiceConfig.Timeout.Read) * time.Second,
		ReadHeaderTimeout: time.Duration(ServiceConfig.Timeout.ReadHeader) * time.Second,
		WriteTimeout:      time.Duration(ServiceConfig.Timeout.Write) * time.Second,
		IdleTimeout:       time.Duration(ServiceConfig.Timeout.Idle) * time.Minute,
		ConnState: func(conn net.Conn, state http.ConnState) {
			log.Info().Str("conn", conn.RemoteAddr().String()).Str("state", state.String()).Send()
		},
	}

	// OPTIONAL --- configuring shutdown hooks ---
	srv.RegisterOnShutdown(func() {
		// initiate a shutdown process of another object but should not wait for it to complete
		go func() {

			// just a simulation
			log.Info().Msg("Shutdown process triggered for foreign object")
		}()
	})

	log.Info().Msgf("Starting server on %s:%d", ServiceConfig.Address, ServiceConfig.Port)
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatal().Err(err).Send()
		}
	}()

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	a := <-ch
	log.Info().Str("signal received", a.String()).Send()
	shutdown(&srv)
}

func shutdown(srv *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	log.Info().Msg("shutting down the server")
	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatal().Err(err).Send()
	}
}

func configureGlobalLogger() {
	buildInfo, _ := debug.ReadBuildInfo()

	logObj := zerolog.New(os.Stdout).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Caller().
		Int("pid", os.Getpid()).
		Str("go_version", buildInfo.GoVersion).
		Logger()

	if ServiceConfig.Debug {
		logObj = logObj.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}
	log.Logger = logObj
}

func loadConfig() *config.ServiceConfig {
	viper.SetConfigFile("./config.yml")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	var serviceConfig config.ServiceConfig
	err = viper.UnmarshalKey("service", &serviceConfig)
	if err != nil {
		panic(err)
	}
	return &serviceConfig
}
