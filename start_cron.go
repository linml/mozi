package main

import (
	"fmt"
	"github.com/robfig/cron"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/service"
)

type GenIssueJob struct {
}

func (g GenIssueJob) Run() {
	service.CreateIssueCorn()
}

type CheckAndCreateOddsJob struct {
}

func (g CheckAndCreateOddsJob) Run() {
	service.CreateIssueCorn()
}

type LottoDayCountJob struct {
}

func (g LottoDayCountJob) Run() {
	fmt.Println(common.GetTimeNowString(), ": start LottoDayCountJob")
	service.RefreshLottoDayCount(common.GetDateNowString())
	fmt.Println(common.GetTimeNowString(), ":  end LottoDayCountJob")

}
func main() {
	service.CreateIssueCorn()
	service.CheckAndCreateOdds()
	c := cron.New()
	c.AddJob("0 0 3 * * *", CheckAndCreateOddsJob{})
	c.AddJob("0 0 4 * * *", GenIssueJob{})
	c.AddJob("0 */1 * * * *", LottoDayCountJob{})
	c.Start()
	select {}
}
