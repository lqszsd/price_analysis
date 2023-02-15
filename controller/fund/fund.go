package fund

import (
	"encoding/json"
	"fmt"
	"get_price/service/fund"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

type ResData struct {
	Page     int    `json:"page"`
	Limit    int    `json:"limit"`
	API      string `json:"api"`
	FundName string `json:"fund_name"`
	FundCode string `json:"fund_code"`
}

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Index(c *gin.Context) {
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	for {
		//读取ws中的数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		var res ResData
		json.Unmarshal(message, &res)
		if res.API == "fund_list" {
			list := fund.Get_fund_list(res.Page, res.Limit, res.FundCode, res.FundName)
			response, _ := json.Marshal(list)
			fmt.Println("回复消息", response)
			err = ws.WriteMessage(mt, response)
			if err != nil {
				break
			}
		}
		if res.API == "fund_new" {
			list := fund.Get_fund_new_list(res.Page, res.Limit, res.FundCode, res.FundName)
			response, _ := json.Marshal(list)
			fmt.Println("回复消息", response)
			err = ws.WriteMessage(mt, response)
			if err != nil {
				break
			}
		}
		fmt.Println("收到消息", string(message), res)
		//写入ws数据
	}
}
