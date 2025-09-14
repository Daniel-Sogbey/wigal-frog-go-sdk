package main

type WigalClient struct {
	config *Config
}

func NewWigalClient(c *Config) *WigalClient {
	return &WigalClient{config: c}
}
