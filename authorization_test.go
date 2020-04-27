package main

import (
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/calebdoxsey/pomerium-integration-tests/internal/flows"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

func TestAuthorization(t *testing.T) {
	ctx, clearTimeout := context.WithTimeout(mainCtx, time.Second*30)
	defer clearTimeout()

	t.Run("public", func(t *testing.T) {
		t.Skip() // pomerium doesn't currently handle unauthenticated public routes

		client := testcluster.NewHTTPClient()

		req, err := http.NewRequestWithContext(ctx, "GET", "https://httpdetails.localhost.pomerium.io", nil)
		if err != nil {
			t.Fatal(err)
		}

		res, err := client.Do(req)
		if !assert.NoError(t, err, "unexpected http error") {
			return
		}
		defer res.Body.Close()

		assert.Equal(t, http.StatusOK, res.StatusCode, "unexpected status code, headers=%v", res.Header)
	})

	t.Run("domains", func(t *testing.T) {
		t.Run("allowed", func(t *testing.T) {
			client := testcluster.NewHTTPClient()
			res, err := flows.Authenticate(ctx, client, mustParseURL("https://httpdetails.localhost.pomerium.io/by-domain"), "bob@dogs.test", []string{"user"})
			if assert.NoError(t, err) {
				assert.Equal(t, http.StatusOK, res.StatusCode, "expected OK for dogs.test")
			}
		})
		t.Run("not allowed", func(t *testing.T) {
			client := testcluster.NewHTTPClient()
			res, err := flows.Authenticate(ctx, client, mustParseURL("https://httpdetails.localhost.pomerium.io/by-domain"), "joe@cats.test", []string{"user"})
			if assert.NoError(t, err) {
				assert.Equal(t, http.StatusForbidden, res.StatusCode, "expected Forbidden for cats.test")
			}
		})
	})
	t.Run("users", func(t *testing.T) {
		t.Run("allowed", func(t *testing.T) {
			client := testcluster.NewHTTPClient()
			res, err := flows.Authenticate(ctx, client, mustParseURL("https://httpdetails.localhost.pomerium.io/by-user"), "bob@dogs.test", []string{"user"})
			if assert.NoError(t, err) {
				assert.Equal(t, http.StatusOK, res.StatusCode, "expected OK for bob@dogs.test")
			}
		})
		t.Run("not allowed", func(t *testing.T) {
			client := testcluster.NewHTTPClient()
			res, err := flows.Authenticate(ctx, client, mustParseURL("https://httpdetails.localhost.pomerium.io/by-user"), "joe@cats.test", []string{"user"})
			if assert.NoError(t, err) {
				assert.Equal(t, http.StatusForbidden, res.StatusCode, "expected Forbidden for joe@cats.test")
			}
		})
	})
	t.Run("groups", func(t *testing.T) {
		t.Run("allowed", func(t *testing.T) {
			client := testcluster.NewHTTPClient()
			res, err := flows.Authenticate(ctx, client, mustParseURL("https://httpdetails.localhost.pomerium.io/by-group"), "bob@dogs.test", []string{"admin", "user"})
			if assert.NoError(t, err) {
				assert.Equal(t, http.StatusOK, res.StatusCode, "expected OK for admin")
			}
		})
		t.Run("not allowed", func(t *testing.T) {
			client := testcluster.NewHTTPClient()
			res, err := flows.Authenticate(ctx, client, mustParseURL("https://httpdetails.localhost.pomerium.io/by-group"), "joe@cats.test", []string{"user"})
			if assert.NoError(t, err) {
				assert.Equal(t, http.StatusForbidden, res.StatusCode, "expected Forbidden for user")
			}
		})
	})
}
func mustParseURL(str string) *url.URL {
	u, err := url.Parse(str)
	if err != nil {
		panic(err)
	}
	return u
}
