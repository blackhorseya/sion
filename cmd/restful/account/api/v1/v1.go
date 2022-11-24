package v1

import (
	"github.com/blackhorseya/irent/cmd/restful/account/api/v1/auth"
	ab "github.com/blackhorseya/irent/pkg/entity/domain/account/biz"
	"github.com/gin-gonic/gin"
)

func Handle(g *gin.RouterGroup, biz ab.IBiz) {
	auth.Handle(g.Group("auth"), biz)
}
