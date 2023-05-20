package orders

import (
	"net/http"

	"github.com/blackhorseya/irent/pkg/response"
	"github.com/gin-gonic/gin"
)

type bookCarRequest struct {
	ID         string `json:"id"`
	ProjectID  string `json:"project_id"`
	Circularly bool   `json:"circularly"`
}

// BookCar will book the car
// @Summary Book the car
// @Description book the car
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path string true "user id"
// @Param car body bookCarRequest true "book car request"
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{data=model.Lease}
// @Failure 400,401,403 {object} er.Error
// @Failure 500 {object} er.Error
// @Router /accounts/{id}/orders [post]
func (i *impl) BookCar(c *gin.Context) {
	// todo: 2023/5/21|sean|impl me
	c.JSON(http.StatusOK, response.OK)
}
