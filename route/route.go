package route

import (
	"get_price/controller"
	"get_price/controller/Index"
	"get_price/controller/fund"
	"github.com/gin-gonic/gin"
)

func RegisterRoute(router *gin.Engine) {
	v1 := router.Group("/api")
	{
		found := v1.Group("/fund")
		{
			found.Any("/fund_list", fund.Index)
		}
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
	v2 := router.Group("/wechat")
	{
		v2.Any("/login", controller.Login)
	}
	router.GET("/", Index.Index)
	router.GET("/fund_list", Index.Index)
	router.GET("/all_recommend", Index.Index)
	router.GET("/my_strategy", Index.Index)
	router.GET("/my_choose", Index.Index)
	router.GET("/recommend", Index.Index)
	router.GET("/fund_new", Index.Index)

}
