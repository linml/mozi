package lotto

import (
	"errors"
	"regexp"
	"strings"
)

func CheckBetLegalReg(me MethodEngine, o *Order) error {
	if ok, _ := regexp.MatchString(me.GetOpt().BetReg, o.BetContent); !ok {
		return errors.New("号码格式错误")
	}
	return nil
}

func CheckBetLegalList(me MethodEngine, o *Order) error {

	if ok := CheckInStr(me.GetOpt().BetList, o.BetContent); !ok {
		return errors.New("号码格式错误")
	}
	return nil
}

func CheckBetRegNotRepeat(me MethodEngine, o *Order) error {
	if ok, _ := regexp.MatchString(me.GetOpt().BetReg, o.BetContent); !ok {
		return errors.New("号码格式错误")
	}
	orderList := strings.Split(o.BetContent, ",")
	oList := SliceRemoveDuplicates(orderList)
	if len(orderList) != len(oList) {
		return errors.New("号码格式错误[1]")
	}

	return nil
}

func CheckBetRegF3DCombo(me MethodEngine, o *Order) error {
	if ok, _ := regexp.MatchString(me.GetOpt().BetReg, o.BetContent); !ok {
		return errors.New("Number error")
	}
	l := strings.Split(o.BetContent, "")
	il, _ := S2IList(l)
	for _, i := range il {
		for _, j := range il[1:] {
			if i > j {
				return errors.New("Number error")
			}
		}
	}

	return nil
}
