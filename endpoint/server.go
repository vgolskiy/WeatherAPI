package main

import (
	"os"

	"github.com/labstack/echo/v4"
)

const URL = "https://api.openweathermap.org/data/3.0/onecall?"

type server struct {
	url    string
	apiKey string
	e      *echo.Echo
}

func newServer() *server {
	s := &server{}

	s.e = echo.New()

	s.url = URL
	s.apiKey = os.Getenv("API_KEY")
	return s
}
