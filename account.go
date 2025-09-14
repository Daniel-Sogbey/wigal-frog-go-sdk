package main

import "net/http"

func (c *Client) GetAccountBalance() (*AccountBalanceResponseModel, error) {
	headers := http.Header{}
	headers.Add("Content-Type", "application/json")
	accountBalanceResponse, err := requester[AccountBalanceResponseModel]("https://frogapi.wigal.com.gh/api/v3/balance",
		http.MethodGet, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	return accountBalanceResponse, nil
}
