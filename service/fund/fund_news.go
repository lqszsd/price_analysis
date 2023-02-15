package fund

import (
	"encoding/json"
	"fmt"
	"get_price/service/SqlliteDb"
	"net/http"
)

type FundNew struct {
	FundCode    string  `json:"基金代码"`
	FundName    string  `json:"基金简称"`
	FundCompany string  `json:"发行公司"`
	FundType    string  `json:"基金类型"`
	BuyTime     string  `json:"集中认购期"`
	Total       float64 `json:"募集份额"`
	StartTime   string  `json:"成立日期"`
	BeginNow    float64 `json:"成立来涨幅"`
	FundManager string  `json:"基金经理"`
	Status      string  `json:"申购状态"`
	BuyPercent  float64 `json:"优惠费率"`
}

var FundNewSql = `
CREATE TABLE IF NOT EXISTS fund_new (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "fund_code" VARCHAR(64),
  "fund_name" VARCHAR(64),
  "fund_company" VARCHAR(64),
  "fund_type" VARCHAR(64),
  "buy_time" VARCHAR(64),
  "total" VARCHAR(64),
  "start_time" VARCHAR(64),
  "begin_now" real(64),
  "fund_manager" VARCHAR(64),
"status" VARCHAR(64),
"buy_percent" VARCHAR(64)
);
CREATE INDEX "fund_new"
ON "fund_money" (
  "fund_code"
);
CREATE INDEX "fund_new"
ON "fund_money" (
  "fund_name"
);
`

func Fund_data_new() {
	db := SqlliteDb.OpenTableDB("fund_new")
	db.Exec("Drop table fund_new")

	db.Exec(FundNewSql)
	res, _ := http.Get("http://127.0.0.1:8080/api/public/fund_new_found_em")
	var fundnew []FundNew
	err := json.NewDecoder(res.Body).Decode(&fundnew)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(fundnew) > 0 {
		for _, s := range fundnew {
			fmt.Println(s, s.FundCode)
			fund_money := FundMoneyTable{}
			db.Table("fund_new").Where("fund_code=?", s.FundCode).First(&fund_money)
			fmt.Println("???????????????????", s)
			if fund_money.ID > 0 {
				continue
			} else {
				db.Table("fund_new").Debug().Create(&s)
			}
		}
	}
}
