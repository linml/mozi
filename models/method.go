package models

import "github.com/xiuos/lotto"

func MethodCheckBetLegal(method string, o *Order) error {
	engine, err := lotto.GetMethodEngine(method)
	if err != nil {
		return err
	}
	err = engine.CheckBetLegal(o)
	return err
}

func MethodBetCount(method string, o *Order) (int, error) {
	engine, err := lotto.GetMethodEngine(method)
	if err != nil {
		return 0, err
	}
	return engine.GetBetCount(o)
}


func MethodFindWin(method string, o *Order, drawNumber []string) (*[]lotto.WinInfo, error) {
	engine, err := lotto.GetMethodEngine(method)
	if err != nil {
		return nil, err
	}
	return engine.FindWin(o, drawNumber)
}