package services

import (
	"fmt"
	"strconv"
)

func VerifyLatitudeLongitude(lat, lon string) (latitude float64, longitude float64, err error) {
	if lat == "" || lon == "" {
		return 0, 0, fmt.Errorf("missing coordinates value")
	}
	latitude, err = strconv.ParseFloat(lat, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("wrong coordinates value")
	}
	longitude, err = strconv.ParseFloat(lon, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("wrong coordinates value")
	}
	if (latitude < -90 || latitude > 90) || (longitude < -180 || longitude > 180) {
		return 0, 0, fmt.Errorf("wrong coordinates value")
	}
	return
}
