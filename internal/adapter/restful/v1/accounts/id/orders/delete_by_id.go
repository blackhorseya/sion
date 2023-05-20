package orders

import (
	"net/http"

	"github.com/blackhorseya/irent/pkg/response"
	"github.com/gin-gonic/gin"
)

// CancelByID will cancel the order by id
// @Summary Cancel the order by id
// @Description cancel the order by id
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path string true "user id"
// @Param order_id path string true "order id"
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{data=string}
// @Failure 400,401,403 {object} er.Error
// @Failure 500 {object} er.Error
// @Router /accounts/{id}/orders/{order_id} [delete]
func (i *impl) CancelByID(c *gin.Context) {
	// todo: 2023/5/21|sean|impl me
	c.JSON(http.StatusOK, response.OK)
}
