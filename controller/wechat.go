package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"get_price/service"
	"github.com/gin-gonic/gin"
	wechat "github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"math/rand"
	"strconv"
	"time"
)

type ListData struct {
	Name          string  `json:"Name"`
	NowPrice      float64 `json:"NowPrice"`
	MonthAvg      float64 `json:"month_avg"`
	MonthMaxMoney float64 `json:"month_max_money"`
	MonthMinMoney float64 `json:"month_min_money"`
	MonthMsg      string  `json:"month_msg"`
	MonthStatus   int     `json:"month_status"`
	SevenAvg      float64 `json:"seven_avg"`
	SevenMaxMoney float64 `json:"seven_max_money"`
	SevenMinMoney float64 `json:"seven_min_money"`
	SevenMsg      string  `json:"seven_msg"`
	SevenStatus   int     `json:"seven_status"`
	Symbol        string  `json:"symbol"`
}
type List []ListData

func Login(c *gin.Context) {
	wc := wechat.NewWechat()
	//设置全局cache，也可以单独为每个操作实例设置
	redisOpts := &cache.RedisOpts{
		Host:     "127.0.0.1:6379",
		Password: "lq123456789",
	}
	redisCache := cache.NewRedis(redisOpts)
	wc.SetCache(redisCache)
	cfg := &offConfig.Config{
		AppID:     "wx1364b54db5971f71",
		AppSecret: "8ee11684cbbf0fd5ff66810789a15417",
		Token:     "wechat",
		//EncodingAESKey: "xxxx",
		//Cache: redisCache, //也可以单独设置
	}
	officialAccount := wc.GetOfficialAccount(cfg)
	server := officialAccount.GetServer(c.Request, c.Writer)
	var ctx = context.Background()
	server.SetMessageHandler(func(msg *message.MixMessage) *message.Reply {
		r := service.GetRedis()
		response, _ := r.Get(ctx, "wechat_recommend").Result()
		fmt.Println("这是结果", response)
		var list List
		json.Unmarshal([]byte(response), &list)
		if len(list) > 0 {
			rand.Seed(time.Now().UnixNano())
			r := rand.Intn(len(list) - 1)
			v := list[r]
			a := v.Name
			symbol := v.Symbol
			month_avg := strconv.FormatFloat(v.MonthAvg, 'f', 1, 64)
			month_max := strconv.FormatFloat(v.MonthMaxMoney, 'f', 1, 64)
			month_min := strconv.FormatFloat(v.MonthMinMoney, 'f', 1, 64)
			now_price := strconv.FormatFloat(v.NowPrice, 'f', 1, 64)
			msg_data := "名称" + a + "代码:" + symbol + ",当前价格:" + now_price + ",当月平均值" + month_avg + ",月最高" + month_max + "月最低:" + month_min
			text := message.NewText(msg_data)
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
		}
		text := message.NewText("暂时没有推荐")
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	})
	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}
	//发送回复的消息
	server.Send()
}
