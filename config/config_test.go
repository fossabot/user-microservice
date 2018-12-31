package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/spf13/viper"
)

func TestConfig(t *testing.T) {
	os.Setenv("test", "true")
	os.Setenv("env", "default")
	LoadConfigFile()
	testTableName := viper.GetString("database.tableName.user")

	os.Setenv("test", "")
	os.Setenv("env", "")
	LoadConfigFile()
	defaultTableName := viper.GetString("database.tableName.user")

	os.Setenv("test", "false")
	os.Setenv("env", "dev")
	LoadConfigFile()
	devTableName := viper.GetString("database.tableName.user")

	assert.NotEqual(t, testTableName, defaultTableName)
	assert.Equal(t, "User", defaultTableName)
	assert.Equal(t, "User", devTableName)
	assert.Equal(t, "User_test", testTableName)
}
