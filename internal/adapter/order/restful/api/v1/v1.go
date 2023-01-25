package v1

import (
	"github.com/blackhorseya/irent/internal/adapter/order/restful/api/v1/arrears"
	"github.com/blackhorseya/irent/internal/adapter/order/restful/api/v1/orders"
	ab "github.com/blackhorseya/irent/pkg/entity/domain/account/biz"
	ob "github.com/blackhorseya/irent/pkg/entity/domain/order/biz"
	"github.com/blackhorseya/irent/pkg/httpheaders"
	"github.com/gin-gonic/gin"
)

func Handle(g *gin.RouterGroup, biz ob.IBiz, auth ab.IBiz) {
	arrears.Handle(g.Group("/arrears"), biz, auth)
	orders.Handle(g.Group("/orders", httpheaders.AddRequiredAuthMiddleware()), biz, auth)
}
