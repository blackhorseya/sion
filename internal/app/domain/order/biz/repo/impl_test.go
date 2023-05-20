package repo

import (
	"reflect"
	"testing"

	"github.com/blackhorseya/irent/pkg/contextx"
	am "github.com/blackhorseya/irent/pkg/entity/domain/account/model"
	om "github.com/blackhorseya/irent/pkg/entity/domain/order/model"
	rm "github.com/blackhorseya/irent/pkg/entity/domain/rental/model"
	"github.com/blackhorseya/irent/pkg/httpx"
	"github.com/blackhorseya/irent/test/testdata"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type SuiteTester struct {
	suite.Suite
	logger     *zap.Logger
	ctrl       *gomock.Controller
	httpclient *httpx.MockClient
	repo       IRepo
}

func (s *SuiteTester) SetupTest() {
	s.logger, _ = zap.NewDevelopment()
	s.ctrl = gomock.NewController(s.T())
	s.httpclient = httpx.NewMockClient(s.ctrl)
	opts := &Options{Endpoint: "http://localhost", AppVersion: "1.0.0"}
	s.repo = CreateRepo(opts, s.httpclient)
}

func (s *SuiteTester) TearDownTest() {
	s.ctrl.Finish()
}

func TestAll(t *testing.T) {
	suite.Run(t, new(SuiteTester))
}

func (s *SuiteTester) Test_impl_FetchArrears() {
	type args struct {
		from   *am.Profile
		target *am.Profile
		mock   func()
	}
	tests := []struct {
		name        string
		args        args
		wantRecords []*om.ArrearsRecord
		wantErr     bool
	}{
		{
			name: "http do return error",
			args: args{from: testdata.Profile1, target: testdata.Profile1, mock: func() {
				s.httpclient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("error")).Times(1)
			}},
			wantRecords: nil,
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotRecords, err := s.repo.FetchArrears(contextx.BackgroundWithLogger(s.logger), tt.args.from, tt.args.target)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchArrears() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRecords, tt.wantRecords) {
				t.Errorf("FetchArrears() gotRecords = %v, want %v", gotRecords, tt.wantRecords)
			}
		})
	}
}

func (s *SuiteTester) Test_impl_BookCar() {
	type args struct {
		from   *am.Profile
		target *rm.Car
		mock   func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *om.Lease
		wantErr  bool
	}{
		{
			name: "http do then error",
			args: args{from: testdata.Profile1, target: testdata.Car1, mock: func() {
				s.httpclient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("error")).Times(1)
			}},
			wantInfo: nil,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfo, err := s.repo.BookCar(contextx.BackgroundWithLogger(s.logger), tt.args.from, tt.args.target)
			if (err != nil) != tt.wantErr {
				t.Errorf("BookCar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("BookCar() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}
		})
	}
}

func (s *SuiteTester) Test_impl_CancelBooking() {
	type args struct {
		from   *am.Profile
		target *om.Lease
		mock   func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "http do then error",
			args: args{from: testdata.Profile1, target: testdata.Lease1, mock: func() {
				s.httpclient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("error")).Times(1)
			}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			if err := s.repo.CancelBooking(contextx.BackgroundWithLogger(s.logger), tt.args.from, tt.args.target); (err != nil) != tt.wantErr {
				t.Errorf("CancelBooking() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
