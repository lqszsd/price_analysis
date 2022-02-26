package main

import (
	"encoding/json"
	"fmt"
	"get_price/service"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {
	r := gin.Default()
	r.Use(Cors())
	r.POST("/", func(c *gin.Context) {
		resBody := c.Request.Body
		param := &struct {
			Symbol    string `json:"symbol"`
			StartDate string `json:"start_date"`
			EndDate   string `json:"end_date"`
		}{}
		json.NewDecoder(resBody).Decode(param)
		fmt.Println(param)
		res, _ := http.Get("http://127.0.0.1:8080/api/public/bond_zh_hs_cov_min?symbol=" + param.Symbol + "&start_date=" + param.StartDate + "&end_date=" + param.EndDate + "&adjust=hfq")
		fmt.Println("http://127.0.0.1:8080/api/public/bond_zh_hs_cov_min?symbol=" + param.Symbol + "&start_date=" + param.StartDate + "&end_date=" + param.EndDate + "&adjust=hfq")
		by, _ := ioutil.ReadAll(res.Body)
		c.String(200, string(by))
	})
	r.Run(":9090")
	select {}
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		resBody := request.Body

		param := &struct {
			Symbol    string `json:"symbol"`
			StartDate string `json:"start_date"`
			EndDate   string `json:"end_date"`
		}{}
		// 通过json解析器解析参数
		json.NewDecoder(resBody).Decode(param)
		fmt.Println(param)
		writer.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
		writer.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		writer.Header().Set("content-type", "application/json")             //返回数据格式是json
		res, _ := http.Get("http://127.0.0.1:8080/api/public/bond_zh_hs_cov_min?symbol=" + param.Symbol + "&start_date=" + param.StartDate + "&end_date=" + param.EndDate + "&adjust=hfq")
		fmt.Println("http://127.0.0.1:8080/api/public/bond_zh_hs_cov_min?symbol=" + param.Symbol + "&start_date=" + param.StartDate + "&end_date=" + param.EndDate + "&adjust=hfq")
		by, _ := ioutil.ReadAll(res.Body)
		writer.Write(by)
	})

	http.HandleFunc("/recommend", func(writer http.ResponseWriter, request *http.Request) {

		// 通过json解析器解析参数

		writer.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
		writer.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		writer.Header().Set("content-type", "application/json")             //返回数据格式是json
		list := service.Filter_expensive_list()
		data, _ := json.Marshal(list)
		writer.Write(data)
	})

	http.HandleFunc("/api/recommend", func(writer http.ResponseWriter, request *http.Request) {
		code := request.FormValue("code")
		price := request.FormValue("price")
		// 通过json解析器解析参数

		writer.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
		writer.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		writer.Header().Set("content-type", "application/json")             //返回数据格式是json
		money, _ := strconv.ParseFloat(price, 64)
		list := service.GetMoneyRecommend(code, money)
		data, _ := json.Marshal(list)
		writer.Write(data)
	})
	err := http.ListenAndServe("127.0.0.1:9090", nil)
	fmt.Println(err)
	for {
		fmt.Println(err)
	}

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
