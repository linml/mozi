package lotto

import "github.com/shopspring/decimal"

type WinInfo struct {
	Info  string
	Odds  decimal.Decimal
	Count int
}

type MethodEngine interface {
	GetOpt() Opt
	CheckBetLegal(o *Order) error
	GetBetCount(o *Order) (int, error)
	FindWin(o *Order, num []string) (*[]WinInfo, error)
}

type Opt struct {
	Type      int
	BetReg    string
	BetList   []string
	Bit       []int
	OptionWin []int
}

type BaseMethod struct {
	opt               Opt
	checkBetLegalFunc func(me MethodEngine, o *Order) error
	getBetCountFunc   func(me MethodEngine, o *Order) (int, error)
	findWinFunc       func(me MethodEngine, o *Order, num []string) (*[]WinInfo, error)
}

func (bm BaseMethod) GetOpt() Opt {
	return bm.opt
}
func (bm BaseMethod) CheckBetLegal(o *Order) error {
	return bm.checkBetLegalFunc(bm, o)
}

func (bm BaseMethod) GetBetCount(o *Order) (int, error) {
	return bm.getBetCountFunc(bm, o)
}

func (bm BaseMethod) FindWin(o *Order, num []string) (*[]WinInfo, error) {
	return bm.findWinFunc(bm, o, num)
}
