package wigal

import (
	"context"
	"net/http"
)

func (c *WigalClient) SentMessagesHistory(ctx context.Context, service Service, serviceType ServiceType,
	dateFrom, dateTo, senderId string,
	status Status, msgId string) (*HistoryResponseModel, error) {

	body := &HistoryRequestModel{
		Service:     service.String(),
		Servicetype: serviceType.String(),
		Datefrom:    dateFrom,
		Dateto:      dateTo,
		Senderid:    senderId,
		Status:      status.String(),
		Msgid:       msgId,
	}

	headers := http.Header{}
	headers.Add("Content-Type", "application/json")
	sentMessagesHistoryResponse, err := requester[HistoryResponseModel](ctx,
		"https://frogapi.wigal.com.gh/api/v3/sms/history",
		http.MethodPost, nil, headers, body)

	if err != nil {
		return nil, err
	}

	return sentMessagesHistoryResponse, nil
}
