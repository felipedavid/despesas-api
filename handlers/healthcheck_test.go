package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/felipedavid/saldop/test"
)

func TestHealthcheck(t *testing.T) {
	t.Parallel()

	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodGet, "/healthcheck", nil)
	if err != nil {
		t.Fatal(err)
	}

	healthcheck(rr, r)

	rs := rr.Result()

	test.Equal(t, rs.StatusCode, http.StatusOK)
}
