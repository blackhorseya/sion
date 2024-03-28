//go:build wireinject

//go:generate wire

package restful

import (
	"github.com/blackhorseya/sion/app/domain/rental/biz"
	"github.com/blackhorseya/sion/pkg/adapterx"
	"github.com/blackhorseya/sion/pkg/linebotx"
	"github.com/blackhorseya/sion/pkg/transports/httpx"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var providerSet = wire.NewSet(
	httpx.NewServer,
	linebotx.NewClient,
	biz.ProviderSet,
)

func New(v *viper.Viper) (adapterx.Servicer, error) {
	panic(wire.Build(newService, providerSet))
}

func NewRestful() (adapterx.Restful, error) {
	panic(wire.Build(newRestful, providerSet))
}
