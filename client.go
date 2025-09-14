package main

type Client struct {
	config *Config
}

func NewClient(c *Config) *Client {
	return &Client{config: c}
}
