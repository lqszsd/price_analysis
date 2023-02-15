package fund

import (
	"encoding/json"
	"fmt"
	"get_price/service/SqlliteDb"
	"net/http"
)

type FundBaseData []struct {
	ID              int    `json:"序号"`
	FundCode        string `json:"基金代码"`
	FundName        string `json:"基金简称"`
	Day             string `json:"日期"`
	AverageNav      string `json:"单位净值"`
	AccumulatedNet  string `json:"累计净值"`
	DailyGrowthRate string `json:"日增长率"`
	Week            string `json:"近1周"`
	Month           string `json:"近1月"`
	ThreeMonth      string `json:"近3月"`
	SixMonth        string `json:"近6月"`
	Year            string `json:"近1年"`
	TwoYear         string `json:"近2年"`
	ThreeYear       string `json:"近3年"`
	BeginYear       string `json:"今年来"`
	BeginNow        string `json:"成立来"`
	DefaultData     string `json:"自定义"`
	Price           string `json:"手续费"`
}

type FundBaseDataTable struct {
	ID              int64  `gorm:"column:id"`
	FundCode        string `gorm:"column:fund_code"`
	FundName        string `gorm:"column:fund_name"`
	Day             string `gorm:"column:day"`
	AverageNav      string `gorm:"column:average_nav"`
	AccumulatedNet  string `gorm:"column:accumulated_net"`
	DailyGrowthRate string `gorm:"column:daily_growth_rate"`
	Week            string `gorm:"column:week"`
	Month           string `gorm:"column:month"`
	ThreeMonth      string `gorm:"column:three_month"`
	SixMonth        string `gorm:"column:six_month"`
	Year            string `gorm:"column:year"`
	TwoYear         string `gorm:"column:two_year"`
	ThreeYear       string `gorm:"column:three_year"`
	BeginYear       string `gorm:"column:begin_year"`
	BeginNow        string `gorm:"column:begin_now"`
	DefaultData     string `gorm:"column:default_data"`
	Price           string `gorm:"column:price"`
}

func (FundBaseDataTable) TableName() string {
	return "fund_base_data"
}

var sql_table = `
CREATE TABLE IF NOT EXISTS fund_base_data(
   id INTEGER PRIMARY KEY AUTOINCREMENT,
   fund_code VARCHAR(64) NULL,
   fund_name VARCHAR(64) NULL,
day VARCHAR(64) NULL,
average_nav VARCHAR(64) NULL,
accumulated_net VARCHAR(64) NULL,
daily_growth_rate VARCHAR(64) NULL,
week VARCHAR(64) NULL,
month VARCHAR(64) NULL,
three_month VARCHAR(64) NULL,
six_month VARCHAR(64) NULL,
year VARCHAR(64) NULL,
two_year VARCHAR(64) NULL,
three_year VARCHAR(64) NULL,
begin_year VARCHAR(64) NULL,
begin_now VARCHAR(64) NULL,
default_data VARCHAR(64) NULL,
price VARCHAR(64) NULL
);
CREATE INDEX "fund_code"
ON "fund_base_data" (
  "fund_code"
);
CREATE INDEX "fund_name"
ON "fund_base_data" (
  "fund_name"
);
`

func Fund_data_init() {
	db := SqlliteDb.OpenTableDB("fund_base_data")
	db.Exec("Drop table fund_base_data")
	db.Exec(sql_table)
	res, _ := http.Get("http://127.0.0.1:8080/api/public/fund_open_fund_rank_em")
	var fundbasedata FundBaseData
	_ = json.NewDecoder(res.Body).Decode(&fundbasedata)
	if len(fundbasedata) > 0 {
		for _, s := range fundbasedata {
			fmt.Println(s, s.FundCode)
			fund_base := FundBaseDataTable{}

			db.Table("fund_base_data").Where("fund_code=?", s.FundCode).First(&fund_base)
			if fund_base.ID > 0 {
				continue
			} else {
				db.Table("fund_base_data").Debug().Create(&s)
			}
		}
	}
}

func Register_fund_data() {
	go Fund_data_money()
	go Fund_data_init()
	go Fund_data_new()
}
