package arrears

import (
	"net/http"

	"github.com/blackhorseya/irent/pkg/response"
	"github.com/gin-gonic/gin"
)

// GetArrears godoc
// @Summary Get arrears list
// @Description Get arrears list
// @Tags Billing
// @Accept json
// @Produce json
// @Param id path string true "user id"
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{data=model.Arrears}
// @Failure 400,401,403 {object} er.Error
// @Failure 500 {object} er.Error
// @Router /v1/accounts/{id}/arrears [get]
func (i *impl) GetArrears(c *gin.Context) {
	// todo: 2023/5/21|sean|impl me
	c.JSON(http.StatusOK, response.OK)
}
