package arrears

import (
	"net/http"

	"github.com/blackhorseya/irent/internal/pkg/errorx"
	"github.com/blackhorseya/irent/pkg/contextx"
	"github.com/blackhorseya/irent/pkg/httpheaders"
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

	from, err := i.auth.GetByAccessToken(ctx, token)
	if err != nil {
		_ = c.Error(err)
		return
	}

	ret, err := i.biz.GetArrears(ctx, from)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(ret))
}
