package config

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	Port string `mapstructure:"SERVICE_PORT" validate:"required"`

	DynamoDBTable string `mapstructure:"DB_TABLE_NAME" validate:"required"`

	AwsRegion string `mapstructure:"AWS_REGION" validate:"required"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	validate := validator.New()
	if err := validate.Struct(&config); err != nil {
		log.Fatalf("Missing required attributes %v\n", err)
	}

	return
}
