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
		v1.POST("/add_strategy", controller.AddStrategy)
		v1.GET("/user_strategy", controller.UserStrategy)
		v1.POST("/change_strategy_status", controller.ChangeStrategyStatus)
		v1.GET("/news_list", controller.NewList)
		v1.GET("/rm_select", controller.RmSelect)
		v1.GET("/all_recommend", controller.AllSymbol)
	}
}
