package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/xiuos/lotto"
)

func main() {
	o := lotto.Order{Content: "1", Odds: decimal.NewFromFloat(9)}
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
