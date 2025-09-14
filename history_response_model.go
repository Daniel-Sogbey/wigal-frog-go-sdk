package main

type HistoryResponseModel struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Content []struct {
			Apimessageid   string  `json:"apimessageid"`
			Status         string  `json:"status"`
			Statusreason   string  `json:"statusreason"`
			Recipient      string  `json:"recipient"`
			Statusdate     string  `json:"statusdate"`
			Bundlecredits  int     `json:"bundlecredits"`
			Charge         float64 `json:"charge"`
			Service        string  `json:"service"`
			Servicetype    string  `json:"servicetype"`
			Messagecount   int     `json:"messagecount"`
			Charactercount int     `json:"charactercount"`
		} `json:"content"`
		Pageable struct {
			PageNumber int `json:"pageNumber"`
			PageSize   int `json:"pageSize"`
			Sort       struct {
				Sorted   bool `json:"sorted"`
				Empty    bool `json:"empty"`
				Unsorted bool `json:"unsorted"`
			} `json:"sort"`
			Offset  int  `json:"offset"`
			Unpaged bool `json:"unpaged"`
			Paged   bool `json:"paged"`
		} `json:"pageable"`
		Last          bool `json:"last"`
		TotalElements int  `json:"totalElements"`
		TotalPages    int  `json:"totalPages"`
		Size          int  `json:"size"`
		Number        int  `json:"number"`
		Sort          struct {
			Sorted   bool `json:"sorted"`
			Empty    bool `json:"empty"`
			Unsorted bool `json:"unsorted"`
		} `json:"sort"`
		First            bool `json:"first"`
		NumberOfElements int  `json:"numberOfElements"`
		Empty            bool `json:"empty"`
	} `json:"data"`
}
