package main_force

//todo 主力成本估算
func MainForce(symbol, market string) {
	///stock_individual_fund_flow
	//第一天 (流入-流出) *当日平均值+第n天 (流入-流出) *(如果今天大于昨天那么出货应该取负数)当日平均值 除以总数量 大概可以得出一个建仓成本
	//数量累加求和多少手
}
