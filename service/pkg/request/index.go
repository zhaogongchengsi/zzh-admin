package request

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type GetGormById struct {
	ID float64 `json:"id" form:"id"` // 主键ID
}

type LimitOffset struct {
	Limit int `json:"limit"`
	Offset int `json:"offset"`
}

func NewLimit(Limit int, offset int) *LimitOffset {
	return &LimitOffset{
		Limit:Limit,
		Offset: offset,
	}
}

func GetLimitOffset(c *gin.Context) *LimitOffset {
	limits := c.DefaultQuery("limit", "0")
	offsets := c.DefaultQuery("offset", "10")
	limit, _ := strconv.Atoi(limits)
	offset, _ := strconv.Atoi(offsets)
	return NewLimit(limit, offset)
}