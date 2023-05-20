package restful

import (
	accountB "github.com/blackhorseya/irent/pkg/entity/domain/account/biz"
	orderB "github.com/blackhorseya/irent/pkg/entity/domain/order/biz"
	rentalB "github.com/blackhorseya/irent/pkg/entity/domain/rental/biz"
	"github.com/gin-gonic/gin"
)

type impl struct {
	router  *gin.Engine
	account accountB.IBiz
	rental  rentalB.IBiz
	order   orderB.IBiz
}
