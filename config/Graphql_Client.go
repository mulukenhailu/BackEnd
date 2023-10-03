package config

import (
	"net/http"

	"github.com/hasura/go-graphql-client"
)

func GraphqlClient() *graphql.Client {
	hasuraEndPoint := "http://localhost:8080/v1/graphql"
	headers := http.Header{}

	client := graphql.NewClient(hasuraEndPoint, &http.Client{
		Transport: &headerTransport{Base: http.DefaultTransport, Headers: headers},
	})

	return client
}

type headerTransport struct {
	Base    http.RoundTripper
	Headers http.Header
}

func (t *headerTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	for key, values := range t.Headers {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}
	return t.Base.RoundTrip(req)
}
