package errors

import (
	"fmt"
	"net/http"
)

type WeatherError struct {
	Code    int    `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

func (e *WeatherError) Error() string {
	return fmt.Sprintf("[%s] %s: %s", e.Type, e.Message, e.Details)
}

var (
	// Validation Errors
	ErrInvalidCoordinates = &WeatherError{
		Code:    http.StatusBadRequest,
		Type:    "INVALID_COORDINATES",
		Message: "Invalid latitude or longitude provided",
	}

	ErrMissingCoordinates = &WeatherError{
		Code:    http.StatusBadRequest,
		Type:    "MISSING_COORDINATES",
		Message: "Latitude and longitude are required",
	}

	ErrCoordinatesOutOfRange = &WeatherError{
		Code:    http.StatusBadRequest,
		Type:    "COORDINATES_OUT_OF_RANGE",
		Message: "Coordinates are outside valid range",
		Details: "Latitude must be between -90 and 90, longitude between -180 and 180",
	}

	// External Service Errors
	ErrWeatherServiceUnavailable = &WeatherError{
		Code:    http.StatusServiceUnavailable,
		Type:    "WEATHER_SERVICE_UNAVAILABLE",
		Message: "Weather service is temporarily unavailable",
	}

	ErrWeatherDataNotFound = &WeatherError{
		Code:    http.StatusNotFound,
		Type:    "WEATHER_DATA_NOT_FOUND",
		Message: "Weather data not available for the specified location",
	}

	// Internal Errors
	ErrInternalServer = &WeatherError{
		Code:    http.StatusInternalServerError,
		Type:    "INTERNAL_SERVER_ERROR",
		Message: "An internal server error occurred",
	}

	// Network Errors
	ErrNetworkError = &WeatherError{
		Code:    http.StatusBadGateway,
		Type:    "NETWORK_ERROR",
		Message: "Network error occurred while fetching weather data",
	}

	ErrConnectionFailed = &WeatherError{
		Code:    http.StatusBadGateway,
		Type:    "CONNECTION_FAILED",
		Message: "Failed to connect to weather service",
	}
)

// Helper functions to create custom errors
func NewValidationError(message, details string) *WeatherError {
	return &WeatherError{
		Code:    http.StatusBadRequest,
		Type:    "VALIDATION_ERROR",
		Message: message,
		Details: details,
	}
}

func NewServiceError(message, details string) *WeatherError {
	return &WeatherError{
		Code:    http.StatusServiceUnavailable,
		Type:    "SERVICE_ERROR",
		Message: message,
		Details: details,
	}
}

func NewInternalError(message, details string) *WeatherError {
	return &WeatherError{
		Code:    http.StatusInternalServerError,
		Type:    "INTERNAL_ERROR",
		Message: message,
		Details: details,
	}
}

// Error wrapping function
func WrapError(err error, weatherErr *WeatherError) *WeatherError {
	return &WeatherError{
		Code:    weatherErr.Code,
		Type:    weatherErr.Type,
		Message: weatherErr.Message,
		Details: fmt.Sprintf("%s: %v", weatherErr.Details, err),
	}
}

// IsWeatherError checks if an error is a WeatherError
func IsWeatherError(err error) (*WeatherError, bool) {
	if weatherErr, ok := err.(*WeatherError); ok {
		return weatherErr, true
	}
	return nil, false
}
