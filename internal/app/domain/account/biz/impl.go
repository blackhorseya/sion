package biz

import (
	"crypto/sha256"
	"fmt"

	"github.com/blackhorseya/irent/internal/app/domain/account/biz/repo"
	"github.com/blackhorseya/irent/internal/pkg/errorx"
	"github.com/blackhorseya/irent/pkg/contextx"
	ab "github.com/blackhorseya/irent/pkg/entity/domain/account/biz"
	am "github.com/blackhorseya/irent/pkg/entity/domain/account/model"
	"github.com/google/wire"
	"go.uber.org/zap"
)

// ProviderSet is a provider set for account biz
var ProviderSet = wire.NewSet(NewImpl, repo.ProviderSet)

type impl struct {
	repo repo.IRepo
}

func NewImpl(repo repo.IRepo) ab.IBiz {
	return &impl{
		repo: repo,
	}
}

func (i *impl) Readiness(ctx contextx.Contextx) error {
	return nil
}

func (i *impl) Liveness(ctx contextx.Contextx) error {
	return nil
}

func (i *impl) Login(ctx contextx.Contextx, id, password string) (info *am.Profile, err error) {
	if len(id) == 0 {
		ctx.Error(errorx.ErrMissingID.Error())
		return nil, errorx.ErrMissingID
	}

	if len(password) == 0 {
		ctx.Error(errorx.ErrMissingPassword.Error())
		return nil, errorx.ErrMissingPassword
	}

	ret, err := i.repo.Login(ctx, id, encryptPassword(password))
	if err != nil {
		ctx.Error(errorx.ErrLogin.Error(), zap.Error(err))
		return nil, errorx.ErrLogin
	}

	return ret, nil
}

func (i *impl) GetByAccessToken(ctx contextx.Contextx, token string) (info *am.Profile, err error) {
	if len(token) == 0 {
		ctx.Error(errorx.ErrMissingToken.Error())
		return nil, errorx.ErrMissingToken
	}

	ret, err := i.repo.GetMemberStatus(ctx, token)
	if err != nil {
		ctx.Error(errorx.ErrGetProfileByToken.Error(), zap.Error(err), zap.String("token", token))
		return nil, errorx.ErrGetProfileByToken
	}

	return ret, nil
}

func encryptPassword(password string) string {
	return fmt.Sprintf("0x%x", sha256.Sum256([]byte(password)))
}
