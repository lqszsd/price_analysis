package fund

import (
	"encoding/json"
	"fmt"
	"get_price/service/SqlliteDb"
	"net/http"
)

type FundMoney []struct {
	Id                int     `json:"序号"`
	FundCode          string  `json:"基金代码"`
	FundName          string  `json:"基金简称"`
	AverageNAV        float64 `json:"单位净值"`
	TotalRaisingScale float64 `json:"总募集规模"`
	RecentTotalShare  float64 `json:"最近总份额"`
	BeginDay          string  `json:"成立日期"`
	FundManager       string  `json:"基金经理"`
	UpdateTime        string  `json:"更新日期"`
	TotalPrice        float64 `gorm:"column:total_price"`
}

var FundMoneySql = `
CREATE TABLE fund_money (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "fund_code" VARCHAR(64),
  "fund_name" VARCHAR(64),
  "average_nav" real(64),
  "total_raising_scale" VARCHAR(64),
  "recent_total_share" VARCHAR(64),
  "begin_day" VARCHAR(64),
  "fund_manager" VARCHAR(64),
  "total_price" real(64),
  "update_time" VARCHAR(64)
);
CREATE INDEX "fund_code"
ON "fund_money" (
  "fund_code"
);
CREATE INDEX "fund_name"
ON "fund_money" (
  "fund_name"
);
`

type FundMoneyTable struct {
	ID                int64   `gorm:"column:id"`
	FundCode          string  `gorm:"column:fund_code"`
	FundName          string  `gorm:"column:fund_name"`
	AverageNav        string  `gorm:"column:average_nav"`
	TotalRaisingScale string  `gorm:"column:total_raising_scale"`
	RecentTotalShare  string  `gorm:"column:recent_total_share"`
	BeginDay          string  `gorm:"column:begin_day"`
	FundManager       string  `gorm:"column:fund_manager"`
	TotalPrice        float64 `gorm:"column:total_price"`
	UpdateTime        string  `gorm:"column:update_time"`
}

func (FundMoney) TableName() string {
	return "fund_money"
}

func Fund_data_money() {
	db := SqlliteDb.OpenTableDB("fund_money")
	db.Exec("Drop table fund_money")

	db.Exec(FundMoneySql)
	res, _ := http.Get("http://127.0.0.1:8080/api/public/fund_scale_open_sina")
	var fundmoney FundMoney
	err := json.NewDecoder(res.Body).Decode(&fundmoney)
	if err != nil {
		return
	}
	if len(fundmoney) > 0 {
		for _, s := range fundmoney {
			fmt.Println(s, s.FundCode)
			fund_money := FundMoneyTable{}
			db.Table("fund_money").Where("fund_code=?", s.FundCode).First(&fund_money)
			if fund_money.ID > 0 {
				continue
			} else {
				s.TotalPrice = s.RecentTotalShare * s.AverageNAV
				db.Table("fund_money").Debug().Create(&s)
			}
		}
	}
}

func Get_fund_list(page int, limit int, fundCode, fundName string) map[string]interface{} {
	db := SqlliteDb.OpenTableDB("fund_money")
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * limit
	type ListData struct {
		FundMoneyTable
		ThreeMonth string `json:"three_month"`
		SixMonth   string `json:"six_month"`
		Year       string `json:"year"`
		TwoYear    string `json:"two_year"`
	}
	var list []ListData
	temp := db.Table("fund_money").Debug().Offset(offset).Limit(limit).Order("total_price desc")
	if fundName != "" {
		temp.Where("fund_name=?", fundName)
	}
	if fundCode != "" {
		temp.Where("fund_code=?", fundCode)
	}
	temp.Find(&list)
	if len(list) > 0 {
		fund_base_db := SqlliteDb.OpenTableDB("fund_base_data")
		for g, s := range list {
			fund_base_info := FundBaseDataTable{}
			fund_base_db.Table("fund_base_data").Where("fund_name=?", s.FundName).First(&fund_base_info)
			if fund_base_info.ID > 0 {
				s.ThreeMonth = fund_base_info.ThreeMonth
				s.SixMonth = fund_base_info.SixMonth
				s.Year = fund_base_info.Year
				s.TwoYear = fund_base_info.TwoYear
			}
			list[g] = s
		}
	}
	var data map[string]interface{}
	data = make(map[string]interface{})
	data["list"] = list
	data["count"] = 2000
	return data
}

func Get_fund_new_list(page int, limit int, fundCode, fundName string) map[string]interface{} {
	db := SqlliteDb.OpenTableDB("fund_new")
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * limit
	type ListData struct {
		FundNew
		ThreeMonth string `json:"three_month"`
		SixMonth   string `json:"six_month"`
		Year       string `json:"year"`
		TwoYear    string `json:"two_year"`
	}
	var list []ListData
	temp := db.Table("fund_new").Debug().Offset(offset).Limit(limit).Order("start_time desc")
	if fundName != "" {
		temp.Where("fund_name=?", fundName)
	}
	if fundCode != "" {
		temp.Where("fund_code=?", fundCode)
	}
	temp.Find(&list)
	if len(list) > 0 {
		fund_base_db := SqlliteDb.OpenTableDB("fund_base_data")
		for g, s := range list {
			fund_base_info := FundBaseDataTable{}
			fund_base_db.Table("fund_base_data").Where("fund_name=?", s.FundName).First(&fund_base_info)
			if fund_base_info.ID > 0 {
				s.ThreeMonth = fund_base_info.ThreeMonth
				s.SixMonth = fund_base_info.SixMonth
				s.Year = fund_base_info.Year
				s.TwoYear = fund_base_info.TwoYear
			}
			list[g] = s
		}
	}
	var data map[string]interface{}
	data = make(map[string]interface{})
	data["list"] = list
	data["count"] = 2000
	return data
}
