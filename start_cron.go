package main

import (
	"github.com/robfig/cron"
	"github.com/xiuos/mozi/service"
)

type GenIssueJob struct {
}

func (g GenIssueJob) Run() {
	service.CreateIssueCorn()
}

type CheckAndCreatOddsJob struct {
}

func (g CheckAndCreatOddsJob) Run() {
	service.CreateIssueCorn()
}
func main() {
	service.CreateIssueCorn()
	service.CheckAndCreateOdds()
	c := cron.New()
	c.AddJob("0 0 3 * * *", CheckAndCreatOddsJob{})
	c.AddJob("0 0 4 * * *", GenIssueJob{})
	c.Start()
	select {}
}
