package repo

import (
	"github.com/blackhorseya/irent/pkg/contextx"
	am "github.com/blackhorseya/irent/pkg/entity/domain/account/model"
)

// IRepo declare account repository functions
//
//go:generate mockery --all --inpackage
type IRepo interface {
	// Login serve caller to log in the system
	Login(ctx contextx.Contextx, id, password string) (info *am.Profile, err error)

	// GetMemberStatus serve caller to given access token to get profile
	GetMemberStatus(ctx contextx.Contextx, token string) (info *am.Profile, err error)
}
