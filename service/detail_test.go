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
	Estimate := (8.340 - 5.655) * 300
	mysql_model := model.SelectModel{
		Symbol:       "002060",
		Cost:         "5.655",
		Price:        8.340,
		Name:         "粤水电",
		Shareholding: "300",
		Estimate:     Estimate,
	}
	data, _ := json.Marshal(mysql_model)
	err := re.Set(ctx, "002060", string(data), 0).Err()
	fmt.Println(err)
	err = re.HMSet(ctx, "my_select", "002060", "002060").Err()
	Estimate = (13.280 - 14.185) * 400
	mysql_model = model.SelectModel{
		Symbol:       "002639",
		Cost:         "14.185",
		Price:        13.280,
		Name:         "雪人股份",
		Shareholding: "400",
		Estimate:     Estimate,
	}
	data, _ = json.Marshal(mysql_model)
	err = re.Set(ctx, "002639", string(data), 0).Err()
	fmt.Println(err)
	err = re.HMSet(ctx, "my_select", "002639", "002639").Err()

	//fmt.Println(re.LIndex(ctx,"test",0).Result())

}
