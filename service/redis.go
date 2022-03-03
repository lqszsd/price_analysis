package service

import (
	"context"
	"encoding/json"
	"get_price/model"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var re *redis.Client

func GetRedis() *redis.Client {
	if re == nil {
		re = redis.NewClient(&redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})
	}
	return re
}

func GetSelectList() ([]model.SelectModel, error) {
	re := GetRedis()
	list, err := re.HGetAll(ctx, "my_select").Result()
	var data []model.SelectModel
	if len(list) > 0 {
		for _, v := range list {
			symbol, _ := re.Get(ctx, v).Result()
			if symbol != "" {
				info := model.SelectModel{}
				_ = json.Unmarshal([]byte(symbol), &info)
				data = append(data, info)
			}
		}
	}
	return data, err
}
