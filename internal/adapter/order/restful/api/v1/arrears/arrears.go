package arrears

import (
	ab "github.com/blackhorseya/irent/pkg/entity/domain/account/biz"
	ob "github.com/blackhorseya/irent/pkg/entity/domain/order/biz"
	"github.com/blackhorseya/irent/pkg/httpheaders"
	"github.com/gin-gonic/gin"
)

func Handle(g *gin.RouterGroup, biz ob.IBiz, auth ab.IBiz) {
	i := &impl{biz: biz, auth: auth}

	g.GET("", httpheaders.AddRequiredAuthMiddleware(), i.GetArrearsByUser)
}

type impl struct {
	biz  ob.IBiz
	auth ab.IBiz
}
