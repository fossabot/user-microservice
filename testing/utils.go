package testing

import "os"

func prepareTest() {
	os.Setenv("test", "true")   // allow to load correct config file
	os.Setenv("env", "default") // allow to set the env config file
}
