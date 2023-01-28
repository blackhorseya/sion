package biz

import (
	"sort"

	"github.com/blackhorseya/irent/internal/app/domain/rental/biz/repo"
	"github.com/blackhorseya/irent/internal/pkg/errorx"
	"github.com/blackhorseya/irent/pkg/contextx"
	"github.com/blackhorseya/irent/pkg/distance"
	rb "github.com/blackhorseya/irent/pkg/entity/domain/rental/biz"
	rm "github.com/blackhorseya/irent/pkg/entity/domain/rental/model"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var RentalSet = wire.NewSet(NewImpl)

type impl struct {
	repo repo.IRepo
}

func NewImpl(repo repo.IRepo) rb.IBiz {
	return &impl{
		repo: repo,
	}
}

func (i *impl) Readiness(ctx contextx.Contextx) error {
	return nil
}

func (i *impl) Liveness(ctx contextx.Contextx) error {
	return nil
}

func (i *impl) ListCars(ctx contextx.Contextx, condition rb.QueryCarCondition) (info []*rm.Car, total int, err error) {
	ret, err := i.repo.ListCars(ctx)
	if err != nil {
		ctx.Error(errorx.ErrListCars.Error(), zap.Error(err))
		return nil, 0, errorx.ErrListCars
	}
	if len(ret) == 0 {
		return nil, 0, nil
	}

	for _, car := range ret {
		car.Distance = distance.CalcWithGeo(condition.Latitude, condition.Longitude, car.Latitude, car.Longitude, "K")
	}

	sort.Slice(ret, func(i, j int) bool {
		return ret[i].Distance < ret[j].Distance
	})

	return ret[:condition.TopNum], len(ret), nil
}

func (i *impl) UpdateInfoCars(ctx contextx.Contextx) (cars []*rm.Car, err error) {
	ret, err := i.repo.FetchAvailableCars(ctx)
	if err != nil {
		ctx.Error(errorx.ErrListCars.Error(), zap.Error(err))
		return nil, errorx.ErrListCars.ReplaceMsg(err.Error())
	}

	err = i.repo.UpdateStatusAllCars(ctx, rm.CarStatus_CAR_STATUS_INUSE)
	if err != nil {
		ctx.Error(errorx.ErrUpdateCar.Error(), zap.Error(err))
		return nil, errorx.ErrUpdateCar
	}

	for _, car := range ret {
		car.Status = rm.CarStatus_CAR_STATUS_AVAILABLE
		err = i.repo.UpsertStatusCar(ctx, car)
		if err != nil {
			ctx.Warn(errorx.ErrUpdateCar.Error(), zap.Error(err), zap.Any("car", car))
			continue
		}
	}

	return ret, nil
}
