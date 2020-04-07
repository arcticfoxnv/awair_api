package awair_api

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestClientUserLatestAirData(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "Bearer abc123", r.Header.Get("Authorization"))

		data, _ := ioutil.ReadFile("testdata/UserLatestAirData.json")

		w.Write([]byte(data))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli := NewClient("abc123", SetHTTPClient(httpClient))

	data, err := cli.UserLatestAirData("awair-c", 0)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(data.Data))
}

func TestClientUserLatestAirDataWithF(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "Bearer abc123", r.Header.Get("Authorization"))
		assert.Equal(t, "true", r.URL.Query().Get("farenheit"))

		data, _ := ioutil.ReadFile("testdata/UserLatestAirData.json")
		w.Write([]byte(data))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli := NewClient("abc123", SetHTTPClient(httpClient))
	cli.UseFarenheit = true

	data, err := cli.UserLatestAirData("awair-c", 0)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(data.Data))
}
