package httpx

import (
	"net/http"
)

// Client http client interface
//
//go:generate mockery --all --inpackage
type Client interface {
	// Do send an HTTP request and returns an HTTP response, following
	// policy (such as redirects, cookies, auth) as configured on the
	// client.
	Do(req *http.Request) (resp *http.Response, err error)
}
