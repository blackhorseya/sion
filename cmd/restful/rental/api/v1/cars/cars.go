package cars

import (
	rb "github.com/blackhorseya/irent/pkg/entity/domain/rental/biz"
	"github.com/gin-gonic/gin"
)

type impl struct {
	biz rb.IBiz
}

func Handle(g *gin.RouterGroup, biz rb.IBiz) {
	i := &impl{biz: biz}

	g.GET("", i.ListCars)
}
