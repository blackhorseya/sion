package repo

import (
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/blackhorseya/irent/pkg/contextx"
	"github.com/blackhorseya/irent/pkg/entity/domain/rental/model"
	"github.com/blackhorseya/irent/pkg/httpx"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type SuiteTester struct {
	suite.Suite
	logger     *zap.Logger
	httpclient *httpx.MockClient
	rw         sqlmock.Sqlmock
	repo       IRepo
}

func (s *SuiteTester) SetupTest() {
	s.logger, _ = zap.NewDevelopment()
	s.httpclient = new(httpx.MockClient)
	opts := &Options{Endpoint: "http://localhost", AppVersion: "1.0.0"}
	db, rw, _ := sqlmock.New(sqlmock.MonitorPingsOption(true))
	s.rw = rw
	s.repo = CreateRepo(opts, s.httpclient, sqlx.NewDb(db, "mysql"))
}

func (s *SuiteTester) assertExpectation(t *testing.T) {
	if err := s.rw.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	s.httpclient.AssertExpectations(t)
}

func TestAll(t *testing.T) {
	suite.Run(t, new(SuiteTester))
}

func (s *SuiteTester) Test_impl_ListCars() {
	type args struct {
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo []*model.Car
		wantErr  bool
	}{
		{
			name: "do request then error",
			args: args{mock: func() {
				s.httpclient.On("Do", mock.Anything).Return(nil, errors.New("error")).Once()
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

			gotInfo, err := s.repo.ListCars(contextx.BackgroundWithLogger(s.logger))
			if (err != nil) != tt.wantErr {
				t.Errorf("ListCars() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("ListCars() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}

			s.assertExpectation(t)
		})
	}
}
