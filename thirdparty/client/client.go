package client

import (
	"context"
	"net/http"

	"github.com/Khan/genqlient/graphql"
)

type Client struct {
	gqlClient graphql.Client
}

func New(endpoint string, httpHeaders map[string]string) *Client {
	httpClient := &http.Client{
		Transport: &authTransport{
			headers: httpHeaders,
			wrapped: http.DefaultTransport,
		},
	}

	return &Client{
		gqlClient: graphql.NewClient(endpoint, httpClient),
	}
}

type MessageCreateResult struct {
	Id      int
	Content string
}

func (c *Client) MessageCreate(ctx context.Context, content string) (*MessageCreateResult, error) {
	resp, err := MessageCreate(ctx, c.gqlClient, content)
	if err != nil {
		return nil, err
	}
	return &MessageCreateResult{
		Id:      resp.MessageCreate.Id,
		Content: resp.MessageCreate.Content,
	}, nil
}

func (c *Client) ErrorCreate(ctx context.Context, errMsg string) error {
	if _, err := ErrorCreate(ctx, c.gqlClient, errMsg); err != nil {
		return err
	}
	return nil
}

const (
	EndpointUrl string = "http://host.docker.internal:8001/graphql"
)

type authTransport struct {
	headers map[string]string
	wrapped http.RoundTripper
}

func (t *authTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	for k, v := range t.headers {
		r.Header.Set(k, v)
	}
	return t.wrapped.RoundTrip(r)
}
