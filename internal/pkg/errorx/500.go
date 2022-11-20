package errorx

import (
	"net/http"

	"github.com/blackhorseya/irent/pkg/errorx"
)

var (
	// ErrLogin means Failed to login then system
	ErrLogin = errorx.NewAPPError(http.StatusInternalServerError, 50010, "Failed to login then system")
)
