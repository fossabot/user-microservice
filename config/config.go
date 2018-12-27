package config

import (
	"os"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Conf Configuration

type Configuration struct {
	Database DatabaseConfiguration
}

func LoadConfigFile() {

	env := os.Getenv("env")
	log.WithFields(logrus.Fields{"env": env}).Info("Loading config files for :")

	var configFileName string
	if env == "" {
		configFileName = "config.default"
	} else {
		configFileName = "config." + env
	}

	log.Info("Trying to load file : " + configFileName)

	viper.SetConfigName(configFileName)
	viper.AddConfigPath("config/")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&Conf)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
}
