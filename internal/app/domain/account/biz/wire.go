//go:generate wire
//go:build wireinject

package biz

import (
	ab "github.com/blackhorseya/irent/pkg/entity/domain/account/biz"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewImpl)

func CreateBiz() ab.IBiz {
	panic(wire.Build(testProviderSet))
}
