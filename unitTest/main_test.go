package unittest

import (
	"album/api"
	"album/config"
	"album/database"
	"album/database/repositories"
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	setUpMigrate()
	resutlCode := m.Run()
	downMigrate()

	os.Exit(resutlCode)
}

func setUpMigrate() {
	config.LoadEnv()

	repositories.DatabaseHandle()
	database.Migrate()
}

func downMigrate() {
	database.MigrateDown()
}

func getRoutes() *gin.Engine {
	return api.Routes()
}

func makeRequest(httpMethod, url string, body interface{}, isAuthRequest bool) (responseWriter *httptest.ResponseRecorder) {

	requestBody, _ := json.Marshal(body)
	request, err := http.NewRequest(
		httpMethod,
		url,
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		log.Fatal(errors.New(err.Error()))
	}

	responseWriter = httptest.NewRecorder()
	routes := getRoutes()
	routes.ServeHTTP(responseWriter, request)
	return
}
