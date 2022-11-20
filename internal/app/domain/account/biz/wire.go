//go:generate wire
//go:build wireinject

package biz

import (
	"github.com/blackhorseya/irent/internal/app/domain/account/biz/repo"
	ab "github.com/blackhorseya/irent/pkg/entity/domain/account/biz"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewImpl)

func CreateBiz(repo repo.IRepo) ab.IBiz {
	panic(wire.Build(testProviderSet))
}
