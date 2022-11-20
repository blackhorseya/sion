//go:generate wire
//go:build wireinject

package repo

import (
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewImpl)

func CreateRepo() IRepo {
	panic(wire.Build(testProviderSet))
}
