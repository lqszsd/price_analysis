package service

import (
	"fmt"
	"testing"
)

func TestGetMoneyList(t *testing.T) {
	fmt.Println(GetMoneyList())
}

func TestGetMoneyRecommend(t *testing.T) {
	fmt.Println(GetMoneyRecommend("002060", 5.30))
}

func TestLast(t *testing.T) {
	//fmt.Println(filter_expensive_list())
}
