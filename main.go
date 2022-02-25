package main

import (
	"encoding/json"
	"fmt"
	"get_price/service"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main(){
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		resBody:=request.Body

		param := &struct {
			Symbol string `json:"symbol"`
			StartDate string `json:"start_date"`
			EndDate string `json:"end_date"`
		}{}
		// 通过json解析器解析参数
		json.NewDecoder(resBody).Decode(param)
		fmt.Println(param)
		writer.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
		writer.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		writer.Header().Set("content-type", "application/json")             //返回数据格式是json
		res,_:=http.Get("http://127.0.0.1:8080/api/bond_zh_hs_cov_min?symbol="+param.Symbol+"&start_date="+param.StartDate+"&end_date="+param.EndDate+"&adjust=hfq")
		fmt.Println("http://127.0.0.1:8080/api/bond_zh_hs_cov_min?symbol="+param.Symbol+"&start_date="+param.StartDate+"&end_date="+param.EndDate+"&adjust=hfq")
		by,_:=ioutil.ReadAll(res.Body)
		writer.Write(by)
	})

	http.HandleFunc("/recommend", func(writer http.ResponseWriter, request *http.Request) {

		// 通过json解析器解析参数

		writer.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
		writer.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		writer.Header().Set("content-type", "application/json")             //返回数据格式是json
		list:=service.Filter_expensive_list()
		data,_:=json.Marshal(list)
		writer.Write(data)
	})

	http.HandleFunc("/api/recommend", func(writer http.ResponseWriter, request *http.Request) {
		code:=request.FormValue("code")
		price:=request.FormValue("price")
		// 通过json解析器解析参数

		writer.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
		writer.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		writer.Header().Set("content-type", "application/json")             //返回数据格式是json
		money,_:=strconv.ParseFloat(price,64)
		list:=service.GetMoneyRecommend(code,money)
		data,_:=json.Marshal(list)
		writer.Write(data)
	})
	err:=http.ListenAndServe("127.0.0.1:9090",nil)
	fmt.Println(err)
	for{
		fmt.Println(err)
	}

}


//https://a1.m1907.cn/api/v/?z=f132c044a68b5a642e89f9157518addb&jx=https://m.iqiyi.com/v_26ex87ts5gc.html&s1ig=11401&g=