//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package biz

import (
	"github.com/blackhorseya/sion/entity/domain/rental/agg"
	"github.com/blackhorseya/sion/entity/domain/rental/model"
	"github.com/blackhorseya/sion/pkg/contextx"
)

// ListByLocationOptions is the options for ListByLocation.
type ListByLocationOptions struct {
	Page int
	Size int
}

// IRentalBiz is an interface for rental biz.
type IRentalBiz interface {
	ListByLocation(
		ctx contextx.Contextx,
		location *model.Location,
		opts ListByLocationOptions,
	) (rentals []*agg.Asset, total int, err error)
}
