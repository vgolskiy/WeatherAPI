package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

const (
	queryParamLatitude  = "lat"
	queryParamLongitude = "lon"
	queryParamAPIKey    = "appid"
	url                 = "https://api.openweathermap.org/data/2.5/weather?"
)

func (s *server) handleGetWeatherForecast() echo.HandlerFunc {
	return func(c echo.Context) error {
		latitude, err := strconv.ParseFloat(c.QueryParam(queryParamLatitude), 64)
		if err != nil {
			return err
		}
		longitude, err := strconv.ParseFloat(c.QueryParam(queryParamLongitude), 64)
		if err != nil {
			return err
		}

		res, err := http.Get(
			fmt.Sprintf("%s%s=%.2f&%s=%.2f&%s=%s",
				url,
				queryParamLongitude, longitude,
				queryParamLatitude, latitude,
				queryParamAPIKey, s.apiKey))
		if err != nil {
			return err
		}

		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		fmt.Println(string(body))
		return c.NoContent(http.StatusNoContent)
	}
}
