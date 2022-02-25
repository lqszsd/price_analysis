package request

import (
	"io/ioutil"
	"net/http"
)

func HttpGet(url string)([]byte){
	client:=http.Client{}
	req,err:=http.NewRequest("GET",url,nil)
	if err != nil {
		// handle error
	}
	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	return body
}