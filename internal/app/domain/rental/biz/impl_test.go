package biz

import (
	"reflect"
	"testing"

	"github.com/blackhorseya/irent/internal/app/domain/rental/biz/repo"
	"github.com/blackhorseya/irent/pkg/contextx"
	rb "github.com/blackhorseya/irent/pkg/entity/domain/rental/biz"
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
	biz    rb.IBiz
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

func (s *SuiteTester) Test_impl_ListCars() {
	type args struct {
		condition rb.QueryCarCondition
		mock      func()
	}
	tests := []struct {
		name      string
		args      args
		wantInfo  []*rm.Car
		wantTotal int
		wantErr   bool
	}{
		{
			name: "invalid top num then error",
			args: args{condition: rb.QueryCarCondition{TopNum: 10}, mock: func() {
				s.repo.EXPECT().ListCars(gomock.Any()).Return(nil, errors.New("error")).Times(1)
			}},
			wantInfo:  nil,
			wantTotal: 0,
			wantErr:   true,
		},
		{
			name: "not found any cars then nil",
			args: args{condition: rb.QueryCarCondition{TopNum: 10}, mock: func() {
				s.repo.EXPECT().ListCars(gomock.Any()).Return(nil, nil).Times(1)
			}},
			wantInfo:  nil,
			wantTotal: 0,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfo, gotTotal, err := s.biz.ListCars(contextx.BackgroundWithLogger(s.logger), tt.args.condition)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListCars() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("ListCars() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}
			if gotTotal != tt.wantTotal {
				t.Errorf("ListCars() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}
		})
	}
}

func (s *SuiteTester) Test_impl_UpdateInfoCars() {
	type args struct {
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantCars []*rm.Car
		wantErr  bool
	}{
		{
			name: "fetch all cars then error",
			args: args{mock: func() {
				s.repo.EXPECT().FetchAvailableCars(gomock.Any()).Return(nil, errors.New("error")).Times(1)
			}},
			wantCars: nil,
			wantErr:  true,
		},
		{
			name: "update status all cars then error",
			args: args{mock: func() {
				s.repo.EXPECT().FetchAvailableCars(gomock.Any()).Return(nil, nil).Times(1)

				s.repo.EXPECT().UpdateStatusAllCars(gomock.Any(), rm.CarStatus_CAR_STATUS_INUSE).Return(errors.New("error")).Times(1)
			}},
			wantCars: nil,
			wantErr:  true,
		},
		{
			name: "for loop each car to upsert then ok",
			args: args{mock: func() {
				s.repo.EXPECT().FetchAvailableCars(gomock.Any()).Return([]*rm.Car{testdata.Car1}, nil).Times(1)

				s.repo.EXPECT().UpdateStatusAllCars(gomock.Any(), rm.CarStatus_CAR_STATUS_INUSE).Return(nil).Times(1)

				s.repo.EXPECT().UpsertStatusCar(gomock.Any(), testdata.Car1).Return(errors.New("error")).Times(1)
			}},
			wantCars: []*rm.Car{testdata.Car1},
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotCars, err := s.biz.UpdateInfoCars(contextx.BackgroundWithLogger(s.logger))
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateInfoCars() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCars, tt.wantCars) {
				t.Errorf("UpdateInfoCars() gotCars = %v, want %v", gotCars, tt.wantCars)
			}
		})
	}
}
