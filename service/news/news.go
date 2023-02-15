package news

import (
	"encoding/json"
	"net/http"
)

type NewList []struct {
	Code       string `json:"关键词"`
	Title      string `json:"新闻标题"`
	Content    string `json:"新闻内容"`
	PublicTime string `json:"发布时间"`
	URL        string `json:"新闻链接"`
}

func GetNewList(symbol string) NewList {
	res, _ := http.Get("http://127.0.0.1:8080/api/public/stock_news_em?symbol=" + symbol)
	var news NewList
	_ = json.NewDecoder(res.Body).Decode(&news)
	return news
}
