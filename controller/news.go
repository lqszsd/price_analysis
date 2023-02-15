package controller

import (
	"encoding/json"
	"fmt"
	"get_price/service/news"
	"github.com/gin-gonic/gin"
	"strings"
)

func NewList(c *gin.Context) {
	symbol := c.DefaultQuery("symbol", "002060")
	symbol = strings.Replace(symbol, "SZ", "", -1)
	symbol = strings.Replace(symbol, "SH", "", -1)
	symbol = strings.Replace(symbol, "sz", "", -1)
	symbol = strings.Replace(symbol, "sh", "", -1)
	fmt.Println("444444444444", symbol)
	list := news.GetNewList(symbol)
	data, _ := json.Marshal(list)
	c.String(200, string(data))
}
