package v1

import (
	"errors"
	"fmt"
	"net/http"
	"slices"
	"strconv"
	"strings"

	models "github.com/Projectoutlast/ypracticum-metrics/internal/model"
)

func (r *V1) getMetric(w http.ResponseWriter, req *http.Request) {
	if !slices.Contains([]string{models.Counter, models.Gauge}, req.PathValue("type")) {
		r.l.Error("v1 - getMetric - req.PathValue(type): metrict type must be counter or gauge")

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	name, err := parseMetricName(req.PathValue("name"))
	if err != nil {
		r.l.Error("v1 - getMetric - parseMetricName: ", "error", err)

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	m, err := r.m.GetByID(req.Context(), name)
	if err != nil {
		if errors.Is(err, models.ErrMetricNotFound) {
			w.WriteHeader(http.StatusNotFound)

			return
		}

		r.l.Error("v1 - getMetric - r.m.GetByID: ", "error", err)

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write([]byte(strconv.FormatFloat(*m.Value, 'f', -1, 64))); err != nil {
		r.l.Error("v1 - getMetric - w.Write: ", "error", err)
	}
}

func (r *V1) metrics(w http.ResponseWriter, req *http.Request) {
	m := r.m.Get(req.Context())

	var res []string

	for k, v := range m {
		var val float64

		if v.Value != nil {
			val = *v.Value
		}

		res = append(res, fmt.Sprintf("%s: %f", k, val))
	}

	metricList := strings.Join(res, "\n\n")

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write([]byte(metricList)); err != nil {
		r.l.Error("v1 - metrics - w.Write: ", "error", err)
	}
}
