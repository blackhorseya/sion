package repo

type anyRentResponse struct {
	Result       string `json:"Result"`
	ErrorCode    string `json:"ErrorCode"`
	NeedRelogin  int    `json:"NeedRelogin"`
	NeedUpgrade  int    `json:"NeedUpgrade"`
	ErrorMessage string `json:"ErrorMessage"`
	Data         struct {
		AnyRentObj []struct {
			MonthlyRentID int     `json:"MonthlyRentId"`
			MonProjNM     string  `json:"MonProjNM"`
			CarWDHours    float64 `json:"CarWDHours"`
			CarHDHours    float64 `json:"CarHDHours"`
			MotoTotalMins int     `json:"MotoTotalMins"`
			WDRateForCar  float64 `json:"WDRateForCar"`
			HDRateForCar  float64 `json:"HDRateForCar"`
			WDRateForMoto float64 `json:"WDRateForMoto"`
			HDRateForMoto float64 `json:"HDRateForMoto"`
			DiscountLabel struct {
				LabelType  string `json:"LabelType"`
				GiveMinute int    `json:"GiveMinute"`
				Describe   string `json:"Describe"`
			} `json:"DiscountLabel"`
			ParkingImgObj  []interface{} `json:"ParkingImgObj"`
			CarNo          string        `json:"CarNo"`
			CarType        string        `json:"CarType"`
			CarTypeName    string        `json:"CarTypeName"`
			CarOfArea      string        `json:"CarOfArea"`
			ProjectName    string        `json:"ProjectName"`
			Rental         float64       `json:"Rental"`
			Mileage        float64       `json:"Mileage"`
			Insurance      int           `json:"Insurance"`
			InsurancePrice int           `json:"InsurancePrice"`
			ShowSpecial    int           `json:"ShowSpecial"`
			SpecialInfo    string        `json:"SpecialInfo"`
			Latitude       float64       `json:"Latitude"`
			Longitude      float64       `json:"Longitude"`
			Operator       string        `json:"Operator"`
			OperatorScore  float64       `json:"OperatorScore"`
			CarTypePic     string        `json:"CarTypePic"`
			Seat           int64         `json:"Seat"`
			ProjID         string        `json:"ProjID"`
			TaxID          string        `json:"TaxID"`
			ReportOrderNo  string        `json:"ReportOrderNo"`
			ParkingSpace   string        `json:"ParkingSpace"`
			ColorTag       int           `json:"ColorTag"`
			NowStationID   string        `json:"nowStationID"`
			AreaRemk       string        `json:"AreaRemk"`
			AreabtnText    string        `json:"AreabtnText"`
			AreaAlert      string        `json:"AreaAlert"`
			ParkingImgJSON string        `json:"ParkingImgJson"`
		} `json:"AnyRentObj"`
	} `json:"Data"`
}
