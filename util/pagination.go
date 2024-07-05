package util

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

// GetPage get page parameters
func GetPage(c *gin.Context) int {
	result := 0
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
	if err != nil {
		return result
	}
	if page > 0 {
		result = (page - 1) * pageSize
	}

	return result
}
