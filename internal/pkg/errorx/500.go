package errorx

import (
	"net/http"

	"github.com/blackhorseya/irent/pkg/errorx"
)

var (
	// ErrContextx means Missing contextx
	ErrContextx = errorx.New(http.StatusInternalServerError, 50001, "Missing contextx")
)

var (
	// ErrLogin means Failed to login then system
	ErrLogin = errorx.New(http.StatusInternalServerError, 50010, "Failed to login then system")

	// ErrGetProfileByToken means Failed to get profile by token
	ErrGetProfileByToken = errorx.New(http.StatusInternalServerError, 50011, "Failed to get profile by token")
)
