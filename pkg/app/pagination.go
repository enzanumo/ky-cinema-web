package app

import (
	"github.com/enzanumo/ky-theater-web/pkg/convert"
	"github.com/gin-gonic/gin"
)

var (
	DefaultPageSize = 10
	MaxPageSize     = 100
)

func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}

	return page
}

func GetPageSize(c *gin.Context) int {
	pageSize := convert.StrTo(c.Query("page_size")).MustInt()
	if pageSize <= 0 {
		return DefaultPageSize
	}
	if pageSize > MaxPageSize {
		return MaxPageSize
	}

	return pageSize
}

func GetPageOffset(c *gin.Context) (offset, limit int) {
	page := convert.StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		page = 1
	}

	limit = convert.StrTo(c.Query("page_size")).MustInt()
	if limit <= 0 {
		limit = DefaultPageSize
	} else if limit > MaxPageSize {
		limit = MaxPageSize
	}
	offset = (page - 1) * limit
	return
}
