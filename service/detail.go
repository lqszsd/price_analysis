package service

import (
	"encoding/json"
	"fmt"
	"get_price/request"
	"strings"
)

type DetailResponse struct {
	Fundcode string `json:"fundcode"`
	Name     string `json:"name"`
	Jzrq     string `json:"jzrq"`
	Dwjz     string `json:"dwjz"`
	Gsz      string `json:"gsz"`
	Gszzl    string `json:"gszzl"`
	Gztime   string `json:"gztime"`
}
type Detail interface {
	Do(url string) string
	Replace(data string) string
	Parse(data string) DetailResponse
	Request(url string) DetailResponse
}

func (d DetailResponse)Do(url string) string {
	data:=request.HttpGet(url)
	return string(data)
}

func (d DetailResponse)Replace(data string) string {
	str:=strings.Replace(data,"jsonpgz(","",-1)
	str=strings.Replace(str,");","",-1)
	return str
}
func (d DetailResponse)Parse(data string) DetailResponse {
	byte:=[]byte(data)
	json.Unmarshal(byte,&d)
	return d
}
func (d DetailResponse)Request(url string) DetailResponse {
	str:=d.Do(url)
	fmt.Println(str)
	str=d.Replace(str)
	fmt.Println(str)
	return d.Parse(str)
}