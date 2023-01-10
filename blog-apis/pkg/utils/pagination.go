package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"coinsky_go_project/blog-apis/pkg/setting"
)

// 分页offset
func GetPagingOffset(c *gin.Context) int {
	offset := 0
	page, _ := strconv.Atoi(c.Query("page"))
	if page > 0 {
		offset = (page - 1) * setting.PageSize
	}
	return offset
}
