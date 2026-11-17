package v1

import (
	"errors"
	"strconv"

	models "github.com/Projectoutlast/ypracticum-metrics/internal/model"
)

func parseMetricValue(mType, value string) (float64, error) {
	var res float64

	switch mType {
	case models.Counter:
		parsed, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return 0, err
		}

		res = float64(parsed)
	case models.Gauge:
		parsed, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return 0, err
		}

		res = parsed
	default:
		return 0, errors.New("unknown metric type")
	}

	return res, nil
}

func parseMetricName(name string) (string, error) {
	if name == "" {
		return "", errors.New("metric name are required")
	}

	if len(name) > maxMetricNameChars {
		return "", errors.New("metrict name length must be less than or equal 100 letters")
	}

	return name, nil
}
