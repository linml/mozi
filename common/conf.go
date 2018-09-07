package common

import (
	"fmt"
	"github.com/robfig/config"
	"os"
)

var Conf *config.Config

func init() {
	var err error
	Conf, err = config.ReadDefault("conf.ini")
	if err != nil {
		fmt.Println("Config file not found")
		os.Exit(1)
	}
}
