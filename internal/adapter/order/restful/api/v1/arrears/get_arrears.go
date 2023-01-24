package arrears

import (
	"net/http"

	"github.com/blackhorseya/irent/pkg/response"
	"github.com/gin-gonic/gin"
)

// GetArrearsByUser
// @Summary Get arrears by user
// @Description Get arrears by user
// @Tags Billing
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{data=model.Arrears}
// @Failure 500 {object} er.Error
// @Router /v1/arrears [get]
func (i *impl) GetArrearsByUser(c *gin.Context) {
	// todo: 2023/1/24|sean|impl me

	c.JSON(http.StatusOK, response.OK)
}
