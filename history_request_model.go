package main

type HistoryRequestModel struct {
	Service     string `json:"service"`
	Servicetype string `json:"servicetype"`
	Datefrom    string `json:"datefrom"`
	Dateto      string `json:"dateto"`
	Senderid    string `json:"senderid"`
	Status      string `json:"status"`
	Msgid       string `json:"msgid"`
}
