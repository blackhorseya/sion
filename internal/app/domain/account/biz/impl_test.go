package biz

import (
	"testing"

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
