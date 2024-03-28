package biz

import (
	"sort"

	"github.com/blackhorseya/sion/entity/domain/rental/agg"
	"github.com/blackhorseya/sion/entity/domain/rental/biz"
	"github.com/blackhorseya/sion/entity/domain/rental/model"
	"github.com/blackhorseya/sion/entity/domain/rental/repo"
	"github.com/blackhorseya/sion/pkg/contextx"
	"go.uber.org/zap"
)

type impl struct {
	assets repo.IAssetRepo
}

// NewRentalBiz will create a new rental biz.
func NewRentalBiz(assets repo.IAssetRepo) biz.IRentalBiz {
	return &impl{assets: assets}
}

func (i *impl) ListByLocation(
	ctx contextx.Contextx,
	location *model.Location,
	opts biz.ListByLocationOptions,
) (rentals []*agg.Asset, total int, err error) {
	cars, err := i.assets.FetchAvailableCars(ctx)
	if err != nil {
		ctx.Error("failed to fetch available cars", zap.Error(err))
		return nil, 0, err
	}

	for _, car := range cars {
		car.Distance = car.Location.DistanceTo(location, "K")
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].Distance < cars[j].Distance
	})

	end := len(cars)
	if opts.Size < end {
		end = opts.Size
	}

	return cars[:end], len(cars), nil
}
