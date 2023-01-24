package restful

import (
	"time"

	"github.com/blackhorseya/irent/internal/adapter/order/restful/api"
	"github.com/blackhorseya/irent/pkg/adapters"
	"github.com/blackhorseya/irent/pkg/contextx"
	"github.com/blackhorseya/irent/pkg/cors"
	ob "github.com/blackhorseya/irent/pkg/entity/domain/order/biz"
	"github.com/blackhorseya/irent/pkg/er"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go.uber.org/zap"
)

type restful struct {
	router *gin.Engine
	biz    ob.IBiz
}

func NewRestful(logger *zap.Logger, router *gin.Engine, biz ob.IBiz) adapters.Restful {
	router.Use(cors.AddAllowAll())
	router.Use(ginzap.RecoveryWithZap(logger, true))
	router.Use(ginzap.GinzapWithConfig(logger, &ginzap.Config{
		TimeFormat: time.RFC3339,
		UTC:        true,
		SkipPaths:  []string{"/api/readiness", "/api/liveness", "/metrics"},
	}))
	router.Use(contextx.AddContextxWitLoggerMiddleware(logger))
	router.Use(er.AddErrorHandlingMiddleware())

	return &restful{
		router: router,
		biz:    biz,
	}
}

func (i *restful) InitRouting() error {
	api.Handle(i.router.Group("/api"), i.biz)

	return nil
}

var OrderSet = wire.NewSet(NewRestful)
