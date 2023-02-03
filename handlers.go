package main

import (
	"fmt"
	"net/http"

	"WeatherAPI/services"
	"github.com/labstack/echo/v4"
)

func (s *server) handleGetWeatherForecast() echo.HandlerFunc {
	return func(c echo.Context) error {
		latitude, longitude, err := services.VerifyLatitudeLongitude(
			c.QueryParam(services.QueryParamLatitude),
			c.QueryParam(services.QueryParamLongitude),
		)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		forecast, err := services.GetWeatherForecastByLatLon(latitude, longitude, s.apiKey)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		fmt.Println(forecast)
		return c.NoContent(http.StatusNoContent)
	}
}
