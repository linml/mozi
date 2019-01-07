package main

import (
	"fmt"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/models"
	"github.com/xiuos/mozi/service"
	"github.com/xiuos/mozi/xlog"
	"os"
	"runtime"
	"strings"
)

func init() {

	savePath, err := common.Conf.String("log", "path")
	if err != nil {
		os.Exit(1)
	}
	filePrefix, err := common.Conf.String("log", "lotto_myself_file_prefix")
	if err != nil {
		os.Exit(1)
	}
	common.Logger = xlog.NewLogger(savePath, filePrefix, xlog.Ldate|xlog.Ltime|xlog.Lmicroseconds|xlog.Llongfile, xlog.LevelInfo)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	lottoMyselfSettings, err := models.GetSysSettings(models.SysLottoCalcMyself)
	fmt.Println("Calc lotto:", lottoMyselfSettings.SysValue)
	if err != nil {
		common.Logger.Warn(fmt.Sprintf("%s", err))
	}
	lottoMyselfStrList := strings.Split(lottoMyselfSettings.SysValue, ",")
	lottoMyselfList, _ := common.S2IList(lottoMyselfStrList)
	for _, lid := range lottoMyselfList {
		fmt.Println("Calc lotto:", lid)
		go service.SelfCalcLottoV1(lid)
	}
	select {}
}
