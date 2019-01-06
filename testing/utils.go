package testing

import "os"

func prepareTest() {
	os.Setenv("TEST", "true") // allow to load correct config file
	os.Setenv("ENV", "")      // allow to set the env config file
}
