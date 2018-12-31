package config

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Conf is the configuration object of the app
var Conf Configuration
var projetPath = os.Getenv("GOPATH") + "/src/github.com/thomaspoignant/user-microservice"

// Configuration is the struct who store configuration
type Configuration struct {
	Database DatabaseConfiguration
}

// LoadConfigFile load configuration from YAML file
func LoadConfigFile() {
	configFileName := composeConfigFileName()
	log.Infof("Trying to load file : %s", configFileName)
	viper.SetConfigName(configFileName)
	viper.AddConfigPath(projetPath + "/config/")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&Conf)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
}

// determine the name of the config file
func composeConfigFileName() string {
	env := os.Getenv("env")
	test := os.Getenv("test")
	var configFileName []string
	configFileName = append(configFileName, "config")

	if strings.Compare("true", test) == 0 {
		configFileName = append(configFileName, "test")
	}

	if strings.Compare("", env) == 0 {
		configFileName = append(configFileName, "default")
	} else {
		configFileName = append(configFileName, env)
	}

	return strings.Join(configFileName, ".")
}
