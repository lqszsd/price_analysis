package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"get_price/model"
	"get_price/service"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

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

func RmSelect(c *gin.Context) {
	symbol := c.Query("symbol")
	var ctx = context.Background()
	re := service.GetRedis()
	re.HDel(ctx, "my_select", symbol, symbol)

	data := Success(symbol)
	response, _ := json.Marshal(Error(data))
	c.String(200, string(response))
	return
}
