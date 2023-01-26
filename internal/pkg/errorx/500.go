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

	// ErrListCars means Failed to list cars via api
	ErrListCars = er.New(http.StatusInternalServerError, 50012, "Failed to list cars via api")

	// ErrGetArrears means Failed to get arrears
	ErrGetArrears = er.New(http.StatusInternalServerError, 50013, "Failed to get arrears")

	// ErrBookRental means Failed to book rental
	ErrBookRental = er.New(http.StatusInternalServerError, 50014, "Failed to book rental")

	// ErrCancelBooking means Failed to cancel booking
	ErrCancelBooking = er.New(http.StatusInternalServerError, 50015, "Failed to cancel booking")

	// ErrGetLease means Failed to get lease
	ErrGetLease = er.New(http.StatusInternalServerError, 50016, "Failed to get lease")
)
