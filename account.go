package main

import (
	"context"
	"net/http"
)

func (c *Client) GetAccountBalance(ctx context.Context) (*AccountBalanceResponseModel, error) {
	headers := http.Header{}
	headers.Add("Content-Type", "application/json")
	accountBalanceResponse, err := requester[AccountBalanceResponseModel](ctx,
		"https://frogapi.wigal.com.gh/api/v3/balance",
		http.MethodGet, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	return accountBalanceResponse, nil
}
