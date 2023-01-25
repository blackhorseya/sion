package orders

import (
	"net/http"

	"github.com/blackhorseya/irent/internal/pkg/errorx"
	"github.com/blackhorseya/irent/pkg/contextx"
	rm "github.com/blackhorseya/irent/pkg/entity/domain/rental/model"
	"github.com/blackhorseya/irent/pkg/httpheaders"
	"github.com/blackhorseya/irent/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type bookCarRequest struct {
	ID         string `json:"id"`
	ProjectID  string `json:"project_id"`
	Circularly bool   `json:"circularly"`
}

// BookCar
// @Summary Book a car
// @Description Book a car
// @Tags Orders
// @Accept application/json
// @Produce application/json
// @Param car body bookCarRequest true "information of car"
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{data=model.Lease}
// @Failure 400 {object} er.Error
// @Failure 500 {object} er.Error
// @Router /v1/orders [post]
func (i *impl) BookCar(c *gin.Context) {
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

	var data bookCarRequest
	err := c.ShouldBindJSON(&data)
	if err != nil {
		ctx.Error(errorx.ErrInvalidRental.Error(), zap.Error(err))
		_ = c.Error(errorx.ErrInvalidRental)
		return
	}

	from, err := i.auth.GetByAccessToken(ctx, token)
	if err != nil {
		_ = c.Error(err)
		return
	}

	// todo: 2023/1/25|sean|refactor: you should be using new function to create car instance
	target := &rm.Car{Id: data.ID, ProjectId: data.ProjectID}
	ret, err := i.biz.BookRental(ctx, from, target)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(ret))
}
