package v1

import (
	"net/http"
	"slices"

	models "github.com/Projectoutlast/ypracticum-metrics/internal/model"
)

const maxMetricNameChars = 100

func (r *V1) updateMetrics(w http.ResponseWriter, req *http.Request) {
	mType := req.PathValue("type")
	if !slices.Contains([]string{models.Counter, models.Gauge}, mType) {
		r.l.Error("v1 - updateMetrics - req.PathValue(type): metrict type must be counter or gauge")

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	name, err := parseMetricName(req.PathValue("name"))
	if err != nil {
		r.l.Error("v1 - updateMetrics - parseMetricName: ", "error", err)

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	mValue := req.PathValue("value")

	parsedValue, err := parseMetricValue(mType, mValue)
	if err != nil {
		r.l.Error("v1 - updateMetrics - parseMetricValue", "error", err)

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	m := &models.Metrics{
		ID:    name,
		MType: mType,
		Value: &parsedValue,
	}

	r.m.Update(req.Context(), m)

	w.WriteHeader(http.StatusOK)
}
