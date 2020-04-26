package main

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

func TestAuthorization(t *testing.T) {
	ctx, clearTimeout := context.WithTimeout(mainCtx, time.Second*30)
	defer clearTimeout()

	t.Run("public", func(t *testing.T) {
		req, err := http.NewRequestWithContext(ctx, "GET", "https://httpdetails.localhost.pomerium.io", nil)
		if err != nil {
			t.Fatal(err)
		}

		res, err := testcluster.Do(req)
		if !assert.NoError(t, err, "unexpected http error") {
			return
		}
		defer res.Body.Close()

		assert.Equal(t, http.StatusOK, res.StatusCode, "unexpected status code, headers=%v", res.Header)
	})
}
