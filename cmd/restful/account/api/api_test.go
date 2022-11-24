package api

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/blackhorseya/irent/pkg/contextx"
	ab "github.com/blackhorseya/irent/pkg/entity/domain/account/biz"
	"github.com/blackhorseya/irent/pkg/er"
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

func (s *SuiteTester) Test_impl_Readiness() {
	Handle(s.r.Group("api"), s.biz)

	type args struct {
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name: "readiness then 500",
			args: args{mock: func() {
				s.biz.On("Readiness", mock.Anything).Return(errors.New("error")).Once()
			}},
			wantCode: 500,
		},
		{
			name: "readiness then 200",
			args: args{mock: func() {
				s.biz.On("Readiness", mock.Anything).Return(nil).Once()
			}},
			wantCode: 200,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri, _ := url.Parse("/api/readiness")
			req := httptest.NewRequest(http.MethodGet, uri.String(), nil)
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			if got.StatusCode != tt.wantCode {
				t.Errorf("Readiness() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)
			}

			s.TearDownTest()
		})
	}
}

func (s *SuiteTester) Test_impl_Liveness() {
	Handle(s.r.Group("api"), s.biz)

	type args struct {
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name: "liveness then 500",
			args: args{mock: func() {
				s.biz.On("Liveness", mock.Anything).Return(errors.New("error")).Once()
			}},
			wantCode: 500,
		},
		{
			name: "liveness then 200",
			args: args{mock: func() {
				s.biz.On("Liveness", mock.Anything).Return(nil).Once()
			}},
			wantCode: 200,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri, _ := url.Parse("/api/liveness")
			req := httptest.NewRequest(http.MethodGet, uri.String(), nil)
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			if got.StatusCode != tt.wantCode {
				t.Errorf("Liveness() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)
			}

			s.TearDownTest()
		})
	}
}
