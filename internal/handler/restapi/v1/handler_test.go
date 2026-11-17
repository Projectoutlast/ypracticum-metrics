package v1

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	models "github.com/Projectoutlast/ypracticum-metrics/internal/model"
	"github.com/Projectoutlast/ypracticum-metrics/pkg/logger"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type fakeMetricUC struct{}

func (f *fakeMetricUC) Get(ctx context.Context) map[string]*models.Metrics {
	return nil
}

func (f *fakeMetricUC) Update(ctx context.Context, metric *models.Metrics) {}

func (f *fakeMetricUC) GetByID(ctx context.Context, id string) (*models.Metrics, error) {
	if id == "Alloc" {
		return &models.Metrics{Value: new(34.4)}, nil
	}

	if id == "PreAlloc" {
		return nil, models.ErrMetricNotFound
	}

	return nil, nil
}

func testRequest(t *testing.T, ts *httptest.Server, method, path string) *http.Response {
	req, err := http.NewRequest(method, ts.URL+path, nil)
	require.NoError(t, err)

	resp, err := ts.Client().Do(req)
	require.NoError(t, err)

	defer resp.Body.Close()

	return resp
}

func updateMetricsRouter() chi.Router {
	r := chi.NewRouter()

	v := &V1{m: &fakeMetricUC{}, l: logger.New("")}

	r.Post("/update/{type}/{name}/{value}", v.updateMetrics)

	return r
}

func getMetricsRouter() chi.Router {
	r := chi.NewRouter()

	v := &V1{m: &fakeMetricUC{}, l: logger.New("")}

	r.Get("/value/{type}/{name}", v.getMetric)

	return r
}

func metricsRouter() chi.Router {
	r := chi.NewRouter()

	v := &V1{m: &fakeMetricUC{}, l: logger.New("")}

	r.Get("/", v.metrics)

	return r
}

func TestV1_updateMetrics(t *testing.T) {
	ts := httptest.NewServer(updateMetricsRouter())
	defer ts.Close()

	tests := []struct {
		name   string
		url    string
		method string
		code   int
	}{
		{
			name:   "StatusOk",
			url:    "/update/counter/Alloc/45",
			method: http.MethodPost,
			code:   200,
		},
		{
			name:   "MethodNotAllowed",
			url:    "/update/counter/Alloc/45",
			method: http.MethodPut,
			code:   405,
		},
		{
			name:   "StatusBadRequest(invalid type)",
			url:    "/update/counterr/Alloc/45",
			method: http.MethodPost,
			code:   400,
		},
		{
			name:   "StatusBadRequest(empty name)",
			url:    "/update/counterr//45",
			method: http.MethodPost,
			code:   400,
		},
		{
			name:   "StatusBadRequest(invalid value)",
			url:    "/update/counterr/Alloc/hello",
			method: http.MethodPost,
			code:   400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := testRequest(t, ts, tt.method, tt.url)
			defer resp.Body.Close()

			assert.Equal(t, tt.code, resp.StatusCode)
		})
	}
}

func TestV1_getMetric(t *testing.T) {
	ts := httptest.NewServer(getMetricsRouter())
	defer ts.Close()

	tests := []struct {
		name   string
		url    string
		method string
		code   int
	}{
		{
			name:   "StatusOk",
			url:    "/value/counter/Alloc",
			method: http.MethodGet,
			code:   200,
		},
		{
			name:   "unknown metric type",
			url:    "/value/evil/Alloc",
			method: http.MethodGet,
			code:   400,
		},
		{
			name:   "empty name",
			url:    "/value//Alloc",
			method: http.MethodGet,
			code:   400,
		},
		{
			name:   "not found",
			url:    "/value/gauge/PreAlloc",
			method: http.MethodGet,
			code:   404,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := testRequest(t, ts, tt.method, tt.url)
			defer resp.Body.Close()

			assert.Equal(t, tt.code, resp.StatusCode)
		})
	}
}

func TestV1_metrics(t *testing.T) {
	ts := httptest.NewServer(metricsRouter())
	defer ts.Close()

	tests := []struct {
		name   string
		url    string
		method string
		code   int
	}{
		{
			name:   "StatusOk",
			url:    "/",
			method: http.MethodGet,
			code:   200,
		},
		{
			name:   "method not allowed",
			url:    "/",
			method: http.MethodPost,
			code:   405,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := testRequest(t, ts, tt.method, tt.url)
			defer resp.Body.Close()

			assert.Equal(t, tt.code, resp.StatusCode)
		})
	}
}
