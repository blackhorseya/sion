//go:generate wire
//go:build wireinject

package biz

import (
	"github.com/blackhorseya/irent/internal/app/domain/order/biz/repo"
	ob "github.com/blackhorseya/irent/pkg/entity/domain/order/biz"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewImpl)

func CreateBiz(repo repo.IRepo) ob.IBiz {
	panic(wire.Build(testProviderSet))
}
