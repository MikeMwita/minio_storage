package util

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/viper"
)

type Config struct {
	HTTPServerAddress string `mapstructure:"HTTP_SERVER_ADDRESS"`
	GRPCServerAddress string `mapstructure:"GRPC_SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading .env file: %v", err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Error unmarshalling configuration: %v", err)
	}

	return config, nil

}
