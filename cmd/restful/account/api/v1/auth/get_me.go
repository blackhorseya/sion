package auth

import (
	"net/http"

	"github.com/blackhorseya/irent/internal/pkg/errorx"
	"github.com/blackhorseya/irent/pkg/contextx"
	_ "github.com/blackhorseya/irent/pkg/entity/domain/account/model" // import struct
	_ "github.com/blackhorseya/irent/pkg/er"                          // import struct
	"github.com/blackhorseya/irent/pkg/httpheaders"
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
// @Failure 401 {object} er.Error
// @Failure 500 {object} er.Error
// @Router /v1/auth/me [get]
func (i *impl) Me(c *gin.Context) {
	ctx, ok := c.MustGet(string(contextx.KeyCtx)).(contextx.Contextx)
	if !ok {
		_ = c.Error(errorx.ErrContextx)
		return
	}

	token, ok := c.MustGet(string(httpheaders.KeyToken)).(string)
	if !ok {
		ctx.Error(errorx.ErrMissingToken.Error())
		_ = c.Error(errorx.ErrMissingToken)
		return
	}

	ret, err := i.biz.GetByAccessToken(ctx, token)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(ret))
}
