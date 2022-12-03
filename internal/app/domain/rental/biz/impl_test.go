package biz

import (
	"reflect"
	"testing"

	"github.com/blackhorseya/irent/internal/app/domain/rental/biz/repo"
	"github.com/blackhorseya/irent/pkg/contextx"
	rb "github.com/blackhorseya/irent/pkg/entity/domain/rental/biz"
	rm "github.com/blackhorseya/irent/pkg/entity/domain/rental/model"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type SuiteTester struct {
	suite.Suite
	logger *zap.Logger
	repo   *repo.MockIRepo
	biz    rb.IBiz
}

func (s *SuiteTester) SetupTest() {
	s.logger, _ = zap.NewDevelopment()
	s.repo = new(repo.MockIRepo)
	s.biz = CreateBiz(s.repo)
}

func (s *SuiteTester) assertExpectation(t *testing.T) {
	s.repo.AssertExpectations(t)
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
				s.repo.On("ListCars", mock.Anything).Return(nil, errors.New("error")).Once()
			}},
			wantInfo:  nil,
			wantTotal: 0,
			wantErr:   true,
		},
		{
			name: "not found any cars then nil",
			args: args{condition: rb.QueryCarCondition{TopNum: 10}, mock: func() {
				s.repo.On("ListCars", mock.Anything).Return(nil, nil).Once()
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

			s.assertExpectation(t)
		})
	}
}
