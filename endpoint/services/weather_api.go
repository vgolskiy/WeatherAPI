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
)

type WeatherConditions struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Alert struct {
	SenderName  string   `json:"sender_name"`
	Event       string   `json:"event"`
	Start       int64    `json:"start"`
	End         int64    `json:"end"`
	Description string   `json:"description"`
	Tags        []string `json:"tags,omitempty"`
}

type WeatherParameters struct {
	Dt         int                 `json:"dt"`
	Sunrise    int                 `json:"sunrise"`
	Sunset     int                 `json:"sunset"`
	Temp       float64             `json:"temp"`
	FeelsLike  float64             `json:"feels_like"`
	Pressure   int                 `json:"pressure"`
	Humidity   int                 `json:"humidity"`
	DewPoint   float64             `json:"dew_point"`
	Uvi        float64             `json:"uvi"`
	Clouds     int                 `json:"clouds"`
	Visibility int                 `json:"visibility"`
	WindSpeed  float64             `json:"wind_speed"`
	WindDeg    int                 `json:"wind_deg"`
	WindGust   float64             `json:"wind_gust"`
	Weather    []WeatherConditions `json:"weather"`
}

type ForecastData struct {
	Lat            float64           `json:"lat"`
	Lon            float64           `json:"lon"`
	Timezone       string            `json:"timezone"`
	TimezoneOffset int64             `json:"timezone_offset"`
	Current        WeatherParameters `json:"current"`
	Alerts         []Alert           `json:"alerts,omitempty"`
}

func GetWeatherForecastByLatLon(lat, lon float64, url, key string) (*ForecastData, error) {
	res, err := http.Get(
		fmt.Sprintf("%s%s=%.2f&%s=%.2f&exclude=minutely,hourly,daily&%s=%s",
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
	if res.StatusCode < 200 || res.StatusCode > 299 {
		return nil, fmt.Errorf("external service error code %d", res.StatusCode)
	}
	return forecast, nil
}
