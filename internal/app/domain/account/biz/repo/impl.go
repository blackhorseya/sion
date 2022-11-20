package repo

import (
	"github.com/blackhorseya/irent/pkg/contextx"
	"github.com/blackhorseya/irent/pkg/entity/domain/account/model"
)

type impl struct {
	// todo: 2022/11/20|sean|breakpoint: inject httpx client
}

func NewImpl() IRepo {
	return &impl{}
}

func (i *impl) Login(ctx contextx.Contextx, id, password string) (info *model.Profile, err error) {
	// TODO implement me
	panic("implement me")
}

func (i *impl) GetMemberStatus(ctx contextx.Contextx, token string) (info *model.Profile, err error) {
	// TODO implement me
	panic("implement me")
}
