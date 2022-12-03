package repo

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/blackhorseya/irent/pkg/contextx"
	rm "github.com/blackhorseya/irent/pkg/entity/domain/rental/model"
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

func (i *impl) ListCars(ctx contextx.Contextx) (info []*rm.Car, err error) {
	uri, err := url.Parse(i.opts.Endpoint + "/AnyRent")
	if err != nil {
		return nil, err
	}

	payload, err := json.Marshal(map[string]interface{}{
		"Radius":    0,
		"Latitude":  0,
		"Longitude": 0,
		"ShowAll":   1,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, uri.String(), bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}

	resp, err := i.httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var got *anyRentResponse
	err = json.NewDecoder(resp.Body).Decode(&got)
	if err != nil {
		return nil, err
	}

	if got.ErrorMessage != "Success" {
		return nil, errors.New(got.ErrorMessage)
	}

	ret := make([]*rm.Car, len(got.Data.AnyRentObj))
	for idx, obj := range got.Data.AnyRentObj {
		ret[idx] = &rm.Car{
			Id:          strings.ReplaceAll(obj.CarNo, " ", ""),
			CarType:     obj.CarType,
			CarTypeName: obj.CarTypeName,
			CarOfArea:   obj.CarOfArea,
			ProjectName: obj.ProjectName,
			ProjectId:   obj.ProjID,
			Latitude:    obj.Latitude,
			Longitude:   obj.Longitude,
			Seat:        obj.Seat,
			Distance:    0,
		}
	}

	return ret, nil
}
