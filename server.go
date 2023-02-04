package main

import (
	"os"

	"github.com/labstack/echo/v4"
)

type server struct {
	url    string
	apiKey string
	e      *echo.Echo
}

func newServer() *server {
	s := &server{}

	s.e = echo.New()

	s.url = os.Getenv("URL")
	s.apiKey = os.Getenv("API_KEY")
	return s
}
