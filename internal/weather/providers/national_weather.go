package providers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/anandxmj/weather-service/internal/weather"
)

const (
	NW_POINTS_URL = "http://api.weather.gov/points/%s,%s"
	NW_CELSIUS    = "C"
	NW_FAHRENHEIT = "F"
)

type NationalWeather struct{}

type PointsResponse struct {
	Properties struct {
		Forecast         string `json:"forecast"`
		RelativeLocation struct {
			Properties struct {
				City  string `json:"city"`
				State string `json:"state"`
			} `json:"properties"`
		} `json:"relativeLocation"`
	} `json:"properties"`
}

type ForecastResponse struct {
	Properties struct {
		Periods []struct {
			Temperature     float64 `json:"temperature"`
			TemperatureUnit string  `json:"temperatureUnit"`
			ShortForecast   string  `json:"shortForecast"`
		} `json:"periods"`
	} `json:"properties"`
}

func (nw *NationalWeather) GetWeather(latitude, longitude string) (weather.WeatherResponse, error) {

	if err := weather.ValidateCoordinates(longitude, latitude); err != nil {
		return weather.WeatherResponse{}, err
	}

	POINTS_URL := fmt.Sprintf(NW_POINTS_URL, latitude, longitude)
	log.Println("Fetching weather data from: ", POINTS_URL)
	resp, err := http.Get(POINTS_URL)

	if err != nil {
		log.Printf("Error fetching weather data: %v", err)
		return weather.WeatherResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return weather.WeatherResponse{}, fmt.Errorf("failed to get weather: %s", resp.Status)
	}

	var pointsResponse PointsResponse
	if err := json.NewDecoder(resp.Body).Decode(&pointsResponse); err != nil {
		return weather.WeatherResponse{}, err
	}

	forecast, err := GetForecast(pointsResponse.Properties.Forecast)
	if err != nil {
		return weather.WeatherResponse{}, err
	}

	weatherResponse := newWeatherResponse(pointsResponse, forecast)

	switch forecast.Unit {
	case NW_CELSIUS:
		weatherResponse.WeatherCharacterization = weather.CharacterizeTemperature(forecast.Temp, weather.METRIC)
	case NW_FAHRENHEIT:
		weatherResponse.WeatherCharacterization = weather.CharacterizeTemperature(forecast.Temp, weather.IMPERIAL)
	}

	return weatherResponse, nil
}

func newWeatherResponse(pointsResponse PointsResponse, forecast Forecast) weather.WeatherResponse {
	var weatherResponse weather.WeatherResponse
	weatherResponse.City = pointsResponse.Properties.RelativeLocation.Properties.City
	weatherResponse.State = pointsResponse.Properties.RelativeLocation.Properties.State
	weatherResponse.Temperature = fmt.Sprintf("%.1f", forecast.Temp)
	weatherResponse.TemparatureUnit = forecast.Unit
	weatherResponse.ShortForecast = forecast.ShortForecast
	return weatherResponse
}

type Forecast struct {
	Temp          float64
	Unit          string
	ShortForecast string
}

func GetForecast(forecastURL string) (Forecast, error) {
	resp, err := http.Get(forecastURL)
	if err != nil {
		return Forecast{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Forecast{}, fmt.Errorf("failed to get forecast: %s", resp.Status)
	}

	var forecastResponse ForecastResponse
	if err := json.NewDecoder(resp.Body).Decode(&forecastResponse); err != nil {
		return Forecast{}, err
	}

	if len(forecastResponse.Properties.Periods) == 0 {
		return Forecast{}, nil
	}

	period := forecastResponse.Properties.Periods[0]
	return Forecast{
		Temp:          period.Temperature,
		Unit:          period.TemperatureUnit,
		ShortForecast: period.ShortForecast,
	}, nil
}
