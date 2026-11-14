package v1

import (
	"net/http"
	"slices"
	"strconv"

	models "github.com/Projectoutlast/ypracticum-metrics/internal/model"
)

const maxMetricNameChars = 100

func (r *V1) updateMetrics(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}

	mType := req.PathValue("type")
	if !slices.Contains([]string{models.Counter, models.Gauge}, mType) {
		r.l.Error("v1 - updateMetrics - req.PathValue(type): metrict type must be counter or gauge")

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	mName := req.PathValue("name")
	if mName == "" {
		r.l.Error("v1 - updateMetrics - req.PathValue(name): metrict name empty - not found")

		w.WriteHeader(http.StatusNotFound)

		return
	}

	if len(mName) > maxMetricNameChars {
		r.l.Error("v1 - updateMetrics - req.PathValue(name): metrict name length must be less than or equal 100 letters")

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	mValue := req.PathValue("value")

	switch mType {
	case models.Counter:
		parsed, err := strconv.ParseInt(mValue, 10, 64)
		if err != nil {
			r.l.Error("v1 - updateMetrics - strconv.ParseInt", "error", err)

			w.WriteHeader(http.StatusBadRequest)

			return
		}

		if err := r.c.Update(req.Context(), parsed); err != nil {
			r.l.Error("v1 - updateMetrics - r.c.Update", "error", err)

			w.WriteHeader(http.StatusInternalServerError)

			return
		}

	case models.Gauge:
		parsed, err := strconv.ParseFloat(mValue, 64)
		if err != nil {
			r.l.Error("v1 - updateMetrics - strconv.ParseFloat", "error", err)

			w.WriteHeader(http.StatusBadRequest)

			return
		}

		if err := r.g.Update(req.Context(), parsed); err != nil {
			r.l.Error("v1 - updateMetrics - r.g.Update", "error", err)

			w.WriteHeader(http.StatusInternalServerError)

			return
		}
	}

	w.WriteHeader(http.StatusOK)
}
