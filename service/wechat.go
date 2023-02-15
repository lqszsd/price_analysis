package service

import "strings"

func GetWechatRecommend() []map[string]interface{} {
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
			if data["seven_msg"] == "当前价格在七日线低点可以买入" && data["month_msg"] == "当前价格在2个月内低于平均值可以考虑买入" {
				real_list = append(real_list, data)
			}
		}
	}
	return real_list
}
