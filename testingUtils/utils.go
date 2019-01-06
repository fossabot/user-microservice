package testingUtils

import (
	"net/http"
	"net/http/httptest"
	"os"
)

func PrepareTest() {
	os.Setenv("TEST", "true") // allow to load correct config file
	os.Setenv("ENV", "")      // allow to set the env config file
}

// utility func who make the request
func PerformHTTPRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
