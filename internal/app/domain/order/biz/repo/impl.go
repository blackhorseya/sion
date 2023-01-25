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
	rm "github.com/blackhorseya/irent/pkg/entity/domain/rental/model"
	"github.com/blackhorseya/irent/pkg/httpx"
	"github.com/blackhorseya/irent/pkg/timex"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const _msg_success = "success"

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

func (i *impl) FetchArrears(ctx contextx.Contextx, from *am.Profile, target *am.Profile) (records []*om.ArrearsRecord, err error) {
	uri, err := url.Parse(i.opts.Endpoint + "/ArrearsQuery")
	if err != nil {
		return nil, err
	}

	payload, err := json.Marshal(map[string]interface{}{
		"IDNO": target.Id,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, uri.String(), bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", from.AccessToken))

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

	if strings.ToLower(data.ErrorMessage) != _msg_success {
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

func (i *impl) BookCar(ctx contextx.Contextx, from *am.Profile, target *rm.Car) (info *om.Lease, err error) {
	uri, err := url.Parse(i.opts.Endpoint + "/Booking")
	if err != nil {
		return nil, err
	}

	payload, err := json.Marshal(map[string]interface{}{
		"ProjID": target.ProjectId,
		"EDate":  "",
		"SDate":  "",
		"CarNo":  target.Id,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, uri.String(), bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", from.AccessToken))

	resp, err := i.httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data *bookResp
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	if strings.ToLower(data.ErrorMessage) != _msg_success {
		return nil, errors.New(data.ErrorMessage)
	}

	t, err := timex.ParseYYYYMMddHHmmss(data.Data.LastPickTime)
	if err != nil {
		return nil, err
	}

	return &om.Lease{
		No:           data.Data.OrderNo,
		CarId:        target.Id,
		CarLatitude:  target.Latitude,
		CarLongitude: target.Longitude,
		StartAt:      nil,
		EndAt:        nil,
		LastPickAt:   timestamppb.New(t),
	}, nil
}

func (i *impl) CancelBooking(ctx contextx.Contextx, from *am.Profile, target *om.Lease) error {
	uri, err := url.Parse(i.opts.Endpoint + "/BookingCancel")
	if err != nil {
		return err
	}

	payload, err := json.Marshal(map[string]interface{}{
		"OrderNo": target.No,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, uri.String(), bytes.NewReader(payload))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", from.AccessToken))

	resp, err := i.httpclient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var data *cancelBookingResp
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return err
	}

	if strings.ToLower(data.ErrorMessage) != _msg_success {
		return errors.New(data.ErrorMessage)
	}

	return nil
}
