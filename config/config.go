package config

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Conf Configuration

// Configuration is the struct who store configuration
type Configuration struct {
	Database DatabaseConfiguration
}

// LoadConfigFile load configuration from YAML file
func LoadConfigFile() {
	env := os.Getenv("env")
	log.Infof("Loading config files for : %s", env)

	var configFileName string
	if env == "" {
		configFileName = "config.default"
	} else {
		configFileName = "config." + env
	}
	log.Debugf("Trying to load file : %s", configFileName)
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
