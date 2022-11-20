package biz

import (
	"testing"

	"github.com/blackhorseya/irent/pkg/contextx"
	ab "github.com/blackhorseya/irent/pkg/entity/domain/account/biz"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type SuiteTester struct {
	suite.Suite
	logger *zap.Logger
	biz    ab.IBiz
}

func (s *SuiteTester) SetupTest() {
	s.logger, _ = zap.NewDevelopment()
	s.biz = CreateBiz()
}

func (s *SuiteTester) TearDownTest() {
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
