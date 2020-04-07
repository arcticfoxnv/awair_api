package awair_api

import (
  "io/ioutil"
  "net/http"
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestClientDeviceAPIUsage(t *testing.T) {
  h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "Bearer abc123", r.Header.Get("Authorization"))

    data, _ := ioutil.ReadFile("testdata/DeviceAPIUsage.json")

		w.Write([]byte(data))
	})

  httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli := NewClient("abc123", SetHTTPClient(httpClient))

	data, err := cli.DeviceAPIUsage("awair-c", 0)

	assert.Nil(t, err)
	assert.Equal(t, 4, len(data.Usages))
}

func TestClientDevices(t *testing.T) {
  h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "Bearer abc123", r.Header.Get("Authorization"))

    data, _ := ioutil.ReadFile("testdata/Devices.json")

		w.Write([]byte(data))
	})

  httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli := NewClient("abc123", SetHTTPClient(httpClient))

	data, err := cli.Devices()

	assert.Nil(t, err)
	assert.Equal(t, 2, len(data.Devices))
}

func TestClientUserInfo(t *testing.T) {
  h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "Bearer abc123", r.Header.Get("Authorization"))

    data, _ := ioutil.ReadFile("testdata/UserInfo.json")

		w.Write([]byte(data))
	})

  httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli := NewClient("abc123", SetHTTPClient(httpClient))

	data, err := cli.UserInfo()

	assert.Nil(t, err)
	assert.Equal(t, "Kim", data.LastName)
	assert.Equal(t, "Steve", data.FirstName)
}
