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

type queryBookingsResp struct {
	Result       string `json:"Result"`
	ErrorCode    string `json:"ErrorCode"`
	NeedRelogin  int    `json:"NeedRelogin"`
	NeedUpgrade  int    `json:"NeedUpgrade"`
	ErrorMessage string `json:"ErrorMessage"`
	Data         struct {
		OrderObj []struct {
			StationInfo struct {
				StationID           string        `json:"StationID"`
				StationName         string        `json:"StationName"`
				Tel                 string        `json:"Tel"`
				ADDR                string        `json:"ADDR"`
				Latitude            float64       `json:"Latitude"`
				Longitude           float64       `json:"Longitude"`
				Content             string        `json:"Content"`
				IsRent              interface{}   `json:"IsRent"`
				ContentForAPP       string        `json:"ContentForAPP"`
				IsRequiredForReturn int           `json:"IsRequiredForReturn"`
				StationPic          []interface{} `json:"StationPic"`
			} `json:"StationInfo"`
			Operator          string      `json:"Operator"`
			OperatorScore     float64     `json:"OperatorScore"`
			CarTypePic        string      `json:"CarTypePic"`
			CarNo             string      `json:"CarNo"`
			CarBrend          string      `json:"CarBrend"`
			CarTypeName       string      `json:"CarTypeName"`
			Seat              int         `json:"Seat"`
			ParkingSection    string      `json:"ParkingSection"`
			IsMotor           int         `json:"IsMotor"`
			CarOfArea         string      `json:"CarOfArea"`
			CarLatitude       float64     `json:"CarLatitude"`
			CarLongitude      float64     `json:"CarLongitude"`
			MotorPowerBaseObj interface{} `json:"MotorPowerBaseObj"`
			ProjType          int         `json:"ProjType"`
			ProjName          string      `json:"ProjName"`
			WorkdayPerHour    int         `json:"WorkdayPerHour"`
			HolidayPerHour    int         `json:"HolidayPerHour"`
			MaxPrice          int         `json:"MaxPrice"`
			MaxPriceH         int         `json:"MaxPriceH"`
			MotorBasePriceObj interface{} `json:"MotorBasePriceObj"`
			OrderStatus       int         `json:"OrderStatus"`
			OrderNo           string      `json:"OrderNo"`
			StartTime         string      `json:"StartTime"`
			PickTime          string      `json:"PickTime"`
			ReturnTime        string      `json:"ReturnTime"`
			StopPickTime      string      `json:"StopPickTime"`
			StopTime          string      `json:"StopTime"`
			OpenDoorDeadLine  string      `json:"OpenDoorDeadLine"`
			CarRentBill       int         `json:"CarRentBill"`
			MileagePerKM      float64     `json:"MileagePerKM"`
			MileageBill       int         `json:"MileageBill"`
			Insurance         int         `json:"Insurance"`
			InsurancePerHour  int         `json:"InsurancePerHour"`
			InsuranceBill     int         `json:"InsuranceBill"`
			TransDiscount     int         `json:"TransDiscount"`
			Bill              int         `json:"Bill"`
			DailyMaxHour      int         `json:"DailyMaxHour"`
			CARMGTSTATUS      int         `json:"CAR_MGT_STATUS"`
			AppStatus         int         `json:"AppStatus"`
		} `json:"OrderObj"`
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
