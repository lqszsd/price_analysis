package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"get_price/model"
	"get_price/service"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var ctx = context.Background()

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
	fmt.Println("http://127.0.0.1:8080/api/public/stock_zh_a_hist_min_em?symbol=" + param.Symbol + "&period=15&start_date=" + param.StartDate + "&end_date=" + param.EndDate)
	res, _ := http.Get("http://127.0.0.1:8080/api/public/stock_zh_a_hist_min_em?symbol=" + param.Symbol + "&period=15&start_date=" + param.StartDate + "&end_date=" + param.EndDate)
	fmt.Println("http://127.0.0.1:8080/api/public/stock_zh_a_hist_min_em?symbol=" + param.Symbol + "&period=15&start_date=" + param.StartDate + "&end_date=" + param.EndDate)
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

func UserSelect(c *gin.Context) {
	list, _ := service.GetSelectList()
	data, _ := json.Marshal(list)
	c.String(200, string(data))
}

func AddSelect(c *gin.Context) {
	var person model.SelectModel
	c.ShouldBind(&person)
	if person.Symbol != "" {
		person.Symbol = strings.Replace(person.Symbol, "SZ", "", -1)
		person.Symbol = strings.Replace(person.Symbol, "SH", "", -1)
		person.Symbol = strings.Replace(person.Symbol, "sz", "", -1)
		person.Symbol = strings.Replace(person.Symbol, "sh", "", -1)
		redis := service.GetRedis()
		price := person.Price
		cost, _ := strconv.ParseFloat(person.Cost, 64)
		Estimate := (price - cost) * 400
		person.Estimate = Estimate
		data, _ := json.Marshal(person)
		err := redis.Set(ctx, person.Symbol, string(data), 0).Err()
		fmt.Println(1111111111, err)
		err = redis.HMSet(ctx, "my_select", person.Symbol, person.Symbol).Err()
		fmt.Println(22222222, err)
		response, _ := json.Marshal(Success(person))
		c.String(200, string(response))
		return
	}
	response, _ := json.Marshal(Error(person))
	c.String(200, string(response))
	return
}
