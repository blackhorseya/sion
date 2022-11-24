package errorx

import (
	"net/http"

	"github.com/blackhorseya/irent/pkg/er"
)

var (
	// ErrAuthHeaderFormat means must provide Authorization header with format `Bearer {token}`
	ErrAuthHeaderFormat = er.New(http.StatusUnauthorized, 40100, "Must provide Authorization header with format `Bearer {token}`")
)
