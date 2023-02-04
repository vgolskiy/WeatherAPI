package services

type ForecastBase struct {
	WeatherConditions []string `json:"weather_conditions"`
	FeelsLike         string   `json:"feels_like"`
}

type WeatherAlert struct {
	Name              string   `json:"name"`
	Description       string   `json:"description"`
	WeatherConditions []string `json:"weather_conditions"`
}

type Forecast struct {
	ForecastBase
	WeatherAlert []WeatherAlert `json:"weather_alert,omitempty"`
}

func ProcessWeatherForecastData(f *ForecastData) *Forecast {
	res := &Forecast{}

	for _, w := range f.Current.Weather {
		res.WeatherConditions = append(res.WeatherConditions, w.Description)
	}

	switch {
	case f.Current.FeelsLike > 303:
		res.FeelsLike = "Hot"
	case f.Current.FeelsLike > 293 && f.Current.FeelsLike <= 303:
		res.FeelsLike = "Warm"
	case f.Current.FeelsLike > 283 && f.Current.FeelsLike <= 293:
		res.FeelsLike = "Moderate"
	case f.Current.FeelsLike > 268 && f.Current.FeelsLike <= 283:
		res.FeelsLike = "Cold"
	case f.Current.FeelsLike <= 268:
		res.FeelsLike = "Freezing"
	}

	if f.Alerts != nil {
		for _, a := range f.Alerts {
			res.WeatherAlert = append(res.WeatherAlert, WeatherAlert{
				Name:              a.Event,
				Description:       a.Description,
				WeatherConditions: a.Tags,
			})
		}
	}
	return res
}
