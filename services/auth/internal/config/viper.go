package config

import (
	"log"

	"github.com/spf13/viper"
)

const ConfigPath = "./"

func InitConfig() Config {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(ConfigPath)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("fatal error config file: ", err)
	}

	var cfg Config

	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatal("fatal error unmarshal config: ", err)
	}

	return cfg
}
