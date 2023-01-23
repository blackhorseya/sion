package repo

import (
	"github.com/blackhorseya/irent/pkg/contextx"
	am "github.com/blackhorseya/irent/pkg/entity/domain/account/model"
	om "github.com/blackhorseya/irent/pkg/entity/domain/order/model"
)

// IRepo declare account repository functions
//
//go:generate mockery --all --inpackage
type IRepo interface {
	// FetchArrears serve caller to fetch arrears via api
	FetchArrears(ctx contextx.Contextx, user *am.Profile) (records []*om.ArrearsRecord, err error)
}
