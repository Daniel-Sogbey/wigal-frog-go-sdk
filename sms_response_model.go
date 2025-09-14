package main

type SMSResponse struct {
	Status  *string     `json:"status,omitempty"`
	Message *string     `json:"message,omitempty"`
	Data    interface{} `json:"data"`
	Errors  interface{} `json:"errors"`
}
