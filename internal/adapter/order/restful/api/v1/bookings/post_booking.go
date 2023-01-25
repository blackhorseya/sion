package bookings

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type bookCarRequest struct {
	ID         string `json:"id"`
	ProjectID  string `json:"project_id"`
	Circularly bool   `json:"circularly"`
}

// BookCar
// @Summary Book a car
// @Description Book a car
// @Tags Bookings
// @Accept application/json
// @Produce application/json
// @Param car body bookCarRequest true "information of car"
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{data=model.Booking}
// @Failure 400 {object} er.Error
// @Failure 500 {object} er.Error
// @Router /v1/bookings [post]
func (i *impl) BookCar(c *gin.Context) {
	// todo: 2023/1/25|sean|impl me
	c.JSON(http.StatusOK, nil)
}
