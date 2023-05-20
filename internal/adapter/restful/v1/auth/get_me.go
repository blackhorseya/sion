package auth

import (
	"net/http"

	"github.com/blackhorseya/irent/pkg/response"
	"github.com/gin-gonic/gin"
)

// Me godoc
// @Summary Get me profile
// @Description Get me profile
// @Tags Auth
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{data=model.Profile}
// @Failure 401,403 {object} er.Error
// @Failure 500 {object} er.Error
// @Router /v1/auth/me [get]
func (i *impl) Me(c *gin.Context) {
	// todo: 2023/5/21|sean|impl me
	c.JSON(http.StatusOK, response.OK)
}
