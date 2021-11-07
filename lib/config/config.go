package config

import (
	"github.com/spf13/viper"
	"log"
)

func SetConfigFile(name, path, extension string) {
	viper.SetConfigFile(name)
	viper.AddConfigPath(path)
	viper.SetConfigType(extension)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config file", err)
	}
}
