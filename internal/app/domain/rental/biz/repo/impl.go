package repo

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/blackhorseya/irent/pkg/contextx"
	rm "github.com/blackhorseya/irent/pkg/entity/domain/rental/model"
	"github.com/blackhorseya/irent/pkg/httpx"
	"github.com/jmoiron/sqlx"
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
	rw         *sqlx.DB
}

func NewImpl(opts *Options, httpclient httpx.Client, rw *sqlx.DB) IRepo {
	return &impl{
		opts:       opts,
		httpclient: httpclient,
		rw:         rw,
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
			Status:      rm.CarStatus_CAR_STATUS_AVAILABLE,
		}
	}

	return ret, nil
}

func (i *impl) FetchAvailableCars(ctx contextx.Contextx) (cars []*rm.Car, err error) {
	defer ctx.Elapsed("[FetchAvailableCars]")()

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
			Status:      rm.CarStatus_CAR_STATUS_AVAILABLE,
		}
	}

	return ret, nil
}

func (i *impl) UpsertStatusCar(ctx contextx.Contextx, target *rm.Car) error {
	timeout, cancelFunc := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancelFunc()

	stmt := `insert into cars (id, area, project_id, project_name, seat, type_name, latitude, longitude, status)
values (:id, :area, :project_id, :project_name, :seat, :type_name, :latitude, :longitude, :status)
on duplicate key update status = :status`

	arg := newCar(target)
	_, err := i.rw.NamedExecContext(timeout, stmt, arg)
	if err != nil {
		return err
	}

	return nil
}

func (i *impl) UpdateStatusAllCars(ctx contextx.Contextx, status rm.CarStatus) error {
	defer ctx.Elapsed("[UpdateStatusAllCars]")()

	timeout, cancelFunc := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancelFunc()

	stmt := `update cars set status=?`

	_, err := i.rw.ExecContext(timeout, stmt, int32(status))
	if err != nil {
		return err
	}

	return nil
}
