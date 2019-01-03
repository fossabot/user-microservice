package server

// Init the configuration of the APIs
func Init() {
	router := SetupRouter()
	router.Run()
}
