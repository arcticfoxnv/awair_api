package awair_api

import (
	"context"
	"crypto/tls"
	"github.com/stretchr/testify/assert"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testingHTTPClient(handler http.Handler) (*http.Client, func()) {
	s := httptest.NewTLSServer(handler)

	cli := &http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, network, _ string) (net.Conn, error) {
				return net.Dial(network, s.Listener.Addr().String())
			},
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	return cli, s.Close
}

func TestClientNewGetRequest(t *testing.T) {
	cli := NewClient("abc123")
	r, err := cli.newGetRequest("v1", "users/self")

	assert.Nil(t, err)
	assert.Equal(t, "GET", r.Method)
	assert.Equal(t, "application/json", r.Header.Get("Accept"))
	assert.Equal(t, "awair_api_client (https://github.com/arcticfoxnv/awair_api)", r.Header.Get("User-Agent"))
	assert.Equal(t, "Bearer abc123", r.Header.Get("Authorization"))
}
