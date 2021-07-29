package utils

import (
	"github.com/caarlos0/env"
	"github.com/spf13/viper"
)

// Config stores all configuration of the application
type Config struct {
	DBUri       string `mapstructure:"DB_URI" env:"DB_URI"`
	SecretKey   string `mapstructure:"SECRET_KEY" env:"SECRET_KEY"`
	RedisPass   string `mapstructure:"REDIS_PASS" env:"REDIS_PASS"`
	RedisHost   string `mapstructure:"REDIS_HOST" env:"REDIS_HOST"`
	RedisDb     int    `mapstructure:"REDIS_DB"   env:"REDIS_DB"`
	MaxFileSize int64  `mapstructure:"MAX_FILE_SIZE" env:"MAX_FILE_SIZE"`
}

func LoadConfig(path string, name string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		err = env.Parse(&config)
		if err != nil {
			return
		}
		if config.SecretKey != "" {
			if config.MaxFileSize != 0 {
				err = nil
				return
			}
		}
		return
	}
	err = viper.Unmarshal(&config)

	return
}
