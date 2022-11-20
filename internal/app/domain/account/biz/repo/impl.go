package repo

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/blackhorseya/irent/pkg/contextx"
	am "github.com/blackhorseya/irent/pkg/entity/domain/account/model"
	"github.com/blackhorseya/irent/pkg/httpx"
	"github.com/google/uuid"
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

func (i *impl) Login(ctx contextx.Contextx, id, password string) (info *am.Profile, err error) {
	uri, err := url.Parse(i.opts.Endpoint + "/Login")
	if err != nil {
		return nil, err
	}

	payload, err := json.Marshal(map[string]interface{}{
		"IDNO":       id,
		"DeviceID":   uuid.New().String(),
		"app":        "1",
		"appVersion": i.opts.AppVersion,
		"PWD":        password,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, uri.String(), bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := i.httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var body *loginResp
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	if body.ErrorMessage != "Success" {
		return nil, errors.New(body.ErrorMessage)
	}

	return &am.Profile{
		Id:          body.Data.UserData.Memidno,
		Name:        body.Data.UserData.Memcname,
		AccessToken: body.Data.Token.AccessToken,
	}, nil
}

func (i *impl) GetMemberStatus(ctx contextx.Contextx, token string) (info *am.Profile, err error) {
	// TODO implement me
	panic("implement me")
}
