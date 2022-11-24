package auth

import (
	ab "github.com/blackhorseya/irent/pkg/entity/domain/account/biz"
	"github.com/gin-gonic/gin"
)

func Handle(g *gin.RouterGroup, biz ab.IBiz) {
	i := &impl{biz: biz}

	g.POST("login", i.Login)
	g.GET("me", i.Me)
}

type impl struct {
	biz ab.IBiz
}
