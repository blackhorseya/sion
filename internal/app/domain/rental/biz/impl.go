package biz

import (
	"github.com/blackhorseya/irent/pkg/contextx"
	rb "github.com/blackhorseya/irent/pkg/entity/domain/rental/biz"
	"github.com/blackhorseya/irent/pkg/entity/domain/rental/model"
)

type impl struct {
}

func NewImpl() rb.IBiz {
	return &impl{}
}

func (i *impl) Readiness(ctx contextx.Contextx) error {
	return nil
}

func (i *impl) Liveness(ctx contextx.Contextx) error {
	return nil
}

func (i *impl) ListCar(ctx contextx.Contextx, condition rb.QueryCarCondition) (info []*model.Car, total int, err error) {
	// TODO implement me
	panic("implement me")
}
