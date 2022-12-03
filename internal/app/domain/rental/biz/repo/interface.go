package repo

import (
	"github.com/blackhorseya/irent/pkg/contextx"
	rm "github.com/blackhorseya/irent/pkg/entity/domain/rental/model"
)

// IRepo declare rental repository functions
//
//go:generate mockery --all --inpackage
type IRepo interface {
	// ListCars serve caller to list all car
	ListCars(ctx contextx.Contextx) (info []*rm.Car, err error)
}
