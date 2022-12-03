//go:generate wire
//go:build wireinject

package biz

import (
	"github.com/blackhorseya/irent/internal/app/domain/rental/biz/repo"
	rb "github.com/blackhorseya/irent/pkg/entity/domain/rental/biz"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewImpl)

func CreateBiz(repo repo.IRepo) rb.IBiz {
	panic(wire.Build(testProviderSet))
}
