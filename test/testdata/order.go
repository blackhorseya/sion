package testdata

import (
	om "github.com/blackhorseya/irent/pkg/entity/domain/order/model"
)

var (
	// Lease1 lease 1
	Lease1 = &om.Lease{
		No:           "id1",
		CarId:        "car1",
		CarLatitude:  0,
		CarLongitude: 0,
		StartAt:      nil,
		EndAt:        nil,
		LastPickAt:   nil,
	}
)
