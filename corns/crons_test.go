package corns

import (
	"testing"
	"time"
)

func TestUpdate(t *testing.T) {
	CronWechat()
	time.Sleep(time.Second * 100)
}
