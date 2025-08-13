package weather

type WeatherResponse struct {
	City                    string `json:"city"`
	State                   string `json:"state"`
	Temperature             string `json:"temperature"`
	TemparatureUnit         string `json:"temperatureUnit"`
	ShortForecast           string `json:"shortForecast"`
	WeatherCharacterization string `json:"weatherCharacterization"`
}
type WeatherService interface {
	GetWeather(longitude, latitude string) (WeatherResponse, error)
}
