package arrears

import (
	"net/http"

	"github.com/blackhorseya/irent/internal/pkg/errorx"
	"github.com/blackhorseya/irent/pkg/contextx"
	am "github.com/blackhorseya/irent/pkg/entity/domain/account/model"
	"github.com/blackhorseya/irent/pkg/httpheaders"
	"github.com/blackhorseya/irent/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type getArrearsByIdRequest struct {
	ID string `uri:"id" binding:"required"`
}

// GetArrearsById
// @Summary Get arrears by user
// @Description Get arrears by user
// @Tags Billing
// @Accept application/json
// @Produce application/json
// @Param id path string true "ID of user"
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{data=model.Arrears}
// @Failure 500 {object} er.Error
// @Router /v1/arrears/{id} [get]
func (i *impl) GetArrearsById(c *gin.Context) {
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

	var req getArrearsByIdRequest
	err := c.ShouldBindUri(&req)
	if err != nil {
		ctx.Error(errorx.ErrMissingID.Error(), zap.Error(err))
		_ = c.Error(errorx.ErrMissingID)
		return
	}
	target := &am.Profile{Id: req.ID}

	from, err := i.auth.GetByAccessToken(ctx, token)
	if err != nil {
		_ = c.Error(err)
		return
	}

	ret, err := i.biz.GetArrears(ctx, from, target)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(ret))
}
