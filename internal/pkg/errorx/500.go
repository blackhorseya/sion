package errorx

import (
	"net/http"

	"github.com/blackhorseya/irent/pkg/er"
)

var (
	// ErrContextx means Missing contextx
	ErrContextx = er.New(http.StatusInternalServerError, 50001, "Missing contextx")
)

var (
	// ErrLogin means Failed to log in then system
	ErrLogin = er.New(http.StatusInternalServerError, 50010, "Failed to login then system")

	// ErrGetProfileByToken means Failed to get profile by token
	ErrGetProfileByToken = er.New(http.StatusInternalServerError, 50011, "Failed to get profile by token")
)
