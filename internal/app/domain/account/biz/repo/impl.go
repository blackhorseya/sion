package repo

import (
	"github.com/blackhorseya/irent/pkg/contextx"
	"github.com/blackhorseya/irent/pkg/entity/domain/account/model"
	"github.com/blackhorseya/irent/pkg/httpx"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// ProviderSet is a provider set for account repo
var ProviderSet = wire.NewSet(NewOptions, NewImpl)

// Options declare app's configuration
type Options struct {
	Endpoint   string `json:"endpoint" yaml:"endpoint"`
	AppVersion string `json:"app_version" yaml:"appVersion"`
}

func NewOptions(v *viper.Viper) (*Options, error) {
	ret := new(Options)

	err := v.UnmarshalKey("app", ret)
	if err != nil {
		return nil, errors.Wrap(err, "load app options failed")
	}

	return ret, nil
}

type impl struct {
	opts       *Options
	httpclient httpx.Client
}

func NewImpl(opts *Options, httpclient httpx.Client) IRepo {
	return &impl{
		opts:       opts,
		httpclient: httpclient,
	}
}

func (i *impl) Login(ctx contextx.Contextx, id, password string) (info *model.Profile, err error) {
	// TODO implement me
	panic("implement me")
}

func (i *impl) GetMemberStatus(ctx contextx.Contextx, token string) (info *model.Profile, err error) {
	// TODO implement me
	panic("implement me")
}
