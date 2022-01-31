package site

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIndexPageRenders(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(Handler())

	req, err := http.NewRequestWithContext(context.Background(), "GET", srv.URL, nil)
	require.NoError(t, err)
	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err, "get index")
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	require.NotEmpty(t, data, "index should have contents")
}