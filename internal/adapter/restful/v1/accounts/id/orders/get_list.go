package orders

import (
	"net/http"

	"github.com/blackhorseya/irent/pkg/response"
	"github.com/gin-gonic/gin"
)

// ListOrders godoc
// @Summary List orders
// @Description List orders
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path string true "user id"
// @Param page query int false "page number" default(1) minimum(1)
// @Param size query int false "page size" default(10) minimum(1) maximum(100)
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{data=[]model.Lease}
// @Failure 400,401,403 {object} er.Error
// @Failure 500 {object} er.Error
// @Router /v1/accounts/{id}/orders [get]
func (i *impl) ListOrders(c *gin.Context) {
	// todo: 2023/5/21|sean|impl me
	c.JSON(http.StatusOK, response.OK)
}
