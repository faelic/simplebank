package util

import (
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBSource            string        `mapstructure:"DB_SOURCE"`
	ServerAddress       string        `mapstructure:"SERVER_ADDRESS"`
	SecretKey           string        `mapstructure:"SECRET_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	viper.BindEnv("DB_SOURCE")
	viper.BindEnv("SERVER_ADDRESS")
	viper.BindEnv("SECRET_KEY")
	viper.BindEnv("ACCESS_TOKEN_DURATION")

	if err := viper.ReadInConfig(); err != nil {
		if !strings.Contains(err.Error(), "Config File") {
			return config, err
		}
	}

	err = viper.Unmarshal(&config)
	return
}
