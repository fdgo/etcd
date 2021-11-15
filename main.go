package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

type CronJob struct {
	expr     *cronexpr.Expression
	nextTime time.Time
}

func main() {
	var (
		cronJob       *CronJob
		expr          *cronexpr.Expression
		now           time.Time
		scheduleTable map[string]*CronJob
	)
	scheduleTable = make(map[string]*CronJob)

	expr = cronexpr.MustParse("*/5 * * * * * *")
	cronJob = &CronJob{
		expr:     expr,
		nextTime: expr.Next(now),
	}
	scheduleTable["job1"] = cronJob

	expr = cronexpr.MustParse("*/6 * * * * * *")
	cronJob = &CronJob{
		expr:     expr,
		nextTime: expr.Next(now),
	}
	scheduleTable["job2"] = cronJob

	go func() {
		var (
			jobName string
			cronJob *CronJob
			now     time.Time
		)
		for {
			now = time.Now()
			for jobName, cronJob = range scheduleTable {
				if cronJob.nextTime.Before(now) || cronJob.nextTime.Equal(now) {
					go func(jobName string) {
						fmt.Println("doing:", jobName)
					}(jobName)

					cronJob.nextTime = cronJob.expr.Next(now)
					fmt.Println(jobName, "next to do time:", cronJob.nextTime)
				}
			}
			select {
			case <- time.NewTimer(100* time.Millisecond).C:}

		}
	}()

	time.Sleep(time.Second*1000)
}
