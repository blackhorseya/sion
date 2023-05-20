package auth

import (
	"net/http"

	"github.com/blackhorseya/irent/pkg/response"
	"github.com/gin-gonic/gin"
)

// Login godoc
// @Summary Login
// @Description Login
// @Tags Auth
// @Accept json
// @Produce json
// @Param id formData string true "user id"
// @Param password formData string true "user password"
// @Success 200 {object} response.Response{data=model.Profile}
// @Failure 400 {object} er.Error
// @Failure 500 {object} er.Error
// @Router /v1/auth/login [post]
func (i *impl) Login(c *gin.Context) {
	// todo: 2023/5/21|sean|impl me
	c.JSON(http.StatusOK, response.OK)
}
