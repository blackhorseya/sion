package biz

import (
	"reflect"
	"testing"

	"github.com/blackhorseya/irent/internal/app/domain/account/biz/repo"
	"github.com/blackhorseya/irent/pkg/contextx"
	ab "github.com/blackhorseya/irent/pkg/entity/domain/account/biz"
	am "github.com/blackhorseya/irent/pkg/entity/domain/account/model"
	"github.com/blackhorseya/irent/test/testdata"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type SuiteTester struct {
	suite.Suite
	logger *zap.Logger
	repo   *repo.MockIRepo
	biz    ab.IBiz
}

func (s *SuiteTester) SetupTest() {
	s.logger, _ = zap.NewDevelopment()
	s.repo = new(repo.MockIRepo)
	s.biz = CreateBiz(s.repo)
}

func (s *SuiteTester) TearDownTest() {
	s.repo.AssertExpectations(s.T())
}

func TestAll(t *testing.T) {
	suite.Run(t, new(SuiteTester))
}

func (s *SuiteTester) Test_impl_Readiness() {
	type args struct {
		mock func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "ok",
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			if err := s.biz.Readiness(contextx.BackgroundWithLogger(s.logger)); (err != nil) != tt.wantErr {
				t.Errorf("Readiness() error = %v, wantErr %v", err, tt.wantErr)
			}

			s.TearDownTest()
		})
	}
}

func (s *SuiteTester) Test_impl_Liveness() {
	type args struct {
		mock func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "ok",
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			if err := s.biz.Liveness(contextx.BackgroundWithLogger(s.logger)); (err != nil) != tt.wantErr {
				t.Errorf("Liveness() error = %v, wantErr %v", err, tt.wantErr)
			}

			s.TearDownTest()
		})
	}
}

func (s *SuiteTester) Test_impl_Login() {
	type args struct {
		id       string
		password string
		mock     func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *am.Profile
		wantErr  bool
	}{
		{
			name:     "id is invalid then error",
			args:     args{id: "", password: "password"},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name:     "password is invalid then error",
			args:     args{id: "id", password: ""},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "login then error",
			args: args{id: "id", password: "password", mock: func() {
				s.repo.On("Login", mock.Anything, "id", encryptPassword("password")).Return(nil, errors.New("error")).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "ok",
			args: args{id: "id", password: "password", mock: func() {
				s.repo.On("Login", mock.Anything, "id", encryptPassword("password")).Return(testdata.Profile1, nil).Once()
			}},
			wantInfo: testdata.Profile1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfo, err := s.biz.Login(contextx.BackgroundWithLogger(s.logger), tt.args.id, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("Login() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}

			s.TearDownTest()
		})
	}
}

func (s *SuiteTester) Test_impl_GetByAccessToken() {
	type args struct {
		token string
		mock  func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *am.Profile
		wantErr  bool
	}{
		{
			name:     "token is empty then error",
			args:     args{token: ""},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "get member status then error",
			args: args{token: "token", mock: func() {
				s.repo.On("GetMemberStatus", mock.Anything, "token").Return(nil, errors.New("error")).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "ok",
			args: args{token: "token", mock: func() {
				s.repo.On("GetMemberStatus", mock.Anything, "token").Return(testdata.Profile1, nil).Once()
			}},
			wantInfo: testdata.Profile1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfo, err := s.biz.GetByAccessToken(contextx.BackgroundWithLogger(s.logger), tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("GetByAccessToken() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}

			s.TearDownTest()
		})
	}
}
