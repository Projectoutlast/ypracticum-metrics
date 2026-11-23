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
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	name, err := parseMetricName(req.PathValue("name"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	m, err := parseMetricValue(mType, req.PathValue("value"), name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	r.m.Update(req.Context(), m)

	w.WriteHeader(http.StatusOK)
}
