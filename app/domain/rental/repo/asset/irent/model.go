package irent

import (
	"strings"

	"github.com/blackhorseya/sion/entity/domain/rental/agg"
	"github.com/blackhorseya/sion/entity/domain/rental/model"
)

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
			ParkingImgObj []struct {
				ParkingImage string `json:"ParkingImage"`
			} `json:"ParkingImgObj"`
			CarNo          string  `json:"CarNo"`
			CarType        string  `json:"CarType"`
			CarTypeName    string  `json:"CarTypeName"`
			CarOfArea      string  `json:"CarOfArea"`
			ProjectName    string  `json:"ProjectName"`
			Rental         float64 `json:"Rental"`
			Mileage        float64 `json:"Mileage"`
			Insurance      int     `json:"Insurance"`
			InsurancePrice int     `json:"InsurancePrice"`
			ShowSpecial    int     `json:"ShowSpecial"`
			SpecialInfo    string  `json:"SpecialInfo"`
			Latitude       float64 `json:"Latitude"`
			Longitude      float64 `json:"Longitude"`
			Operator       string  `json:"Operator"`
			OperatorScore  float64 `json:"OperatorScore"`
			CarTypePic     string  `json:"CarTypePic"`
			Seat           int     `json:"Seat"`
			ProjID         string  `json:"ProjID"`
			TaxID          string  `json:"TaxID"`
			ReportOrderNo  string  `json:"ReportOrderNo"`
			ParkingSpace   string  `json:"ParkingSpace"`
			ColorTag       int     `json:"ColorTag"`
			NowStationID   string  `json:"nowStationID"`
			AreaRemk       string  `json:"AreaRemk"`
			AreabtnText    string  `json:"AreabtnText"`
			AreaAlert      string  `json:"AreaAlert"`
			ParkingImgJSON string  `json:"ParkingImgJson"`
		} `json:"AnyRentObj"`
		ParkAreaObj []struct {
			ParkingPID  string  `json:"ParkingPID"`
			ParkingName string  `json:"ParkingName"`
			ParkingLat  float64 `json:"ParkingLat"`
			ParkingLng  float64 `json:"ParkingLng"`
			ProParking  int     `json:"ProParking"`
			HaveCar     string  `json:"HaveCar"`
			HaveQuota   string  `json:"HaveQuota"`
			HaveConn    string  `json:"HaveConn"`
			HaveFilter  string  `json:"HaveFilter"`
		} `json:"ParkAreaObj"`
		RegionalFlag int `json:"RegionalFlag"`
	} `json:"Data"`
}

// ToAggregate is a function to convert anyRentResponse to asset aggregate.
func (x *anyRentResponse) ToAggregate() []*agg.Asset {
	var ret []*agg.Asset
	for _, v := range x.Data.AnyRentObj {
		ret = append(ret, &agg.Asset{
			Car: &model.Car{
				ID: strings.ReplaceAll(v.CarNo, " ", ""),
				Location: model.Location{
					Latitude:  v.Latitude,
					Longitude: v.Longitude,
				},
			},
		})
	}

	return ret
}
