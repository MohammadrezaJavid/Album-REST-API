package unittest

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"gotest.tools/assert"
)

func TestRegister(t *testing.T) {
	for _, user := range Users {
		responseWriter := makeRequest(
			http.MethodPost,
			"/api/user/register",
			user,
			false,
		)
		assert.Equal(t, http.StatusCreated, responseWriter.Code)
	}
}

func TestGetToken(t *testing.T) {
	for _, user := range Users {
		responseWriter := makeRequest(
			http.MethodPost,
			"/api/token",
			user,
			false,
		)

		assert.Equal(t, http.StatusOK, responseWriter.Code)
		var response map[string]string
		json.Unmarshal(responseWriter.Body.Bytes(), &response)

		responseToken := response["token"]

		if len(strings.TrimSpace(responseToken)) == 0 {
			t.Error("Token is nil!")
		}
	}
}
