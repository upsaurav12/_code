package user

import "net/http"

type MockHttpClient struct {
	Response *http.Response
	Err      error
}
