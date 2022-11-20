package errorx

import (
	"net/http"

	"github.com/blackhorseya/irent/pkg/errorx"
)

var (
	// ErrLogin means Failed to login then system
	ErrLogin = errorx.NewAPPError(http.StatusInternalServerError, 50010, "Failed to login then system")

	// ErrGetProfileByToken means Failed to get profile by token
	ErrGetProfileByToken = errorx.NewAPPError(http.StatusInternalServerError, 50011, "Failed to get profile by token")
)
