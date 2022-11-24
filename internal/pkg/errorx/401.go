package errorx

import (
	"net/http"

	"github.com/blackhorseya/irent/pkg/errors"
)

var (
	// ErrAuthHeaderFormat means must provide Authorization header with format `Bearer {token}`
	ErrAuthHeaderFormat = errors.New(http.StatusUnauthorized, 40100, "Must provide Authorization header with format `Bearer {token}`")
)
