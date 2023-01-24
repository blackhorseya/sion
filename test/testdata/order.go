package testdata

import (
	om "github.com/blackhorseya/irent/pkg/entity/domain/order/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	// Booking1 booking 1
	Booking1 = &om.Booking{
		No:         "id1",
		LastPickAt: timestamppb.Now(),
	}
)
