package cars

import (
	"net/http"

	"github.com/blackhorseya/irent/pkg/response"
	"github.com/gin-gonic/gin"
)

// ListCars godoc
// @Summary List cars
// @Description List cars
// @Tags Cars
// @Accept json
// @Produce json
// @Param page query int false "page" default(1) minimum(1)
// @Param size query int false "size" default(10) minimum(1) maximum(100)
// @Param latitude query number false "latitude" default(0)
// @Param longitude query number false "longitude" default(0)
// @Success 200 {object} response.Response{data=[]model.Car}
// @Failure 400 {object} er.Error
// @Failure 500 {object} er.Error
// @Header 200 {int} x-total-count "total count"
// @Router /v1/cars [get]
func (i *impl) ListCars(c *gin.Context) {
	// todo: 2023/5/21|sean|impl me
	c.JSON(http.StatusOK, response.OK)
}
