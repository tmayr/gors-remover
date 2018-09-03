package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

// ForwardingHTTPClient is a wrapper for http.Client
type ForwardingHTTPClient struct {
	client *http.Client // Internal HTTP client.
}

// NewForwardingHTTPClient creates an HTTP client
func NewForwardingHTTPClient() *ForwardingHTTPClient {
	c := &ForwardingHTTPClient{
		client: http.DefaultClient,
	}
	return c
}

// NewForwardingRequest accepts a Request object and forwards the exact headers, type and data to ?url
func (fw *ForwardingHTTPClient) NewForwardingRequest(r *http.Request) (*http.Response, error) {
	originalURL := r.URL.Query().Get("url")

	if originalURL == "" {
		return nil, errors.New("url query param cannot be empty")
	}

	_, err := url.ParseRequestURI(originalURL)
	if err != nil {
		return nil, errors.New("invalid url")
	}

	req, err := http.NewRequest(r.Method, originalURL, r.Body)
	if err != nil {
		return nil, err
	}

	// add all headers from the incoming request
	for name, headers := range r.Header {
		for _, h := range headers {
			req.Header.Set(name, h)
		}
	}

	// we dont want to deal with encodings tbh, let's not accept any
	req.Header.Del("accept-encoding")

	// print the request to the console for easier debuging
	fmt.Println(FormatRequest(req))

	res, err := fw.client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, err
}
