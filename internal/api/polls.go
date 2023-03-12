package api

import (
	"fmt"
	"github.com/bharat-rajani/go-polls/internal/api/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"runtime/debug"
)

var ServiceConfig *config.ServiceConfig

func StartService() {
	ServiceConfig = loadConfig()
	configureGlobalLogger()
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte(`HelloResponse`))
		if err != nil {
			return
		}
	})
	log.Info().Msgf("Starting server on %s:%d", ServiceConfig.Address, ServiceConfig.Port)
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", ServiceConfig.Address, ServiceConfig.Port), nil)
	if err != nil {
		panic(err)
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
