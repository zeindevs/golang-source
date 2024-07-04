package internal

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

type HttpClient struct {
	client *http.Client
	Header http.Header
}

func NewHttpClient() *HttpClient {
	return &HttpClient{
		client: &http.Client{},
		Header: http.Header{},
	}
}

func (hc *HttpClient) Get(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	req.Header = hc.Header
	if err != nil {
		return nil, err
	}

	res, err := hc.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected with status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
