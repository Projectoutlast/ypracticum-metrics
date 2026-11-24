package v1

import (
	"testing"

	models "github.com/Projectoutlast/ypracticum-metrics/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseMetricValue(t *testing.T) {
	tests := []struct {
		name  string
		mType string
		value string
		want  *models.Metrics
	}{
		{
			name:  "positive",
			mType: "gauge",
			value: "10.0",
			want:  &models.Metrics{MType: "gauge", Value: new(float64(10.0))},
		},
		{
			name:  "unknown type",
			mType: "gaugee",
			value: "10.0",
			want:  nil,
		},
		{
			name:  "unparsed value",
			mType: "counter",
			value: "10",
			want:  &models.Metrics{MType: "counter", Delta: new(int64(10))},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseMetricValue(tt.mType, tt.value, "")

			if tt.mType == "gaugee" {
				require.Error(t, err)
			} else {
				switch got.MType {
				case models.Counter:
					assert.Equal(t, *tt.want.Delta, *got.Delta)
				case models.Gauge:
					assert.Equal(t, *tt.want.Value, *got.Value)
				}
			}

		})
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  string
	}{
		{
			name:  "positive",
			value: "counter",
			want:  "counter",
		},
		{
			name:  "empty name",
			value: "",
			want:  "",
		},
		{
			name:  "name too long",
			value: "4jfj28jf283jf823jf23f823jf8j230fj0238fj8230jf082j3f8j20f823jf02jf023j0f2j30923j0f9j20fj20fj230fj230fj023jf02jf023jf0923j0f9j2309f",
			want:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := parseMetricName(tt.value)

			assert.Equal(t, tt.want, got)
		})
	}
}
