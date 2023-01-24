package repo

type fetchArrearsResp struct {
	Result       string `json:"Result"`
	ErrorCode    string `json:"ErrorCode"`
	NeedRelogin  int    `json:"NeedRelogin"`
	NeedUpgrade  int    `json:"NeedUpgrade"`
	ErrorMessage string `json:"ErrorMessage"`
	Data         struct {
		ArrearsInfos []struct {
			NPR330SaveID        int    `json:"NPR330Save_ID"`
			RentAmount          int    `json:"Rent_Amount"`
			TicketAmount        int    `json:"Ticket_Amount"`
			ParkAmount          int    `json:"Park_Amount"`
			ETAGAmount          int    `json:"ETAG_Amount"`
			OperatingLossAmount int    `json:"OperatingLoss_Amount"`
			TotalAmount         int    `json:"Total_Amount"`
			StartDate           string `json:"StartDate"`
			EndDate             string `json:"EndDate"`
			OrderNo             string `json:"OrderNo"`
			ShortOrderNo        string `json:"ShortOrderNo"`
			StationName         string `json:"StationName"`
			CarType             string `json:"CarType"`
			IsMotor             int    `json:"IsMotor"`
		} `json:"ArrearsInfos"`
		TotalAmount  int    `json:"TotalAmount"`
		TradeOrderNo string `json:"TradeOrderNo"`
	} `json:"Data"`
}

type bookResp struct {
	Result       string `json:"Result"`
	ErrorCode    string `json:"ErrorCode"`
	NeedRelogin  int    `json:"NeedRelogin"`
	NeedUpgrade  int    `json:"NeedUpgrade"`
	ErrorMessage string `json:"ErrorMessage"`
	Data         struct {
		OrderNo      string `json:"OrderNo"`
		LastPickTime string `json:"LastPickTime"`
	} `json:"Data"`
}

type cancelBookingResp struct {
	Result       string `json:"Result"`
	ErrorCode    string `json:"ErrorCode"`
	NeedRelogin  int    `json:"NeedRelogin"`
	NeedUpgrade  int    `json:"NeedUpgrade"`
	ErrorMessage string `json:"ErrorMessage"`
	Data         struct {
	} `json:"Data"`
}
