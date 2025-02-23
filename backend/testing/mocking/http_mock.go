package main

import "net/http"

type mockHTTPClient struct {
	GetResult *http.Response
	GetError  error
}

func (m *mockHTTPClient) Get(url string) (*http.Response, error) {
	return m.GetResult, m.GetError
}
