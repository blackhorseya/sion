package repo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/blackhorseya/irent/pkg/contextx"
	am "github.com/blackhorseya/irent/pkg/entity/domain/account/model"
	om "github.com/blackhorseya/irent/pkg/entity/domain/order/model"
	"github.com/blackhorseya/irent/pkg/httpx"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

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

func (i *impl) FetchArrears(ctx contextx.Contextx, user *am.Profile) (records []*om.ArrearsRecord, err error) {
	uri, err := url.Parse(i.opts.Endpoint + "/ArrearsQuery")
	if err != nil {
		return nil, err
	}

	payload, err := json.Marshal(map[string]interface{}{
		"IDNO": user.Id,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, uri.String(), bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", user.AccessToken))

	resp, err := i.httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data *fetchArrearsResp
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	if strings.ToLower(data.ErrorMessage) != "success" {
		return nil, errors.New(data.ErrorMessage)
	}

	var ret []*om.ArrearsRecord
	for _, info := range data.Data.ArrearsInfos {
		ret = append(ret, &om.ArrearsRecord{
			OrderNo:     info.OrderNo,
			TotalAmount: int64(info.TotalAmount),
		})
	}

	return ret, nil
}
