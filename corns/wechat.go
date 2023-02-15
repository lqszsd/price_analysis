package corns

import (
	"context"
	"encoding/json"
	"fmt"
	"get_price/service"
)

func CronWechat() {
	r := service.GetRedis()
	response, _ := r.Get(ctx, "wechat_recommend").Result()
	fmt.Println("这是结果", response)
	list := service.GetWechatRecommend()
	if len(list) > 0 {
		var ctx = context.Background()
		data, _ := json.Marshal(list)
		err := r.Set(ctx, "wechat_recommend", string(data), 0).Err()
		fmt.Println(err)
		response, _ := r.Get(ctx, "wechat_recommend").Result()
		fmt.Println("这是结果", response)
	} else {
		fmt.Println("真的没有推荐")
	}
	//if len(list)==0{
	//	bd.SendText(nil, "今日无推荐")
	//
	//}
	//for _,v:=range list{
	//	a := v["Name"].(string)
	//	symbol:=v["symbol"].(string)
	//	month_avg:=v["seven_avg"].(string)
	//	month_max:=v["seven_max_money"].(string)
	//	month_min:=v["seven_min_money"].(string)
	//	msg_data:="名称:"+a+"代码:"+symbol+"当月平均值"+month_avg+",月最高"+month_max+"月最低:"+month_min
	//	bd.SendText(nil, msg_data)
	//}
	//处理消息接收以及回复
}
