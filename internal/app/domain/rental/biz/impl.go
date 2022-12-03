package biz

import (
	"github.com/blackhorseya/irent/internal/app/domain/rental/biz/repo"
	"github.com/blackhorseya/irent/pkg/contextx"
	rb "github.com/blackhorseya/irent/pkg/entity/domain/rental/biz"
	rm "github.com/blackhorseya/irent/pkg/entity/domain/rental/model"
	"github.com/google/wire"
)

// ProviderSet is a provider set for rental biz
var ProviderSet = wire.NewSet(NewImpl, repo.ProviderSet)

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

func (i *impl) ListCar(ctx contextx.Contextx, condition rb.QueryCarCondition) (info []*rm.Car, total int, err error) {
	// TODO implement me
	panic("implement me")
}
