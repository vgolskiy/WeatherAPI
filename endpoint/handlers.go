package main

import (
	"net/http"

	"WeatherAPI/models"
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

		forecast, err := services.GetWeatherForecastByLatLon(latitude, longitude, s.url, s.apiKey)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		simplifiedForecast := models.ProcessWeatherForecastData(forecast)
		return c.JSON(http.StatusOK, simplifiedForecast)
	}
}
