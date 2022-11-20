package errorx

import (
	"net/http"

	"github.com/blackhorseya/irent/pkg/errorx"
)

var (
	// ErrMissingID means ID must be NOT empty
	ErrMissingID = errorx.NewAPPError(http.StatusBadRequest, 40010, "ID must be NOT empty")

	// ErrMissingPassword means Password must be NOT empty
	ErrMissingPassword = errorx.NewAPPError(http.StatusBadRequest, 40011, "Password must be NOT empty")

	// ErrMissingToken means Token must be NOT empty
	ErrMissingToken = errorx.NewAPPError(http.StatusBadRequest, 40012, "Token must be NOT empty")
)
