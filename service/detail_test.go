package service

import (
	"context"
	"encoding/json"
	"fmt"
	"get_price/model"
	"testing"
)

func TestDetailTest(t *testing.T) {
	var d Detail
	d = new(DetailResponse)
	//招商白酒
	d.Request("https://fundgz.1234567.com.cn/js/161725.js?rt=ZSZZBJZSLOFA")
	fmt.Println(d)
}

func TestGetRedis(t *testing.T) {
	var ctx = context.Background()
	re := GetRedis()
	Estimate := (6.22 - 6.12) * 100
	mysql_model := model.SelectModel{
		Symbol:       "002822",
		Cost:         "6.22",
		Price:        6.12,
		Name:         "中装建设",
		Shareholding: "100",
		Estimate:     Estimate,
	}
	data, _ := json.Marshal(mysql_model)
	err := re.Set(ctx, "002822", string(data), 0).Err()
	fmt.Println(err)
	err = re.HMSet(ctx, "my_select", "002822", "002822").Err()
	re.HDel(ctx, "my_select", "601868")
	//fmt.Println(re.LIndex(ctx,"test",0).Result())

}
