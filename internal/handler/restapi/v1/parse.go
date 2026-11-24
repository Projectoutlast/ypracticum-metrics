package v1

import (
	"errors"
	"strconv"

	models "github.com/Projectoutlast/ypracticum-metrics/internal/model"
)

func parseMetricValue(mType, value, name string) (*models.Metrics, error) {
	m := &models.Metrics{ID: name, MType: mType}

	switch mType {
	case models.Counter:
		parsed, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, err
		}

		m.Delta = new(parsed)
	case models.Gauge:
		parsed, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, err
		}

		m.Value = new(parsed)
	default:
		return nil, errors.New("unknown metric type")
	}

	return m, nil
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
