package main

import (
	"fmt"
	"net/http"
)

func main() {
	s := newServer()

	s.e.GET("/", s.handleGetWeatherForecast())

	port := 2000
	s.e.Logger.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), s.e))
}
