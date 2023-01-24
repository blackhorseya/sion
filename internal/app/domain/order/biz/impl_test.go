package biz

import (
	"reflect"
	"testing"

	"github.com/blackhorseya/irent/internal/app/domain/order/biz/repo"
	"github.com/blackhorseya/irent/pkg/contextx"
	am "github.com/blackhorseya/irent/pkg/entity/domain/account/model"
	ob "github.com/blackhorseya/irent/pkg/entity/domain/order/biz"
	om "github.com/blackhorseya/irent/pkg/entity/domain/order/model"
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
	biz    ob.IBiz
}

func (s *SuiteTester) SetupTest() {
	s.logger, _ = zap.NewDevelopment()
	s.repo = new(repo.MockIRepo)
	s.biz = CreateBiz(s.repo)
}

func (s *SuiteTester) AssertExpectations(t *testing.T) {
	s.repo.AssertExpectations(t)
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
			name: "fetch then error",
			args: args{from: testdata.Profile1, target: testdata.Profile1, mock: func() {
				s.repo.On("FetchArrears", mock.Anything, testdata.Profile1, testdata.Profile1).Return(nil, errors.New("error")).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "ok",
			args: args{from: testdata.Profile1, target: testdata.Profile1, mock: func() {
				s.repo.On("FetchArrears", mock.Anything, testdata.Profile1, testdata.Profile1).Return(nil, nil).Once()
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

			s.AssertExpectations(t)
		})
	}
}
