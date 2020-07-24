package util

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"gin-blog/pkg/setting"
)

/**
获取offset
第一页 offset 0 10
第二页 offset 10 10
第三页 offset 20 10
*/
func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.AppSetting.PageSize
	}

	return result
}
