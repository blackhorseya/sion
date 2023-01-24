package biz

import (
	"github.com/blackhorseya/irent/internal/app/domain/order/biz/repo"
	"github.com/blackhorseya/irent/internal/pkg/errorx"
	"github.com/blackhorseya/irent/pkg/contextx"
	am "github.com/blackhorseya/irent/pkg/entity/domain/account/model"
	ob "github.com/blackhorseya/irent/pkg/entity/domain/order/biz"
	om "github.com/blackhorseya/irent/pkg/entity/domain/order/model"
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

func (i *impl) GetArrears(ctx contextx.Contextx, from *am.Profile) (info *om.Arrears, err error) {
	records, err := i.repo.FetchArrears(ctx, from)
	if err != nil {
		ctx.Error(errorx.ErrGetArrears.Error(), zap.Error(err), zap.Any("from", from))
		return nil, errorx.ErrGetArrears
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
