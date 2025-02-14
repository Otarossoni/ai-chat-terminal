package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	OpenAiKey string `mapstructure:"OPENAI_API_KEY"`
}

var configuration = &Config{}

func Load() error {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&configuration); err != nil {
		return err
	}

	return nil
}

func Get() *Config {
	return configuration
}
