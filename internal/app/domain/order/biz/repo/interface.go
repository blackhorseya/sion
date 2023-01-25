package repo

import (
	"github.com/blackhorseya/irent/pkg/contextx"
	am "github.com/blackhorseya/irent/pkg/entity/domain/account/model"
	om "github.com/blackhorseya/irent/pkg/entity/domain/order/model"
	rm "github.com/blackhorseya/irent/pkg/entity/domain/rental/model"
	"github.com/google/wire"
)

// IRepo declare account repository functions
//
//go:generate mockery --all --inpackage
type IRepo interface {
	// FetchArrears serve caller to fetch arrears via api
	FetchArrears(ctx contextx.Contextx, from *am.Profile, target *am.Profile) (records []*om.ArrearsRecord, err error)

	// BookCar serve caller to given user and car to book
	BookCar(ctx contextx.Contextx, from *am.Profile, target *rm.Car) (info *om.Booking, err error)

	// CancelBooking serve caller to given booking to cancel
	CancelBooking(ctx contextx.Contextx, from *am.Profile, target *om.Booking) error
}

var OrderSet = wire.NewSet(NewImpl, NewOptions)
