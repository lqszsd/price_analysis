package service

import (
	"fmt"
	"testing"
)

func TestDetailTest(t *testing.T){
	var d Detail
	d=new(DetailResponse)
	//招商白酒
	d.Request("https://fundgz.1234567.com.cn/js/161725.js?rt=ZSZZBJZSLOFA")
	fmt.Println(d)
}
