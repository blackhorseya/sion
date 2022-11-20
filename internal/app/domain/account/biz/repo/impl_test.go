package repo

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type SuiteTester struct {
	suite.Suite
	logger *zap.Logger
	repo   IRepo
}

func (s *SuiteTester) SetupTest() {
	s.logger, _ = zap.NewDevelopment()
	s.repo = CreateRepo()
}

func (s *SuiteTester) TearDownTest() {
}

func TestAll(t *testing.T) {
	suite.Run(t, new(SuiteTester))
}
