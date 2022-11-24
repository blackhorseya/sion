package auth

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/blackhorseya/irent/pkg/contextx"
	ab "github.com/blackhorseya/irent/pkg/entity/domain/account/biz"
	"github.com/blackhorseya/irent/pkg/er"
	"github.com/blackhorseya/irent/pkg/response"
	"github.com/blackhorseya/irent/test/testdata"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type SuiteTester struct {
	suite.Suite
	logger *zap.Logger
	r      *gin.Engine
	biz    *ab.MockIBiz
	impl   *impl
}

func (s *SuiteTester) SetupTest() {
	s.logger, _ = zap.NewDevelopment()
	s.biz = new(ab.MockIBiz)

	gin.SetMode(gin.TestMode)
	s.r = gin.New()
	s.r.Use(contextx.AddContextxWitLoggerMiddleware(s.logger))
	s.r.Use(er.AddErrorHandlingMiddleware())

	s.impl = &impl{biz: s.biz}
}

func (s *SuiteTester) TearDownTest() {
	s.biz.AssertExpectations(s.T())
}

func TestAll(t *testing.T) {
	suite.Run(t, new(SuiteTester))
}

func (s *SuiteTester) Test_impl_Login() {
	Handle(s.r.Group("/api/v1/auth"), s.biz)

	type args struct {
		id       string
		password string
		mock     func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
		wantResp *response.Response
	}{
		{
			name:     "missing id then 400",
			args:     args{id: "", password: "password"},
			wantCode: 400,
			wantResp: nil,
		},
		{
			name:     "missing password then 400",
			args:     args{id: "id", password: ""},
			wantCode: 400,
			wantResp: nil,
		},
		{
			name: "login then 500",
			args: args{id: "id", password: "password", mock: func() {
				s.biz.On("Login", mock.Anything, "id", "password").Return(nil, errors.New("error")).Once()
			}},
			wantCode: 500,
			wantResp: nil,
		},
		{
			name: "login then 200",
			args: args{id: "id", password: "password", mock: func() {
				s.biz.On("Login", mock.Anything, "id", "password").Return(testdata.Profile1, nil).Once()
			}},
			wantCode: 200,
			wantResp: nil,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri, _ := url.Parse("/api/v1/auth/login")
			values := url.Values{
				"id":       []string{tt.args.id},
				"password": []string{tt.args.password},
			}
			req := httptest.NewRequest(http.MethodPost, uri.String(), strings.NewReader(values.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			if got.StatusCode != tt.wantCode {
				t.Errorf("Login() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)
			}

			// todo: 2022/11/25|sean|compare response body is equal

			s.TearDownTest()
		})
	}
}
