package repo

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"testing"

	"github.com/blackhorseya/irent/pkg/contextx"
	am "github.com/blackhorseya/irent/pkg/entity/domain/account/model"
	"github.com/blackhorseya/irent/pkg/httpx"
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
			name: "do request then error",
			args: args{id: "id", password: "password", mock: func() {
				s.httpclient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("error")).Times(1)
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "got error message then error",
			args: args{id: "id", password: "password", mock: func() {
				payload, _ := json.Marshal(&loginResp{
					ErrorMessage: "error",
				})
				body := io.NopCloser(bytes.NewReader(payload))
				s.httpclient.EXPECT().Do(gomock.Any()).Return(&http.Response{
					StatusCode: http.StatusOK,
					Body:       body,
				}, nil).Times(1)
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "ok",
			args: args{id: "id", password: "password", mock: func() {
				payload, _ := json.Marshal(&loginResp{
					Result:       "",
					ErrorCode:    "",
					NeedRelogin:  0,
					NeedUpgrade:  0,
					ErrorMessage: "Success",
				})
				body := io.NopCloser(bytes.NewReader(payload))
				s.httpclient.EXPECT().Do(gomock.Any()).Return(&http.Response{
					StatusCode: http.StatusOK,
					Body:       body,
				}, nil).Times(1)
			}},
			wantInfo: &am.Profile{},
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfo, err := s.repo.Login(contextx.BackgroundWithLogger(s.logger), tt.args.id, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("Login() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}
		})
	}
}

func (s *SuiteTester) Test_impl_GetMemberStatus() {
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
			name: "do request then error",
			args: args{token: "token", mock: func() {
				s.httpclient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("error")).Times(1)
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "got error message then error",
			args: args{token: "token", mock: func() {
				payload, _ := json.Marshal(&getMemberStatusResp{
					ErrorMessage: "error",
				})
				body := io.NopCloser(bytes.NewReader(payload))
				s.httpclient.EXPECT().Do(gomock.Any()).Return(&http.Response{
					StatusCode: http.StatusOK,
					Body:       body,
				}, nil).Times(1)
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "ok",
			args: args{token: "token", mock: func() {
				payload, _ := json.Marshal(&getMemberStatusResp{
					ErrorMessage: "Success",
				})
				body := io.NopCloser(bytes.NewReader(payload))
				s.httpclient.EXPECT().Do(gomock.Any()).Return(&http.Response{
					StatusCode: http.StatusOK,
					Body:       body,
				}, nil).Times(1)
			}},
			wantInfo: &am.Profile{AccessToken: "token"},
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfo, err := s.repo.GetMemberStatus(contextx.BackgroundWithLogger(s.logger), tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMemberStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("GetMemberStatus() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}
		})
	}
}
