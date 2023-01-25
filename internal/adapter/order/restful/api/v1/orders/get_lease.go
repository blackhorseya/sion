package orders

import (
	"net/http"

	"github.com/blackhorseya/irent/pkg/response"
	"github.com/gin-gonic/gin"
)

// ListLease
// @Summary List all lease
// @Description List all lease
// @Tags Orders
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{data=[]model.Lease}
// @Failure 400 {object} er.Error
// @Failure 500 {object} er.Error
// @Router /v1/orders [get]
func (i *impl) ListLease(c *gin.Context) {
	// todo: 2023/1/25|sean|impl me
	c.JSON(http.StatusOK, response.OK)
}
