package models

import "github.com/xiuos/mozi/models/lotto"

func MethodCheckBetLegal(method string, o *lotto.Order) error {
	engine, err := lotto.GetMethodEngine(method)
	if err != nil {
		return err
	}
	err = engine.CheckBetLegal(o)
	return err
}

func MethodBetCount(method string, o *lotto.Order) (int, error) {
	engine, err := lotto.GetMethodEngine(method)
	if err != nil {
		return 0, err
	}
	return engine.GetBetCount(o)
}

func MethodFindWin(method string, o *lotto.Order, drawNumber []string) (*[]lotto.WinInfo, error) {
	engine, err := lotto.GetMethodEngine(method)
	if err != nil {
		return nil, err
	}
	return engine.FindWin(o, drawNumber)
}
