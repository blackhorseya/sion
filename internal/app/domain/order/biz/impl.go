package biz

import (
	"github.com/blackhorseya/irent/internal/app/domain/order/biz/repo"
	"github.com/blackhorseya/irent/internal/pkg/errorx"
	"github.com/blackhorseya/irent/pkg/contextx"
	am "github.com/blackhorseya/irent/pkg/entity/domain/account/model"
	ob "github.com/blackhorseya/irent/pkg/entity/domain/order/biz"
	om "github.com/blackhorseya/irent/pkg/entity/domain/order/model"
	rm "github.com/blackhorseya/irent/pkg/entity/domain/rental/model"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var OrderSet = wire.NewSet(NewImpl)

type impl struct {
	repo repo.IRepo
}

func NewImpl(repo repo.IRepo) ob.IBiz {
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

func (i *impl) GetArrears(ctx contextx.Contextx, from *am.Profile, target *am.Profile) (info *om.Arrears, err error) {
	if len(from.AccessToken) == 0 {
		ctx.Error(errorx.ErrMissingToken.Error(), zap.Any("from", from))
		return nil, errorx.ErrMissingToken
	}

	if len(target.Id) == 0 {
		ctx.Error(errorx.ErrMissingID.Error(), zap.Any("target", target))
		return nil, errorx.ErrMissingID
	}

	records, err := i.repo.FetchArrears(ctx, from, target)
	if err != nil {
		ctx.Error(errorx.ErrGetArrears.Error(), zap.Error(err), zap.Any("from", from), zap.Any("target", target))
		return nil, errorx.ErrGetArrears.ReplaceMsg(err.Error())
	}
	if len(records) == 0 {
		return nil, nil
	}

	ret := &om.Arrears{
		Records:     records,
		TotalAmount: 0,
	}
	for _, record := range records {
		ret.TotalAmount += record.TotalAmount
	}

	return ret, nil
}

func (i *impl) BookRental(ctx contextx.Contextx, from *am.Profile, target *rm.Car) (info *om.Booking, err error) {
	if len(from.AccessToken) == 0 {
		ctx.Error(errorx.ErrMissingToken.Error(), zap.Any("from", from))
		return nil, errorx.ErrMissingToken
	}

	if len(target.Id) == 0 || len(target.ProjectId) == 0 {
		ctx.Error(errorx.ErrInvalidRental.Error(), zap.Any("target", target))
		return nil, errorx.ErrInvalidRental
	}

	ret, err := i.repo.BookCar(ctx, from, target)
	if err != nil {
		ctx.Error(errorx.ErrBookRental.Error(), zap.Error(err), zap.Any("from", from), zap.Any("target", target))
		return nil, errorx.ErrBookRental.ReplaceMsg(err.Error())
	}

	return ret, nil
}

func (i *impl) CancelBooking(ctx contextx.Contextx, from *am.Profile, target *om.Booking) error {
	if len(from.AccessToken) == 0 {
		ctx.Error(errorx.ErrMissingToken.Error(), zap.Any("from", from))
		return errorx.ErrMissingToken
	}

	if len(target.No) == 0 {
		ctx.Error(errorx.ErrInvalidBooking.Error(), zap.Any("target", target))
		return errorx.ErrInvalidBooking
	}

	err := i.repo.CancelBooking(ctx, from, target)
	if err != nil {
		ctx.Error(errorx.ErrCancelBooking.Error(), zap.Error(err), zap.Any("from", from), zap.Any("target", target))
		return errorx.ErrCancelBooking.ReplaceMsg(err.Error())
	}

	return nil
}
