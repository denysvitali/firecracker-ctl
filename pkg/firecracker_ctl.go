package firecrackerctl

import (
	"fmt"
	"github.com/peterbourgon/unixtransport"
	"net/http"
)

type Client struct {
	httpClient *http.Client
	sockUrl string
}

func New(url string) Client {
	t := http.Transport{}
	unixtransport.Register(&t)
	httpClient := &http.Client{Transport: &t}
	return Client{
		httpClient: httpClient,
		sockUrl: url,
	}
}

func (c *Client) url(s string) string {
	return fmt.Sprintf("http+unix://%s:%s", c.sockUrl, s)
}