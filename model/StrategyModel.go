package model

type Strategy struct {
	Symbol            string  `json:"symbol"` //代码
	Name              string  `json:"name"`
	PlanCopyTheBottom string  `json:"plan_copy_the_bottom"` //计划抄底价
	PlanSale          string  `json:"plan_sale"`            //计划卖出价
	Status            bool    `json:"status"`               //0准备抄底，1抄底成功，2已卖出
	Price             float64 `json:"price"`                //当前价格
	EmailCount        int     `json:"email_count"`          //发送邮件次数
}
