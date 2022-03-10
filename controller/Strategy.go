package controller

import (
	"encoding/json"
	"fmt"
	"get_price/model"
	"get_price/service"
	"get_price/service/strategy"
	"github.com/gin-gonic/gin"
	"strings"
)

func AddStrategy(c *gin.Context) {
	var person model.Strategy
	c.ShouldBind(&person)
	if person.Symbol != "" {
		person.Symbol = strings.Replace(person.Symbol, "SZ", "", -1)
		person.Symbol = strings.Replace(person.Symbol, "SH", "", -1)
		person.Symbol = strings.Replace(person.Symbol, "sz", "", -1)
		person.Symbol = strings.Replace(person.Symbol, "sh", "", -1)
		person.Symbol = person.Symbol + "-lq"
		redis := service.GetRedis()
		person.Status = false
		data, _ := json.Marshal(person)
		err := redis.Set(ctx, person.Symbol, string(data), 0).Err()
		fmt.Println(1111111111, err)
		err = redis.HMSet(ctx, "my_strategy", person.Symbol, person.Symbol).Err()
		fmt.Println(22222222, err)
		response, _ := json.Marshal(Success(person))
		c.String(200, string(response))
		return
	}
	response, _ := json.Marshal(Error(person))
	c.String(200, string(response))
	return
}

func UserStrategy(c *gin.Context) {
	list, _ := strategy.GetStrategyList()
	data, _ := json.Marshal(list)
	c.String(200, string(data))
}

func ChangeStrategyStatus(c *gin.Context) {
	var person model.Strategy
	c.ShouldBind(&person)
	if person.Symbol != "" {
		person.Symbol = strings.Replace(person.Symbol, "SZ", "", -1)
		person.Symbol = strings.Replace(person.Symbol, "SH", "", -1)
		person.Symbol = strings.Replace(person.Symbol, "sz", "", -1)
		person.Symbol = strings.Replace(person.Symbol, "sh", "", -1)
		person.Symbol = person.Symbol
		redis := service.GetRedis()
		data, _ := json.Marshal(person)
		err := redis.Set(ctx, person.Symbol, string(data), 0).Err()
		fmt.Println(1111111111, err)
		err = redis.HMSet(ctx, "my_strategy", person.Symbol, person.Symbol).Err()
		fmt.Println(22222222, err)
		response, _ := json.Marshal(Success(person))
		c.String(200, string(response))
		return
	}
	response, _ := json.Marshal(Error(person))
	c.String(200, string(response))
	return
}
