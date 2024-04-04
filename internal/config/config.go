package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	configName       = "SEARCH_API_AUTH_KEY"
	errConfigMissing = fmt.Errorf("missing config: %s", configName)
)

type Config struct {
	APIKey string `mapstructure:"SEARCH_API_AUTH_KEY"`
}

func NewConfig(path string) (*Config, error) {
	viper.SetConfigFile(fmt.Sprintf("%s/%s", path, ".env"))

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		return nil, errConfigMissing
	}

	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		return nil, errConfigMissing
	}

	return config, nil
}
