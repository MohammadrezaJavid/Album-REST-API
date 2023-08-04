package unittest

import (
	"net/http"
	"testing"

	"gotest.tools/assert"
)

func TestPostNewAlbum(t *testing.T) {
	for _, album := range Albums {
		responseWriter := makeRequest(
			http.MethodPost,
			"/api/jwt/albums",
			album,
			true,
		)
		assert.Equal(t, http.StatusCreated, responseWriter.Code)
	}
}

func TestGetAlbums(t *testing.T) {
	responseWriter := makeRequest(
		http.MethodGet,
		"/api/jwt/albums",
		nil,
		true,
	)
	assert.Equal(t, http.StatusOK, responseWriter.Code)
}

func TestGetAlbumByID(t *testing.T) {
	responseWriter := makeRequest(
		http.MethodGet,
		"/api/jwt/albums/"+string(Albums[0].ID),
		nil,
		true,
	)
	assert.Equal(t, http.StatusOK, responseWriter.Code)
}

func TestPutAlbum(t *testing.T) {
	responseWriter := makeRequest(
		http.MethodPut,
		"/api/jwt/albums",
		PutAlbum,
		true,
	)
	assert.Equal(t, http.StatusOK, responseWriter.Code)
}

func TestDeleteAlbumByID(t *testing.T) {
	responseWriter := makeRequest(
		http.MethodDelete,
		"/api/jwt/albums/"+string(Albums[0].ID),
		nil,
		true,
	)
	assert.Equal(t, http.StatusOK, responseWriter.Code)
}
