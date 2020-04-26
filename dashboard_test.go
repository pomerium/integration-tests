package main

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

func TestDashboard(t *testing.T) {
	ctx := mainCtx
	ctx, clearTimeout := context.WithTimeout(ctx, time.Second*30)
	defer clearTimeout()

	t.Run("image asset", func(t *testing.T) {
		req, err := http.NewRequestWithContext(ctx, "GET", "https://httpdetails.localhost.pomerium.io/.pomerium/assets/img/pomerium.svg", nil)
		if err != nil {
			t.Fatal(err)
		}

		res, err := testcluster.Do(req)
		if !assert.NoError(t, err, "unexpected http error") {
			return
		}
		defer res.Body.Close()

		assert.Equal(t, http.StatusOK, res.StatusCode, "unexpected status code")
		assert.Equal(t, "image/svg+xml", res.Header.Get("Content-Type"))
	})
}
