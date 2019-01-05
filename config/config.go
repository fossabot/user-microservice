package config

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var ProjectPath = os.Getenv("GOPATH") + "/src/github.com/thomaspoignant/user-microservice"

// LoadConfigFile load configuration from YAML file
func LoadConfigFile() {
	if os.Getenv("ENV") == "" {
		configFileName := composeConfigFileName()
		log.Infof("Trying to load file : %s", configFileName)
		viper.SetConfigName(configFileName)
		viper.AddConfigPath(ProjectPath + "/config/")
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file, %s", err)
		}
	} else {
		viper.AutomaticEnv()
	}
	viper.SetDefault("APP_PORT", "8080")
	viper.SetDefault("GIN_MODE", gin.ReleaseMode)
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
