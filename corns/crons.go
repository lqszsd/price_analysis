package corns

import (
	"context"
	"encoding/json"
	"fmt"
	"get_price/model"
	"get_price/service"
	"github.com/robfig/cron/v3"
	"net/http"
	"strconv"
	"time"
)

var corns *cron.Cron
var ctx = context.Background()

func RegisterCrons() {
	if corns == nil {
		corns = cron.New()
	}
	fmt.Println("启动定时任务")
	corns.AddFunc("*/5 * * * *", func() {
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
						res, _ := http.Get("http://127.0.0.1:8080/api/public/stock_zh_a_hist_min_em?symbol=" + info.Symbol + "&period=15&start_date=" + logDay + "&end_date=" + now)
						fmt.Println("http://127.0.0.1:8080/api/public/stock_zh_a_hist_min_em?symbol=" + info.Symbol + "&period=15&start_date=" + logDay + "&end_date=" + now)
						symbolList := model.SymbolList{}
						json.NewDecoder(res.Body).Decode(&symbolList)
						if len(symbolList) > 0 {
							info.Price = symbolList[len(symbolList)-1].End
							cost, _ := strconv.ParseFloat(info.Cost, 64)
							shareholding, _ := strconv.ParseFloat(info.Shareholding, 64)
							info.Estimate = (info.Price - cost) * shareholding
							data, _ := json.Marshal(info)
							err := re.Set(ctx, info.Symbol, string(data), 0).Err()
							fmt.Println("更新当前价格and预估收益", err, info.Estimate, info.Price)
						}

					}
				}
			}
		}
	})
	corns.Start()

}
