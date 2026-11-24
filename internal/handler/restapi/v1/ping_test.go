package v1

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestV1_ping(t *testing.T) {
	type want struct {
		code        int
		response    string
		contentType string
	}

	api := &V1{}

	tests := []struct {
		name string
		want want
	}{
		{
			name: "",
			want: want{
				code:        200,
				response:    `{"status":"ok"}`,
				contentType: "application/json",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/ping", nil)

			w := httptest.NewRecorder()

			api.ping(w, req)

			res := w.Result()

			assert.Equal(t, res.StatusCode, tt.want.code)

			defer res.Body.Close()

			resBody, err := io.ReadAll(res.Body)

			require.NoError(t, err)

			assert.JSONEq(t, tt.want.response, string(resBody))
			assert.Equal(t, res.Header.Get("Content-Type"), tt.want.contentType)
		})
	}
}
