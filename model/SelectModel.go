package model

type SelectModel struct {
	Symbol       string  `json:symbol`
	Cost         string  `json:cost`
	Price        float64 `json:price`
	Name         string  `json:name`
	Shareholding string  `json:shareholding`
	Estimate     float64 `json:estimate`
}

type List struct {
	Time         string  `json:"时间"`
	Begin        float64 `json:"开盘"`
	End          float64 `json:"收盘"`
	Max          float64 `json:"最高"`
	Min          float64 `json:"最低"`
	Percent      float64 `json:"涨跌幅"`
	PercentPrice float64 `json:"涨跌额"`
	Num          int     `json:"成交量"`
	NumPrice     float64 `json:"成交额"`
	Amplitude    float64 `json:"振幅"`
	TurnoverRate float64 `json:"换手率"`
}
type SymbolList []List
