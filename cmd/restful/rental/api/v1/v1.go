package v1

import (
	"github.com/blackhorseya/irent/cmd/restful/rental/api/v1/cars"
	rb "github.com/blackhorseya/irent/pkg/entity/domain/rental/biz"
	"github.com/gin-gonic/gin"
)

func Handle(g *gin.RouterGroup, biz rb.IBiz) {
	cars.Handle(g.Group("cars"), biz)
}