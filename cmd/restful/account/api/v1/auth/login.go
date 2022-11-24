package auth

import (
	"net/http"

	"github.com/blackhorseya/irent/internal/pkg/errorx"
	"github.com/blackhorseya/irent/pkg/contextx"
	_ "github.com/blackhorseya/irent/pkg/entity/domain/account/model" // import struct
	_ "github.com/blackhorseya/irent/pkg/errors"                      // import struct
	"github.com/blackhorseya/irent/pkg/response"
	"github.com/gin-gonic/gin"
)

// Login
// @Summary Login
// @Description Login
// @Tags Auth
// @Accept x-www-form-urlencoded
// @Produce application/json
// @Param id formData string true "user id"
// @Param password formData string true "user password"
// @Success 201 {object} response.Response{data=model.Profile}
// @Failure 400 {object} errors.Error
// @Failure 500 {object} errors.Error
// @Router /v1/auth/login [post]
func (i *impl) Login(c *gin.Context) {
	ctx, ok := c.MustGet(string(contextx.KeyCtx)).(contextx.Contextx)
	if !ok {
		_ = c.Error(errorx.ErrContextx)
		return
	}

	id := c.PostForm("id")
	if len(id) == 0 {
		ctx.Error(errorx.ErrMissingID.Error())
		_ = c.Error(errorx.ErrMissingID)
		return
	}

	password := c.PostForm("password")
	if len(password) == 0 {
		ctx.Error(errorx.ErrMissingPassword.Error())
		_ = c.Error(errorx.ErrMissingPassword)
		return
	}

	ret, err := i.biz.Login(ctx, id, password)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(ret))
}
