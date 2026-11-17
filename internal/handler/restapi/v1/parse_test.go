package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMetricValue(t *testing.T) {
	tests := []struct {
		name  string
		mType string
		value string
		want  float64
	}{
		{
			name:  "positive",
			mType: "gauge",
			value: "10.0",
			want:  10.0,
		},
		{
			name:  "unknown type",
			mType: "gaugee",
			value: "10.0",
			want:  0,
		},
		{
			name:  "unparsed value",
			mType: "counter",
			value: "hello",
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := parseMetricValue(tt.mType, tt.value)

			assert.Equal(t, tt.want, got)
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
