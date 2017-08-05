package poloniex

import (
	"context"
	"net/http"
)

type Client interface {
}

type httpClient struct {
	client *http.Client
}

func (c *httpClient) Do(ctx *context.Context, url string) {
}
