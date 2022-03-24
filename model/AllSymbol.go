package model

type AllSymbol []struct {
	ID               int         `json:"序号"`
	Code             string      `json:"代码"`
	Name             string      `json:"名称"`
	Price            float64     `json:"最新价"`
	FluctuationRange float64     `json:"涨跌幅"`
	RiseAndFall      float64     `json:"涨跌额"`
	TurnoverNum      float64     `json:"成交量"`
	Turnover         float64     `json:"成交额"`
	Amplitude        float64     `json:"振幅"`
	Max              float64     `json:"最高"`
	Min              float64     `json:"最低"`
	Begin            float64     `json:"今开"`
	End              float64     `json:"昨收"`
	Percent          interface{} `json:"量比"`
	ChangeHand       float64     `json:"换手率"`
	Dynamic          float64     `json:"市盈率-动态"`
	Dynamics         float64     `json:"市净率"`
}

type InstitutionalRecommendation struct {
	Code            string  `json:"股票代码"`
	Name            string  `json:"股票名称"`
	Grade           string  `json:"最新评级"`
	FuturePrice     float64 `json:"目标价"`
	GradeTime       string  `json:"评级日期"`
	AvgGrade        string  `json:"综合评级"`
	AverageIncrease string  `json:"平均涨幅"`
	Industry        string  `json:"行业"`
}
type Institutional []InstitutionalRecommendation
