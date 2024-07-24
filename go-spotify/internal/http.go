package internal

import (
	"io"
	"net/http"
)

type Http struct {
	http   *http.Client
	Header http.Header
}

func NewHttp() *Http {
	return &Http{
		http:   &http.Client{},
		Header: http.Header{},
	}
}

func (c *Http) Get(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	req.Header = c.Header
	if err != nil {
		return nil, err
	}
	res, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Http) Post(url string, data io.Reader) ([]byte, error) {
	req, err := http.NewRequest("POST", url, data)
	req.Header = c.Header
	if err != nil {
		return nil, err
	}
	res, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
