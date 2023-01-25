package bookings

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CancelBookingByID
// @Summary Cancel a booking by id
// @Description Cancel a booking by id
// @Tags Bookings
// @Accept application/json
// @Produce application/json
// @Param id path string true "ID of booking"
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{data=string}
// @Failure 400 {object} er.Error
// @Failure 500 {object} er.Error
// @Router /v1/bookings/{id} [delete]
func (i *impl) CancelBookingByID(c *gin.Context) {
	// todo: 2023/1/25|sean|impl me
	c.JSON(http.StatusOK, nil)
}
