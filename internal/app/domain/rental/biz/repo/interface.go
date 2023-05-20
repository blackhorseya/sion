//go:generate mockgen -destination=./mock_${GOFILE} -package=repo -source=${GOFILE}
package repo

import (
	"github.com/blackhorseya/irent/pkg/contextx"
	rm "github.com/blackhorseya/irent/pkg/entity/domain/rental/model"
	"github.com/google/wire"
)

// IRepo declare rental repository functions
type IRepo interface {
	// ListCars serve caller to list all car
	ListCars(ctx contextx.Contextx) (info []*rm.Car, err error)

	// FetchAvailableCars serve caller to fetch available cars via official api
	FetchAvailableCars(ctx contextx.Contextx) (cars []*rm.Car, err error)

	// UpsertStatusCar serve caller to given status to insert or update the car
	UpsertStatusCar(ctx contextx.Contextx, target *rm.Car) error

	// UpdateStatusAllCars serve caller to reset status of all cars to inuse
	UpdateStatusAllCars(ctx contextx.Contextx, status rm.CarStatus) error
}

var RentalSet = wire.NewSet(NewOptions, NewImpl)
