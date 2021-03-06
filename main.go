package main

import (
	"get_price/corns"
	"get_price/route"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	go corns.RegisterCrons()
	r := gin.Default()
	r.Use(Cors())
	route.RegisterRoute(r)
	r.Run(":9090")
}
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

//https://a1.m1907.cn/api/v/?z=f132c044a68b5a642e89f9157518addb&jx=https://m.iqiyi.com/v_26ex87ts5gc.html&s1ig=11401&g=
