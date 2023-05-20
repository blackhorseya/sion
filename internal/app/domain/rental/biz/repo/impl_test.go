package repo

import (
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/blackhorseya/irent/pkg/contextx"
	"github.com/blackhorseya/irent/pkg/entity/domain/rental/model"
	"github.com/blackhorseya/irent/pkg/httpx"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type SuiteTester struct {
	suite.Suite
	logger     *zap.Logger
	ctrl       *gomock.Controller
	httpclient *httpx.MockClient
	rw         sqlmock.Sqlmock
	repo       IRepo
}

func (s *SuiteTester) SetupTest() {
	s.logger, _ = zap.NewDevelopment()
	s.ctrl = gomock.NewController(s.T())
	s.httpclient = httpx.NewMockClient(s.ctrl)
	opts := &Options{Endpoint: "http://localhost", AppVersion: "1.0.0"}
	db, rw, _ := sqlmock.New(sqlmock.MonitorPingsOption(true))
	s.rw = rw
	s.repo = CreateRepo(opts, s.httpclient, sqlx.NewDb(db, "mysql"))
}

func (s *SuiteTester) TearDownTest() {
	s.ctrl.Finish()
}

func (s *SuiteTester) assert(t *testing.T) {
	if err := s.rw.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
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

			gotInfo, err := s.repo.ListCars(contextx.BackgroundWithLogger(s.logger))
			if (err != nil) != tt.wantErr {
				t.Errorf("ListCars() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("ListCars() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}

			s.assert(t)
		})
	}
}
