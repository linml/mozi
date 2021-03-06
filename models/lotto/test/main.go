package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/models/lotto"
	"github.com/xiuos/mozi/service"
)

func bet() {
	uid := 7
	lottoID := 1
	issueInfo, err := lotto.GetCurIssue(lottoID)
	if err != nil {
		fmt.Println(err)
		return
	}

	oo := lotto.Order{
		UserID:     uid,
		LottoID:    lottoID,
		Issue:      issueInfo.Issue,
		MethodCode: "10001",
		PlayCode:   "2",
		BetContent: "2",
		Amount:     decimal.NewFromFloat(7.1),
		IP:         common.InetAton("127.0.0.1"),
	}

	err = service.Bet(&oo)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("下单成功")
	}
}

func main() {
	for i := 0; i < 5; i++ {
		bet()
	}

	o := lotto.Order{BetContent: "1", Odds: decimal.NewFromFloat(9)}
	engine := lotto.PlayMethod[lotto.SscNumber1]
	err := engine.CheckBetLegal(&o)
	if err != nil {
		fmt.Println(err)
		return
	}
	betCount, err := engine.GetBetCount(&o)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("bet count:", betCount)

	w, err := engine.FindWin(&o, []string{"1", "2", "3", "4", "5"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Win: ", w)

}
