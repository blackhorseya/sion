package biz

import (
	"github.com/blackhorseya/irent/pkg/contextx"
	ab "github.com/blackhorseya/irent/pkg/entity/domain/account/biz"
	"github.com/blackhorseya/irent/pkg/entity/domain/account/model"
)

type impl struct {
}

func NewImpl() ab.IBiz {
	return &impl{}
}

func (i *impl) Readiness(ctx contextx.Contextx) error {
	return nil
}

func (i *impl) Liveness(ctx contextx.Contextx) error {
	return nil
}

func (i *impl) Login(ctx contextx.Contextx, id, password string) (info *model.Profile, err error) {
	// TODO implement me
	panic("implement me")
}

func (i *impl) GetByAccessToken(ctx contextx.Contextx, token string) (info *model.Profile, err error) {
	// TODO implement me
	panic("implement me")
}
