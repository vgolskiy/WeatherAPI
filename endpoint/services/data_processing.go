package services

type ForecastBase struct {
	WeatherConditions []string `json:"weather_conditions"`
	FeelsLike         string   `json:"feels_like"`
}

type WeatherAlert struct {
	Name              string   `json:"name"`
	WeatherConditions []string `json:"weather_conditions"`
}

type Forecast struct {
	ForecastBase
	HaveAlerts   bool           `json:"have_alerts"`
	WeatherAlert []WeatherAlert `json:"weather_alert,omitempty"`
}

func ProcessWeatherForecastData(f *ForecastData) *Forecast {
	res := &Forecast{}

	for _, w := range f.Current.Weather {
		res.WeatherConditions = append(res.WeatherConditions, w.Main)
	}

	switch {
	case f.Current.FeelsLike > 303:
		res.FeelsLike = "Hot"
	case f.Current.FeelsLike > 268 && f.Current.FeelsLike <= 303:
		res.FeelsLike = "Moderate"
	case f.Current.FeelsLike <= 268:
		res.FeelsLike = "Cold"
	}

	if f.Alerts != nil {
		res.HaveAlerts = true
		for _, a := range f.Alerts {
			res.WeatherAlert = append(res.WeatherAlert, WeatherAlert{
				Name:              a.Event,
				WeatherConditions: a.Tags,
			})
		}
	}
	return res
}
