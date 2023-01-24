package arrears

import (
	ob "github.com/blackhorseya/irent/pkg/entity/domain/order/biz"
	"github.com/blackhorseya/irent/pkg/httpheaders"
	"github.com/gin-gonic/gin"
)

func Handle(g *gin.RouterGroup, biz ob.IBiz) {
	i := &impl{biz: biz}

	g.GET("", httpheaders.AddRequiredAuthMiddleware(), i.GetArrearsByUser)
}

type impl struct {
	biz ob.IBiz
}
