package route

import (
	"get_price/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRoute(router *gin.Engine) {
	v1 := router.Group("/api")
	{
		v1.Any("/", controller.Index)
		v1.Any("/recommend", controller.Recommend)
		v1.Any("/hot/list", controller.HotList)
		v1.Any("/user_select", controller.UserSelect)
		v1.POST("/add_select", controller.AddSelect)
	}
}
