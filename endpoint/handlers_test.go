package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type endpointTest struct {
	lat     string
	lon     string
	code    int
	message string
}

type errorMessage struct {
	Message string `json:"message"`
}

var tests = []endpointTest{
	{"63.46", "-170.37", http.StatusOK, ""},
}

var negTests = []endpointTest{
	{"100", "200", http.StatusBadRequest, "wrong coordinates value"},
	{"p", "100", http.StatusBadRequest, "wrong coordinates value"},
	{"", "", http.StatusBadRequest, "missing coordinates value"},
}

func TestHandleGetWeatherForecast(t *testing.T) {
	os.Setenv("URL", "https://api.openweathermap.org/data/3.0/onecall?")
	os.Setenv("API_KEY", "42eaa4ce5c82be16a896ba2cc17d6e27")
	defer os.Unsetenv("URL")
	defer os.Unsetenv("API_KEY")
	s := newServer()

	for _, test := range tests {
		req := httptest.NewRequest(
			http.MethodGet,
			fmt.Sprintf("/?lat=%s&lon=%s", test.lat, test.lon),
			nil)
		rec := httptest.NewRecorder()
		ctx := s.e.NewContext(req, rec)

		if assert.NoError(t, s.handleGetWeatherForecast()(ctx)) {
			assert.Equal(t, test.code, rec.Code)
		}
	}
}

func TestHandleGetWeatherForecastNegative(t *testing.T) {
	os.Setenv("URL", "https://api.openweathermap.org/data/3.0/onecall?")
	os.Setenv("API_KEY", "42eaa4ce5c82be16a896ba2cc17d6e27")
	defer os.Unsetenv("URL")
	defer os.Unsetenv("API_KEY")
	s := newServer()
	s.e.GET("/", s.handleGetWeatherForecast())

	for _, test := range negTests {
		req := httptest.NewRequest(
			http.MethodGet,
			fmt.Sprintf("/?lat=%s&lon=%s", test.lat, test.lon),
			nil)
		rec := httptest.NewRecorder()
		s.e.ServeHTTP(rec, req)

		assert.Equal(t, test.code, rec.Code)

		var m errorMessage
		json.Unmarshal(rec.Body.Bytes(), &m)
		assert.Equal(t, test.message, m.Message)
	}
}
