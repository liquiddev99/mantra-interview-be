package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	MlServerAddress string `mapstructure:"ML_SERVER_ADDRESS"`
	ServerSecretKey string `mapstructure:"SERVER_SECRET_KEY"`
	OriginAllowed   string `mapstructure:"ORIGIN_ALLOWED"`
	SymmetricKey    string `mapstructure:"SYMMETRIC_KEY"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
