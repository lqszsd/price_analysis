package corns

import (
	"encoding/json"
	"fmt"
	"get_price/model"
	"get_price/service"
	"net/http"
	"strconv"
	"time"
)

func SelectData() {
	re := service.GetRedis()
	list, err := re.HGetAll(ctx, "my_select").Result()
	if err != nil {
		return
	}
	fmt.Println(list)
	if len(list) > 0 {
		for _, v := range list {
			symbol, _ := re.Get(ctx, v).Result()
			if symbol != "" {
				info := model.SelectModel{}
				_ = json.Unmarshal([]byte(symbol), &info)
				if info.Symbol != "" {
					fmt.Println(info.Symbol)
					nTime := time.Now()
					start_time := time.Now().Add(-time.Minute * 10)
					logDay := start_time.Format("2006-01-02")
					now := nTime.Format("2006-01-02")
					res, err := http.Get("http://127.0.0.1:8080/api/public/stock_zh_a_hist_min_em?symbol=" + info.Symbol + "&period=15&start_date=" + logDay + "&end_date=" + now)
					fmt.Println("http://127.0.0.1:8080/api/public/stock_zh_a_hist_min_em?symbol=" + info.Symbol + "&period=15&start_date=" + logDay + "&end_date=" + now)
					symbolList := model.SymbolList{}
					if err != nil {
						return
					}
					json.NewDecoder(res.Body).Decode(&symbolList)
					if len(symbolList) > 0 {
						info.Price = symbolList[len(symbolList)-1].End
						cost, _ := strconv.ParseFloat(info.Cost, 64)
						shareholding, _ := strconv.ParseFloat(info.Shareholding, 64)
						if shareholding < 1 {
							info.Estimate = 0
						} else {
							info.Estimate = (info.Price - cost) * shareholding
						}
						data, _ := json.Marshal(info)
						err := re.Set(ctx, info.Symbol, string(data), 0).Err()
						fmt.Println("更新当前价格and预估收益", err, info.Estimate, info.Price)
					}

				}
			}
		}
	}
}
