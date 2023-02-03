package services

type ForecastBase struct {
	WeatherConditions []string `json:"weather_conditions"`
	FeelsLike         string   `json:"feels_like"`
}

type Forecast struct {
	Area string `json:"area"`
	ForecastBase
	WeatherAlert *ForecastBase `json:"weather_alert,omitempty"`
}

func ProcessWeatherForecastData(f *ForecastData) *Forecast {
	res := &Forecast{
		Area: f.Name,
	}

	for _, w := range f.Weather {
		res.WeatherConditions = append(res.WeatherConditions, w.Description)
	}

	switch {
	case f.Main.FeelsLike > 86:
		res.FeelsLike = "Hot"
	case f.Main.FeelsLike > 68 && f.Main.FeelsLike <= 86:
		res.FeelsLike = "Warm"
	case f.Main.FeelsLike > 50 && f.Main.FeelsLike <= 68:
		res.FeelsLike = "Moderate"
	case f.Main.FeelsLike > 23 && f.Main.FeelsLike <= 50:
		res.FeelsLike = "Cold"
	case f.Main.FeelsLike <= 23:
		res.FeelsLike = "Freezing"
	}
	return res
}
