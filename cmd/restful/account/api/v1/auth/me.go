package auth

import (
	"net/http"

	_ "github.com/blackhorseya/irent/pkg/entity/domain/account/model" // import struct
	_ "github.com/blackhorseya/irent/pkg/errors"                      // import struct
	"github.com/blackhorseya/irent/pkg/response"
	"github.com/gin-gonic/gin"
)

// Me
// @Summary Get me profile
// @Description Get me profile
// @Tags Auth
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{data=model.Profile}
// @Failure 401 {object} errors.Error
// @Failure 500 {object} errors.Error
// @Router /v1/auth/me [get]
func (i *impl) Me(c *gin.Context) {
	c.JSON(http.StatusOK, response.OK)
}
