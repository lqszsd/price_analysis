package controller

import (
	"encoding/json"
	"fmt"
	"get_price/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

type ResData struct {
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
	API   string `json:"api"`
}

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//webSocket请求ping 返回pong
func AllSymbol(c *gin.Context) {
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
		if res.API == "recommend" {
			list := service.Get_Page_Filter_expensive_list(res.Page, res.Limit)
			response, _ := json.Marshal(list)
			err = ws.WriteMessage(mt, response)
			if err != nil {
				break
			}
		}
		fmt.Println("收到消息", string(message), res)
		//写入ws数据
	}
}
