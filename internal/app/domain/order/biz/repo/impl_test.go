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
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type SuiteTester struct {
	suite.Suite
	logger     *zap.Logger
	httpclient *httpx.MockClient
	repo       IRepo
}

func (s *SuiteTester) SetupTest() {
	s.logger, _ = zap.NewDevelopment()
	s.httpclient = new(httpx.MockClient)
	opts := &Options{Endpoint: "http://localhost", AppVersion: "1.0.0"}
	s.repo = CreateRepo(opts, s.httpclient)
}

func (s *SuiteTester) AssertExpectations(t *testing.T) {
	s.httpclient.AssertExpectations(t)
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
				s.httpclient.On("Do", mock.Anything).Return(nil, errors.New("error")).Once()
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

			s.AssertExpectations(t)
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
		wantInfo *om.Booking
		wantErr  bool
	}{
		{
			name: "http do then error",
			args: args{from: testdata.Profile1, target: testdata.Car1, mock: func() {
				s.httpclient.On("Do", mock.Anything).Return(nil, errors.New("error")).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		// TODO: Add test cases.
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

			s.AssertExpectations(t)
		})
	}
}
