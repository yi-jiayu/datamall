package datamall

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type MockHTTPClient struct {
	StatusCode int
	Body       string
}

func (c *MockHTTPClient) Do(*http.Request) (*http.Response, error) {
	res := &http.Response{
		StatusCode: c.StatusCode,
		Body:       ioutil.NopCloser(strings.NewReader(c.Body)),
	}
	return res, nil
}
