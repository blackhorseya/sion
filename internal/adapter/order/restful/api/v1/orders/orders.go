package orders

import (
	ab "github.com/blackhorseya/irent/pkg/entity/domain/account/biz"
	ob "github.com/blackhorseya/irent/pkg/entity/domain/order/biz"
	"github.com/gin-gonic/gin"
)

func Handle(g *gin.RouterGroup, biz ob.IBiz, auth ab.IBiz) {
	i := &impl{biz: biz, auth: auth}

	g.GET("", i.ListLease)
	g.POST("", i.BookCar)
	g.DELETE(":id", i.CancelLeaseByID)
}

type impl struct {
	biz  ob.IBiz
	auth ab.IBiz
}
