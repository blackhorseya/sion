package errorx

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddErrorHandlingMiddleware global handle *gin.Context error middleware
func AddErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if c.Errors.Last() == nil {
				return
			}

			err := c.Errors.Last()
			c.Errors = c.Errors[:0]

			switch err.Err.(type) {
			case *APPError:
				appError := err.Err.(*APPError)
				c.AbortWithStatusJSON(appError.Status, appError)
				break
			default:
				appError := NewAPPError(http.StatusInternalServerError, 50000, err.Err.Error())
				c.AbortWithStatusJSON(http.StatusInternalServerError, appError)
				break
			}
		}()

		c.Next()
	}
}
