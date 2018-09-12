package main

import (
	"github.com/robfig/cron"
	"mozi/service/game"
	"github.com/xiuos/mozi/service"
)

type GenIssueJob struct {
}

func (g GenIssueJob) Run() {
	service.CreateIssueCorn()
}

func main() {
	game.CreateIssueCorn()
	c := cron.New()
	c.AddJob("0 0 4 * * *", GenIssueJob{})
	c.Start()
	select {}
}
