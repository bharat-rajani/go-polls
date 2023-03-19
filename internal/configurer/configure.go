package configurer

import (
	"context"
	"github.com/bharat-rajani/go-polls/internal/service"
	"github.com/spf13/viper"
)

func StartAPIService(ctx context.Context) error {
	svcConfig, err := loadConfig()
	if err != nil {
		return err
	}
	svc := service.NewServiceFromConfig(ctx, svcConfig)
	RegisterRoutes(svc)
	svc.Run()
	return nil
}

func loadConfig() (*service.ServiceConfig, error) {
	viper.SetConfigFile("./config.yml")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		return nil, err
	}
	var serviceConfig *service.ServiceConfig
	err = viper.UnmarshalKey("service", &serviceConfig)
	if err != nil {
		return nil, err
	}
	return serviceConfig, nil
}
