package wigal

type AccountBalanceResponseModel struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Paidcashbalance int     `json:"paidcashbalance"`
		Cashbalance     float64 `json:"cashbalance"`
		Bundles         struct {
			Voice     int `json:"VOICE"`
			Kycverify int `json:"KYCVERIFY"`
			Simactive int `json:"SIMACTIVE"`
			Sms       int `json:"SMS"`
			Ussd      int `json:"USSD"`
		} `json:"bundles"`
		Invoicesummary []struct {
			Invoicetype string `json:"invoicetype"`
			Count       int    `json:"count"`
			Totalamount int    `json:"totalamount"`
		} `json:"invoicesummary"`
		Activebundleinvoices []struct {
			ID          int    `json:"id"`
			Description string `json:"description"`
			Enddate     string `json:"enddate"`
			Invoicetype string `json:"invoicetype"`
			Details     []struct {
				Service  string `json:"service"`
				Quantity int    `json:"quantity"`
				Used     int    `json:"used"`
			} `json:"details"`
		} `json:"activebundleinvoices"`
	} `json:"data"`
}
