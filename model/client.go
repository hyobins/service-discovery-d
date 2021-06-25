package model

import "net/http"

type Client struct {
	address    string
	httpClient *http.Client
}

func NewClient(address string) *Client {
	return &Client{
		address:    address,
		httpClient: &http.Client{},
	}
}

func (c *Client) GetPods(req *GetPodsRequest) (string, error) {

}
