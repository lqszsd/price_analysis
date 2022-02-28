package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

//http://127.0.0.1:8080/api/public/stock_hot_rank_em

type MoneyList []struct {
	Rank   int     `json:"当前排名"`
	Code   string  `json:"代码"`
	Name   string  `json:"股票名称"`
	Price  float64 `json:"最新价"`
	Change float64 `json:"涨跌幅"`
}

type MoneyTileList []struct {
	Day           string  `json:"日期"`
	BeginPrice    float64 `json:"开盘"`
	EndPrice      float64 `json:"收盘"`
	MaxPrice      float64 `json:"最高"`
	MinProce      float64 `json:"最低"`
	SuccessMember int     `json:"成交量"`
	SuccessPrice  int     `json:"成交额"`
	Amplitude     float64 `json:"振幅"`
	ChangePercent float64 `json:"涨跌幅"`
	ChangePrice   float64 `json:"涨跌额"`
	ChangeHand    float64 `json:"换手率"`
}

func Filter_expensive_list() []map[string]interface{} {
	var real_list []map[string]interface{}
	real_list = make([]map[string]interface{}, 0)
	list := GetMoneyList()
	if len(list) > 0 {
		for _, g := range list {
			v := g
			//不考虑花费大于15的list
			if v.Price > 15 {
				continue
			}
			v.Code = strings.Replace(v.Code, "SZ", "", -1)
			v.Code = strings.Replace(v.Code, "SH", "", -1)
			v.Code = strings.Replace(v.Code, "sz", "", -1)
			v.Code = strings.Replace(v.Code, "sh", "", -1)
			data := GetMoneyRecommend(v.Code, v.Price)
			data["Name"] = v.Name
			data["NowPrice"] = v.Price
			data["symbol"] = v.Code
			real_list = append(real_list, data)
		}
	}
	return real_list
}

func GetMoneyList() MoneyList {
	resp, _ := http.Get("http://127.0.0.1:8080/api/public/stock_hot_rank_em")
	by, _ := ioutil.ReadAll(resp.Body)
	var money MoneyList
	json.Unmarshal(by, &money)
	return money
}

func getPriceInfo(symbol string, last_day, today string) map[string]float64 {
	res, _ := http.Get("http://127.0.0.1:8080/api/public/stock_zh_a_hist?symbol=" + symbol + "&period=daily&start_date=" + last_day + "&end_date=" + today)
	var money_time MoneyTileList
	err := json.NewDecoder(res.Body).Decode(&money_time)
	var min_money, max_money float64
	var response_data map[string]float64
	response_data = make(map[string]float64)
	fmt.Println(money_time, err, "http://127.0.0.1:8080/api/public/stock_zh_a_hist?symbol="+symbol+"&period=daily&start_date="+last_day+"&end_date="+today)
	if len(money_time) > 0 {
		for _, v := range money_time {
			if v.MaxPrice >= max_money {
				max_money = v.MaxPrice
			}
			if min_money < 1 {
				min_money = v.MinProce
			}
			if v.MinProce < min_money {
				min_money = v.MinProce
			}
		}
		avg := (min_money + max_money) / 2
		//如果当前价格大于7日线 短线不可买入
		response_data["min_money"] = min_money
		response_data["max_money"] = max_money
		response_data["avg"] = avg
	}
	return response_data
}
func GetMoneyRecommend(symbol string, price_now float64) map[string]interface{} {
	//判断是不是短线
	//获取上周一
	var response_data map[string]interface{}
	response_data = make(map[string]interface{})
	last_day := GetLastWeekFirstDate()
	now := time.Now()
	today := now.Format("20060102")
	price_info := getPriceInfo(symbol, last_day, today)
	//如果当前价格大于7日线 短线不可买入
	if price_info["avg"] > price_now {
		response_data["seven_status"] = 1
		response_data["seven_msg"] = "当前价格在七日线低点可以买入"
	} else {
		response_data["seven_status"] = 0
		response_data["seven_max_msg"] = "当前价格在七日线高点不可以买入"
	}
	response_data["seven_min_money"] = price_info["min_money"]
	response_data["seven_max_money"] = price_info["max_money"]
	response_data["seven_avg"] = price_info["avg"]

	back_month := now.AddDate(0, -1, 0)
	last_month := back_month.AddDate(0, -1, 0)
	tow_month_ago := last_month.Format("20060102")
	price_info = getPriceInfo(symbol, tow_month_ago, today)
	if price_info["avg"] > price_now {
		response_data["month_status"] = 1
		response_data["month_msg"] = "当前价格在2个月内低于平均值可以考虑买入"
	} else {
		response_data["month_status"] = 0
		response_data["month_msg"] = "当前价格在2个月内高于平均值不推荐买入"
	}
	response_data["month_min_money"] = price_info["min_money"]
	response_data["month_max_money"] = price_info["max_money"]
	response_data["month_avg"] = price_info["avg"]
	return response_data
}

/**
获取本周周一的日期
*/
func GetFirstDateOfWeek() (weekMonday string) {
	now := time.Now()

	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}

	weekStartDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	weekMonday = weekStartDate.Format("2006-01-02")
	return
}

/**
获取上周的周一日期
*/
func GetLastWeekFirstDate() (weekMonday string) {
	thisWeekMonday := GetFirstDateOfWeek()
	TimeMonday, _ := time.Parse("2006-01-02", thisWeekMonday)
	lastWeekMonday := TimeMonday.AddDate(0, 0, -7)
	weekMonday = lastWeekMonday.Format("20060102")
	return
}
