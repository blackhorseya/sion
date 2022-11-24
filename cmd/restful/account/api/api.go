package api

import (
	"net/http"

	_ "github.com/blackhorseya/irent/api/account/docs" // import swagger spec
	v1 "github.com/blackhorseya/irent/cmd/restful/account/api/v1"
	"github.com/blackhorseya/irent/internal/pkg/errorx"
	"github.com/blackhorseya/irent/pkg/contextx"
	ab "github.com/blackhorseya/irent/pkg/entity/domain/account/biz"
	_ "github.com/blackhorseya/irent/pkg/errors" // import struct
	"github.com/blackhorseya/irent/pkg/response"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Handle(g *gin.RouterGroup, biz ab.IBiz) {
	ret := impl{biz: biz}

	if gin.Mode() != gin.ReleaseMode {
		g.GET("account/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	g.GET("readiness", ret.Readiness)
	g.GET("liveness", ret.Liveness)

	v1.Handle(g.Group("v1"), biz)
}

type impl struct {
	biz ab.IBiz
}

// Readiness
// @Summary Readiness
// @Description Show application was ready to start accepting traffic
// @Tags Health
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response
// @Failure 500 {object} errors.Error
// @Router /readiness [get]
func (i *impl) Readiness(c *gin.Context) {
	ctx, ok := c.MustGet(string(contextx.KeyCtx)).(contextx.Contextx)
	if !ok {
		_ = c.Error(errorx.ErrContextx)
		return
	}

	err := i.biz.Readiness(ctx)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK)
}

// Liveness
// @Summary Liveness
// @Description to know when to restart an application
// @Tags Health
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response
// @Failure 500 {object} errors.Error
// @Router /liveness [get]
func (i *impl) Liveness(c *gin.Context) {
	ctx, ok := c.MustGet(string(contextx.KeyCtx)).(contextx.Contextx)
	if !ok {
		_ = c.Error(errorx.ErrContextx)
		return
	}

	err := i.biz.Liveness(ctx)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK)
}
