package weather

type TemperatureRanges struct {
	Label string
	Min   float64
	Max   float64
}

var celsiusRanges = []TemperatureRanges{
	{"Freezing", -273.15, 0},
	{"Cold", 0, 15},
	{"Moderate", 15, 25},
	{"Hot", 25, 35},
	{"Very Hot", 35, 1000},
}

var fahrenheitRanges = []TemperatureRanges{
	{"Freezing", -459.67, 32},
	{"Cold", 32, 59},
	{"Moderate", 59, 77},
	{"Hot", 77, 95},
	{"Very Hot", 95, 180},
}

const (
	METRIC   = "METRIC"
	IMPERIAL = "IMPERIAL"
)

func CharacterizeTemperature(temp float64, unit string) string {
	var ranges []TemperatureRanges
	switch unit {
	case METRIC:
		ranges = celsiusRanges
	case IMPERIAL:
		ranges = fahrenheitRanges
	default:
		return "Unknown unit"
	}

	for _, r := range ranges {
		if temp >= r.Min && temp < r.Max {
			return r.Label
		}
	}
	return "Unknown"
}
