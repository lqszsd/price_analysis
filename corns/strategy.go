package corns

import (
	"encoding/json"
	"fmt"
	"get_price/model"
	"get_price/service"
	"get_price/service/email"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func CronStrategy() {
	re := service.GetRedis()
	list, err := re.HGetAll(ctx, "my_strategy").Result()
	if err != nil {
		return
	}
	fmt.Println(list)
	if len(list) > 0 {
		for _, v := range list {
			symbol, _ := re.Get(ctx, v).Result()
			if symbol != "" {
				info := model.Strategy{}
				_ = json.Unmarshal([]byte(symbol), &info)
				if info.Symbol != "" {
					info.Symbol = strings.Replace(info.Symbol, "-lq", "", -1)
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
						plan_copy_the_bottom, _ := strconv.ParseFloat(info.PlanCopyTheBottom, 64)
						//如果查询到当前的价格低于计划抄底价 而且邮件发送次数不超过5次就发送邮件
						price := strconv.FormatFloat(info.Price, 'E', -1, 64)
						if info.Price <= plan_copy_the_bottom {
							count_num, _ := re.Get(ctx, info.Symbol+"-plan_copy_the_bottom").Result()
							num, _ := strconv.Atoi(count_num)
							if count_num == "" || num <= 3 {
								fmt.Println("1111111111")
								num++
								email.SetMail(info.Name+"可以抄底当前价格为:"+price+"计划买入价格为:"+info.PlanCopyTheBottom, info.Symbol, info.Symbol)
								re.Set(ctx, info.Symbol+"-plan_copy_the_bottom", num, 0)
							} else {
								fmt.Println("2222222222")
							}
						}
						plan_sale, _ := strconv.ParseFloat(info.PlanSale, 64)
						if info.Price >= plan_sale {
							count_num, _ := re.Get(ctx, info.Symbol+"-plan_copy_the_bottom").Result()
							num, _ := strconv.Atoi(count_num)
							if count_num == "" || num <= 3 {
								fmt.Println("33333333333")
								num++
								email.SetMail(info.Name+"可以冒出当前价格为:"+price+"计划卖出价格为:"+info.PlanSale, info.Symbol, info.Symbol)
								re.Set(ctx, info.Symbol+"-plan_sale", num, 0)
							} else {
								fmt.Println("4444444444")
							}
						}

					}

				}
			}
		}
	}
}
