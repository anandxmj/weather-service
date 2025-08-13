package weather

import (
	"strconv"

	"github.com/anandxmj/weather-service/internal/errors"
)

func ValidateCoordinates(longitude string, latitude string) error {
	lonFlt, err := strconv.ParseFloat(longitude, 64)
	if err != nil {
		return errors.NewValidationError("Invalid longitude", longitude)
	}

	latFlt, err := strconv.ParseFloat(latitude, 64)
	if err != nil {
		return errors.NewValidationError("Invalid latitude", latitude)
	}

	if latFlt < -90 || latFlt > 90 {
		return errors.ErrCoordinatesOutOfRange
	}

	if lonFlt < -180 || lonFlt > 180 {
		return errors.ErrCoordinatesOutOfRange
	}

	return nil
}
