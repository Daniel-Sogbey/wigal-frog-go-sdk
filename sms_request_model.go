package main

type SMSRequest struct {
	SenderId     string        `json:"senderid,omitempty"`
	Destinations []Destination `json:"destinations,omitempty"`
}

type Destination struct {
	Destination string `json:"destination,omitempty"`
	Message     string `json:"message,omitempty"`
	MsgId       string `json:"msgid,omitempty"`
	SmsType     string `json:"smstype,omitempty"`
}
