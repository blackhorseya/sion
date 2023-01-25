package testdata

import (
	rm "github.com/blackhorseya/irent/pkg/entity/domain/rental/model"
)

var (
	// Car1 car 1
	Car1 = &rm.Car{
		Id:          "id1",
		CarType:     "",
		CarTypeName: "",
		CarOfArea:   "",
		ProjectName: "",
		ProjectId:   "pid1",
		Latitude:    0,
		Longitude:   0,
		Seat:        0,
		Distance:    0,
	}
)
