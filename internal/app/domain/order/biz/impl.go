package biz

import (
	"github.com/blackhorseya/irent/internal/app/domain/order/biz/repo"
	"github.com/blackhorseya/irent/pkg/contextx"
	am "github.com/blackhorseya/irent/pkg/entity/domain/account/model"
	ob "github.com/blackhorseya/irent/pkg/entity/domain/order/biz"
	om "github.com/blackhorseya/irent/pkg/entity/domain/order/model"
)

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
	// todo: 2023/1/23|sean|impl me
	panic("implement me")
}
