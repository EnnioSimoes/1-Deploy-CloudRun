package configs

import (
	"github.com/spf13/viper"
)

type conf struct {
	Weatherapi_key string `mapstructure:"WEATHERAPI_KEY"`
}

func LoadConfig(path string) (*conf, error) {
	var config *conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
