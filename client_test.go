package apex_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"go.reefassistant.com/apex"
)

func TestHTTPClientTimeout(t *testing.T) {
	srv := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(100 * time.Millisecond)
			w.Write([]byte("foobar"))
		}),
	)
	defer srv.Close()

	client, err := apex.NewClient(srv.URL, apex.ClientConfig{
		HTTPClient: &http.Client{
			Timeout: 50 * time.Millisecond,
		},
	})
	require.NoError(t, err)

	req := client.PublicV1().DefaultApi.GetStatus(context.Background())
	_, _, err = client.PublicV1().DefaultApi.GetStatusExecute(req)
	require.Contains(t, err.Error(), context.DeadlineExceeded.Error())
}

func TestContextClientTimeout(t *testing.T) {
	srv := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(100 * time.Millisecond)
			w.Write([]byte("foobar"))
		}),
	)
	defer srv.Close()

	client, err := apex.NewDefaultClient(srv.URL)
	require.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	req := client.PublicV1().DefaultApi.GetStatus(ctx)
	_, _, err = client.PublicV1().DefaultApi.GetStatusExecute(req)
	require.Contains(t, err.Error(), context.DeadlineExceeded.Error())
}
