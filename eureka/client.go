package eureka

import (
	"bytes"
	"net/http"
	"time"
)

type Config struct {
	CertFile    string        `json:"certFile"`
	KeyFile     string        `json:"keyFile"`
	CaCertFile  []string      `json:"caCertFiles"`
	DialTimeout time.Duration `json:"timeout"`
}

type Client struct {
	URL string
}

func NewClient(baseURL string) *Client {
	client := &Client{
		URL: baseURL,
	}
	return client
}

func (c *Client) SendRequest(request *Request) (*http.Response, error) {
	req, err := http.NewRequest(request.Method, c.URL + request.Path, bytes.NewBuffer(request.Body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return response, nil
}
