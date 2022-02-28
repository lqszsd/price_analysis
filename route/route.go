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
	}
}
