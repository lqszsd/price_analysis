package strategy

import (
	"context"
	"encoding/json"
	"get_price/model"
	"get_price/service"
)

var ctx = context.Background()

func GetStrategyList() ([]model.Strategy, error) {
	re := service.GetRedis()
	list, err := re.HGetAll(ctx, "my_strategy").Result()
	var data []model.Strategy
	if len(list) > 0 {
		for _, v := range list {
			symbol, _ := re.Get(ctx, v).Result()
			if symbol != "" {
				info := model.Strategy{}
				_ = json.Unmarshal([]byte(symbol), &info)
				data = append(data, info)
			}
		}
	}
	return data, err
}
