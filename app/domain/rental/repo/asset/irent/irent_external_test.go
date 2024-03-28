//go:build external

package irent

import (
	"testing"

	"github.com/blackhorseya/sion/entity/domain/rental/repo"
	"github.com/blackhorseya/sion/pkg/configx"
	"github.com/blackhorseya/sion/pkg/contextx"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type suiteExternal struct {
	suite.Suite

	repo repo.IAssetRepo
}

func (s *suiteExternal) SetupTest() {
	err := configx.LoadWithPathAndName("", "sion")
	s.Require().NoError(err)

	s.repo, err = NewAssetRepo()
	s.NoError(err)
}

func TestExternal(t *testing.T) {
	suite.Run(t, new(suiteExternal))
}

func (s *suiteExternal) TestFetchAvailableCars() {
	ctx := contextx.Background()
	cars, err := s.repo.FetchAvailableCars(ctx)
	s.NoError(err)
	ctx.Debug("cars", zap.Any("cars", cars))
}
