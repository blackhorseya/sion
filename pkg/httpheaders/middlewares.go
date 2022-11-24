package httpheaders

import (
	"strings"

	"github.com/blackhorseya/irent/internal/pkg/errorx"
	"github.com/gin-gonic/gin"
)

// Key declare key string
type Key string

var (
	// KeyToken token key string
	KeyToken = Key("token")
)

type authHeader struct {
	Value string `header:"Authorization"`
}

// AddRequiredAuthMiddleware serve caller to extract authorization header value
func AddRequiredAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := authHeader{}
		err := c.ShouldBindHeader(&h)
		if err != nil {
			_ = c.Error(errorx.ErrMissingToken)
			c.Abort()
			return
		}

		headers := strings.Split(h.Value, "Bearer ")
		if len(headers) < 2 {
			_ = c.Error(errorx.ErrAuthHeaderFormat)
			c.Abort()
			return
		}

		c.Set(string(KeyToken), headers[1])

		c.Next()
	}
}
