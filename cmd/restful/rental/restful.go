package main

import (
	"time"

	"github.com/blackhorseya/irent/cmd"
	"github.com/blackhorseya/irent/cmd/restful/rental/api"
	"github.com/blackhorseya/irent/pkg/contextx"
	"github.com/blackhorseya/irent/pkg/cors"
	rb "github.com/blackhorseya/irent/pkg/entity/domain/rental/biz"
	"github.com/blackhorseya/irent/pkg/er"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type restful struct {
	router *gin.Engine
	biz    rb.IBiz
}

func NewRestful(logger *zap.Logger, router *gin.Engine, biz rb.IBiz) cmd.Restful {
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
	api.Handle(i.router.Group("api"), i.biz)

	return nil
}
