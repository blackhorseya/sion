//go:generate wire
//go:build wireinject

package repo

import (
	"github.com/blackhorseya/irent/pkg/httpx"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewImpl)

func CreateRepo(opts *Options, httpclient httpx.Client) IRepo {
	panic(wire.Build(testProviderSet))
}
