package cars

import (
	"net/http"
	"strconv"

	"github.com/blackhorseya/irent/internal/pkg/errorx"
	"github.com/blackhorseya/irent/pkg/contextx"
	rb "github.com/blackhorseya/irent/pkg/entity/domain/rental/biz"
	"github.com/blackhorseya/irent/pkg/entity/domain/rental/model"
	"github.com/blackhorseya/irent/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	querySize = "size"

	queryLatitude = "latitude"

	queryLongitude = "longitude"
)

type listCarsResp struct {
	List  []*model.Car `json:"list"`
	Total int          `json:"total"`
}

// ListCars
// @Summary List cars
// @Description List cars
// @Tags Rental
// @Accept application/json
// @Produce application/json
// @Param page query integer false "page" default(1)
// @Param size query integer false "size" default(5)
// @Param latitude query number false "latitude" default(0)
// @Param longitude query number false "longitude" default(0)
// @Success 200 {object} response.Response{data=listCarsResp}
// @Failure 400,500 {object} er.Error
// @Router /v1/cars [get]
func (i *impl) ListCars(c *gin.Context) {
	ctx, ok := c.MustGet(string(contextx.KeyCtx)).(contextx.Contextx)
	if !ok {
		_ = c.Error(errorx.ErrContextx)
		return
	}

	size, err := strconv.Atoi(c.DefaultQuery(querySize, "5"))
	if err != nil {
		ctx.Error("Failed to parse [size] query to int", zap.Error(err), zap.String(querySize, c.Query(querySize)))
		_ = c.Error(errorx.ErrInvalidSize)
		return
	}

	latitude, err := strconv.ParseFloat(c.DefaultQuery(queryLatitude, "0"), 64)
	if err != nil {
		ctx.Error("Failed to parse [latitude] query to float64", zap.Error(err), zap.String(queryLatitude, c.Query(queryLongitude)))
		_ = c.Error(errorx.ErrInvalidLatitude)
		return
	}

	longitude, err := strconv.ParseFloat(c.DefaultQuery(queryLongitude, "0"), 64)
	if err != nil {
		ctx.Error("Failed to parse [longitude] query to float64", zap.Error(err), zap.String(queryLongitude, c.Query(queryLongitude)))
		_ = c.Error(errorx.ErrInvalidLongitude)
		return
	}

	cond := rb.QueryCarCondition{
		TopNum:    size,
		Latitude:  latitude,
		Longitude: longitude,
	}
	ret, total, err := i.biz.ListCars(ctx, cond)
	if err != nil {
		ctx.Error("Failed to list cars", zap.Error(err), zap.Any("condition", cond))
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(listCarsResp{
		List:  ret,
		Total: total,
	}))
}
