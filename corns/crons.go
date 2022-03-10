package corns

import (
	"context"
	"fmt"
	"github.com/robfig/cron/v3"
)

var corns *cron.Cron
var ctx = context.Background()

func RegisterCrons() {
	if corns == nil {
		corns = cron.New()
	}
	fmt.Println("启动定时任务")
	corns.AddFunc("*/1 * * * *", SelectData)
	corns.AddFunc("*/1 * * * *", CronStrategy)

	corns.Start()

}
