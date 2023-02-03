package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	QueryParamLatitude  = "lat"
	QueryParamLongitude = "lon"
	queryParamAPIKey    = "appid"
	url                 = "https://api.openweathermap.org/data/2.5/weather?"
)

type Coordinates struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type WeatherConditions struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type WeatherParameters struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
	SeaLevel  int     `json:"sea_level"`
	GrndLevel int     `json:"grnd_level"`
}

type WindParameters struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
	Gust  float64 `json:"gust"`
}

type RainVolume struct {
	H1 float64 `json:"1h"`
	H3 float64 `json:"3h"`
}

type Cloudiness struct {
	All int `json:"all"`
}

type SystemParameters struct {
	Type    int    `json:"type"`
	Id      int    `json:"id"`
	Country string `json:"country"`
	Sunrise int    `json:"sunrise"`
	Sunset  int    `json:"sunset"`
}

type ForecastData struct {
	Coord      Coordinates         `json:"coord"`
	Weather    []WeatherConditions `json:"weather"`
	Base       string              `json:"base"`
	Main       WeatherParameters   `json:"main"`
	Visibility int                 `json:"visibility"`
	Wind       WindParameters      `json:"wind"`
	Rain       RainVolume          `json:"rain"`
	Clouds     Cloudiness          `json:"clouds"`
	Dt         int                 `json:"dt"`
	Sys        SystemParameters    `json:"sys"`
	Timezone   int                 `json:"timezone"`
	Id         int                 `json:"id"`
	Name       string              `json:"name"`
	Cod        int                 `json:"cod"`
}

func GetWeatherForecastByLatLon(lat, lon float64, key string) (*ForecastData, error) {
	res, err := http.Get(
		fmt.Sprintf("%s%s=%.2f&%s=%.2f&%s=%s",
			url,
			QueryParamLongitude, lon,
			QueryParamLatitude, lat,
			queryParamAPIKey, key))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)

	forecast := &ForecastData{}
	err = decoder.Decode(forecast)
	if err != nil {
		return nil, err
	}
	if forecast.Cod != 200 {
		return nil, fmt.Errorf("external service error %d", forecast.Cod)
	}
	return forecast, nil
}
