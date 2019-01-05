package config

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

type conf struct {
	GinMode string `yaml:"GIN_MODE"`
}

// TestConfigLocal is testing that we read the value from the config file (config.default.yaml)
func TestConfigLocal(t *testing.T) {
	expected := "debug"
	LoadConfigFile()
	got := viper.GetString("GIN_MODE")
	assert.Equal(t, expected, got)
}

// TestConfigRelease is testing we reading config from environnement variable in release
func TestConfigReleaseFromEnvVar(t *testing.T) {
	//setting the environnement variable "ENV" to DEV
	os.Setenv("ENV", "DEV")
	envVarName := "APP_PORT"
	expected := "8585"
	os.Setenv(envVarName, expected)
	LoadConfigFile()
	got := viper.GetString(envVarName)
	assert.Equal(t, expected, got)
	//clean var after test
	os.Setenv(envVarName, "")
}

// TestConfigRelease is testing we reading config from environnement variable in release
func TestConfigReleaseNoValueUsingDefault(t *testing.T) {
	//setting the environnement variable "ENV" to DEV
	os.Setenv("ENV", "DEV")
	envVarName := "APP_PORT"
	expected := "8080"
	LoadConfigFile()
	got := viper.GetString(envVarName)
	assert.Equal(t, expected, got)
	//clean var after test
	os.Setenv(envVarName, "")

}
