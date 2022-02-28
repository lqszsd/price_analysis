package controller

import (
	"encoding/json"
	"fmt"
	"get_price/service"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

func Index(c *gin.Context) {
	resBody := c.Request.Body
	param := &struct {
		Symbol    string `json:"symbol"`
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
	}{}
	json.NewDecoder(resBody).Decode(param)
	fmt.Println(param)
	param.Symbol = strings.Replace(param.Symbol, "SZ", "", -1)
	param.Symbol = strings.Replace(param.Symbol, "SH", "", -1)
	param.Symbol = strings.Replace(param.Symbol, "sz", "", -1)
	param.Symbol = strings.Replace(param.Symbol, "sh", "", -1)
	fmt.Println("http://127.0.0.1:8080/api/public/stock_zh_a_hist_min_em?symbol=" + param.Symbol + "&period=15&start_date=" + param.StartDate + "&end_date=" + param.EndDate + "&adjust=hfq")
	res, _ := http.Get("http://127.0.0.1:8080/api/public/stock_zh_a_hist_min_em?symbol=" + param.Symbol + "&period=15&start_date=" + param.StartDate + "&end_date=" + param.EndDate + "&adjust=hfq")
	fmt.Println("http://127.0.0.1:8080/api/public/stock_zh_a_hist_min_em?symbol=" + param.Symbol + "&period=15&start_date=" + param.StartDate + "&end_date=" + param.EndDate + "&adjust=hfq")
	by, _ := ioutil.ReadAll(res.Body)
	c.String(200, string(by))
}

func Recommend(c *gin.Context) {
	fmt.Println("??????????")
	list := service.Filter_expensive_list()
	data, _ := json.Marshal(list)
	c.String(200, string(data))
}

func HotList(c *gin.Context) {
	list := service.GetMoneyList()
	data, _ := json.Marshal(list)
	c.String(200, string(data))
}
