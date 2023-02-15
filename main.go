package main

import (
	"embed"
	"get_price/corns"
	"get_price/route"
	"get_price/service/fund"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

//go:embed dist/css
var css embed.FS

//go:embed dist/css
var js embed.FS

//go:embed dist/css
var fonts embed.FS

//go:embed dist/index.html
var html embed.FS

func main() {

	go corns.RegisterCrons()
	go fund.Register_fund_data()
	r := gin.Default()
	tpl := template.Must(template.ParseFS(html, "dist/index.html"))
	r.SetHTMLTemplate(tpl)
	//r.Any("/js", func(c *gin.Context) {
	//	staticServer := http.FileServer(http.FS(js))
	//	staticServer.ServeHTTP(c.Writer, c.Request)
	//})
	//res,_:=css.ReadDir("css")
	//data,_:=css.ReadFile("app.5a6ad033.css")
	//fmt.Println(string(data),"读取文件信息",res)
	//r.Any("/css", func(c *gin.Context) {
	//	staticServer := http.FileServer(http.FS(css))
	//	staticServer.ServeHTTP(c.Writer, c.Request)
	//})
	//r.StaticFS("/js", gin.DirByFile(http.FS(js), false))
	//r.StaticFS("/css", http.FS(css))
	//r.StaticFS("/js", http.FS(js))
	r.Static("/css", "./dist/css")
	r.Static("/js", "./dist/js")
	r.Static("/fonts", "./dist/fonts")
	r.Any("/fonts", func(c *gin.Context) {
		staticServer := http.FileServer(http.FS(fonts))
		staticServer.ServeHTTP(c.Writer, c.Request)
	})
	//r.StaticFS("/fonts", http.FS(fonts))

	//r.Static("/js", "/views/dist/js")
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
