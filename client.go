package awair_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	AccessToken  string
	UseFarenheit bool
	UserAgent    string

	httpClient *http.Client
}

func NewClient(accessToken string, options ...func(*Client)) *Client {
	c := &Client{
		AccessToken: accessToken,
		UserAgent:   "awair_api_client (https://github.com/arcticfoxnv/awair_api)",
	}

	for _, option := range options {
		option(c)
	}

	return c
}

func (c *Client) getEndpoint(version, endpoint string) string {
	var base string

	switch version {
	case "v1":
		base = AwairV1
	default:
		base = AwairV1
	}

	return fmt.Sprintf("%s/%s", base, endpoint)
}

func (c *Client) appendQueryParam(req *http.Request, k, v string) {
	q := req.URL.Query()
	q.Set(k, v)
	req.URL.RawQuery = q.Encode()
}

func (c *Client) newGetRequest(version, endpoint string) (*http.Request, error) {
	url := c.getEndpoint(version, endpoint)
	return c.newRequest("GET", url, nil)
}

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, path, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))
	return req, nil
}

func (c *Client) do(req *http.Request, data interface{}) error {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return err
	}

	return nil
}
