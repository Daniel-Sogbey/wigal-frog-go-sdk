package main

import (
	"context"
	"net/http"
)

func (c *WigalClient) SendMessage(ctx context.Context, to, message, msgId, smsType string) (*SMSResponse, error) {
	body := SMSRequest{
		SenderId: c.config.SenderId,
		Destinations: []Destination{
			{
				Destination: to,
				Message:     message,
				MsgId:       msgId,
				SmsType:     smsType,
			},
		},
	}

	headers := http.Header{}
	headers.Add("Content-Type", "application/json")
	headers.Add("API-KEY", c.config.ApiKey)
	headers.Add("USERNAME", c.config.Username)

	smsResponse, err := requester[SMSResponse](ctx, "https://frogapi.wigal.com.gh/api/v3/sms/send", http.MethodPost,
		nil,
		headers, body)

	if err != nil {
		return nil, err
	}

	return smsResponse, nil
}
