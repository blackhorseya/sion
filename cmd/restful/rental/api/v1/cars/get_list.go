package cars

import (
	"github.com/gin-gonic/gin"
)

// ListCars
// @Summary List cars
// @Description List cars
// @Tags Rental
// @Accept application/json
// @Produce application/json
// @Param page query integer false "page" default(1)
// @Param size query integer false "size" default(5)
// @Param latitude query number false "latitude" default(0)
// @Param longitude query number false "longitude" default(0)
// @Success 200 {object} response.Response{data=[]model.Car}
// @Failure 400,500 {object} er.Error
// @Router /v1/cars [get]
func (i *impl) ListCars(c *gin.Context) {
	// todo: 2022/12/4|sean|impl me
	panic("impl me")
}
