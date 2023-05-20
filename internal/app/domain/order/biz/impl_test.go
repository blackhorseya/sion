package biz

import (
	"reflect"
	"testing"

	"github.com/blackhorseya/irent/internal/app/domain/order/biz/repo"
	"github.com/blackhorseya/irent/pkg/contextx"
	am "github.com/blackhorseya/irent/pkg/entity/domain/account/model"
	ob "github.com/blackhorseya/irent/pkg/entity/domain/order/biz"
	om "github.com/blackhorseya/irent/pkg/entity/domain/order/model"
	rm "github.com/blackhorseya/irent/pkg/entity/domain/rental/model"
	"github.com/blackhorseya/irent/test/testdata"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type SuiteTester struct {
	suite.Suite
	logger *zap.Logger
	ctrl   *gomock.Controller
	repo   *repo.MockIRepo
	biz    ob.IBiz
}

func (s *SuiteTester) SetupTest() {
	s.logger, _ = zap.NewDevelopment()
	s.ctrl = gomock.NewController(s.T())
	s.repo = repo.NewMockIRepo(s.ctrl)
	s.biz = CreateBiz(s.repo)
}

func (s *SuiteTester) TearDownTest() {
	s.ctrl.Finish()
}

func TestAll(t *testing.T) {
	suite.Run(t, new(SuiteTester))
}

func (s *SuiteTester) Test_impl_GetArrears() {
	type args struct {
		from   *am.Profile
		target *am.Profile
		mock   func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *om.Arrears
		wantErr  bool
	}{
		{
			name:     "check token of from then error",
			args:     args{from: &am.Profile{AccessToken: ""}, target: testdata.Profile1},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name:     "check id of target then error",
			args:     args{from: testdata.Profile1, target: &am.Profile{Id: ""}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "fetch then error",
			args: args{from: testdata.Profile1, target: testdata.Profile1, mock: func() {
				s.repo.EXPECT().FetchArrears(gomock.Any(), testdata.Profile1, testdata.Profile1).Return(nil, errors.New("error")).Times(1)
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "ok",
			args: args{from: testdata.Profile1, target: testdata.Profile1, mock: func() {
				s.repo.EXPECT().FetchArrears(gomock.Any(), testdata.Profile1, testdata.Profile1).Return(nil, nil).Times(1)
			}},
			wantInfo: nil,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfo, err := s.biz.GetArrears(contextx.BackgroundWithLogger(s.logger), tt.args.from, tt.args.target)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetArrears() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("GetArrears() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}
		})
	}
}

func (s *SuiteTester) Test_impl_BookRental() {
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
			name:     "check token of from then error",
			args:     args{from: &am.Profile{AccessToken: ""}, target: testdata.Car1},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name:     "check target then error",
			args:     args{from: testdata.Profile1, target: &rm.Car{Id: "", ProjectId: ""}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "book then error",
			args: args{from: testdata.Profile1, target: testdata.Car1, mock: func() {
				s.repo.EXPECT().BookCar(gomock.Any(), testdata.Profile1, testdata.Car1).Return(nil, errors.New("error")).Times(1)
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "ok",
			args: args{from: testdata.Profile1, target: testdata.Car1, mock: func() {
				s.repo.EXPECT().BookCar(gomock.Any(), testdata.Profile1, testdata.Car1).Return(nil, nil).Times(1)
			}},
			wantInfo: nil,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfo, err := s.biz.BookRental(contextx.BackgroundWithLogger(s.logger), tt.args.from, tt.args.target)
			if (err != nil) != tt.wantErr {
				t.Errorf("BookRental() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("BookRental() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}
		})
	}
}

func (s *SuiteTester) Test_impl_CancelLease() {
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
			name:    "check from then error",
			args:    args{from: &am.Profile{AccessToken: ""}, target: testdata.Lease1},
			wantErr: true,
		},
		{
			name:    "check target then error",
			args:    args{from: testdata.Profile1, target: &om.Lease{No: ""}},
			wantErr: true,
		},
		{
			name: "cancel then error",
			args: args{from: testdata.Profile1, target: testdata.Lease1, mock: func() {
				s.repo.EXPECT().CancelBooking(gomock.Any(), testdata.Profile1, testdata.Lease1).Return(errors.New("error")).Times(1)
			}},
			wantErr: true,
		},
		{
			name: "ok",
			args: args{from: testdata.Profile1, target: testdata.Lease1, mock: func() {
				s.repo.EXPECT().CancelBooking(gomock.Any(), testdata.Profile1, testdata.Lease1).Return(nil).Times(1)
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			if err := s.biz.CancelLease(contextx.BackgroundWithLogger(s.logger), tt.args.from, tt.args.target); (err != nil) != tt.wantErr {
				t.Errorf("CancelLease() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func (s *SuiteTester) Test_impl_ListLease() {
	type args struct {
		from *am.Profile
		mock func()
	}
	tests := []struct {
		name       string
		args       args
		wantOrders []*om.Lease
		wantErr    bool
	}{
		{
			name:       "check from then error",
			args:       args{from: &am.Profile{AccessToken: ""}},
			wantOrders: nil,
			wantErr:    true,
		},
		{
			name: "query bookings then error",
			args: args{from: testdata.Profile1, mock: func() {
				s.repo.EXPECT().QueryBookings(gomock.Any(), testdata.Profile1).Return(nil, errors.New("error")).Times(1)
			}},
			wantOrders: nil,
			wantErr:    true,
		},
		{
			name: "ok",
			args: args{from: testdata.Profile1, mock: func() {
				s.repo.EXPECT().QueryBookings(gomock.Any(), testdata.Profile1).Return(nil, nil).Times(1)
			}},
			wantOrders: nil,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotOrders, err := s.biz.ListLease(contextx.BackgroundWithLogger(s.logger), tt.args.from)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListLease() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOrders, tt.wantOrders) {
				t.Errorf("ListLease() gotOrders = %v, want %v", gotOrders, tt.wantOrders)
			}
		})
	}
}
