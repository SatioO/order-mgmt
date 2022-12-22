package config

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/satioO/order-mgmt/pkg/constants"
	"github.com/spf13/viper"
)

var configPath string

type Config struct {
	ServerName string `mapstructure:"serviceName"`
	Logger     Logger `mapstructure:"logger"`
	GRPC       GRPC   `mapstructure:"grpc"`
	AWS        AWS    `mapstructure:"aws"`
}

type AWS struct {
	Region string `mapstructure:"region"`
}

type GRPC struct {
	Port        string `mapstructure:"port"`
	Development bool   `mapstructure:"development"`
}

type Logger struct {
	LogLevel string `mapstructure:"level"`
	DevMode  bool   `mapstructure:"devMode"`
	Encoder  string `mapstructure:"encoder"`
}

func LoadConfig() (config *Config, err error) {
	if configPath == "" {
		configPathFromEnv := os.Getenv(constants.ConfigPath)
		if configPathFromEnv != "" {
			configPath = configPathFromEnv
		} else {
			getwd, err := os.Getwd()
			if err != nil {
				return nil, errors.Wrap(err, "os.Getwd")
			}
			configPath = fmt.Sprintf("%s/config/config.yaml", getwd)
		}
	}

	cfg := &Config{}

	viper.SetConfigType(constants.Yaml)
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "viper.ReadInConfig")
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, errors.Wrap(err, "viper.Unmarshel")
	}

	return cfg, nil
}
