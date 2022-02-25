package request

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func HttpPost(url string,data string)([]byte){
	client:=http.Client{}
	req,err:=http.NewRequest("POST",url,strings.NewReader(data))
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