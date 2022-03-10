package news

import (
	"encoding/json"
	"net/http"
)

type NewList []struct {
	Code       string `json:"code"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	PublicTime string `json:"public_time"`
	URL        string `json:"url"`
}

func GetNewList(symbol string) NewList {
	res, _ := http.Get("http://127.0.0.1:8080/api/public/stock_news_em?stock=" + symbol)
	var news NewList
	_ = json.NewDecoder(res.Body).Decode(&news)
	return news
}
