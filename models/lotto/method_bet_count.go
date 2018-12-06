package lotto

import (
	"errors"
	"regexp"
	"strings"
)

func GetCount(me MethodEngine, o *Order) (int, error) {
	return 1, nil
}

func GetCountNotRepeat(me MethodEngine, o *Order) (int, error) {
	if ok, _ := regexp.MatchString(me.GetOpt().BetReg, o.Content); !ok {
		return 0, errors.New("号码格式错误")
	}
	orderList := strings.Split(o.Content, ",")
	oList := SliceRemoveDuplicates(orderList)
	if len(orderList) != len(oList) {
		return 0, errors.New("号码格式错误[1]")
	}
	return 1, nil
}

// 11选5任选
func GetCountSyxwOptionWin(me MethodEngine, o *Order) (int, error) {
	if ok, _ := regexp.MatchString(me.GetOpt().BetReg, o.Content); !ok {
		return 0, errors.New("号码格式错误")
	}
	orderList := strings.Split(o.Content, ",")
	oList := SliceRemoveDuplicates(orderList)
	if len(orderList) != len(oList) {
		return 0, errors.New("号码格式错误[1]")
	}

	return 1, nil
}

// 11选5组选
func GetCountSyxwGroup(me MethodEngine, o *Order) (int, error) {
	if ok, _ := regexp.MatchString(me.GetOpt().BetReg, o.Content); !ok {
		return 0, errors.New("号码格式错误")
	}
	orderList := strings.Split(o.Content, ",")
	oList := SliceRemoveDuplicates(orderList)
	if len(orderList) != len(oList) {
		return 0, errors.New("号码格式错误[1]")
	}

	c := Combination(len(orderList), me.GetOpt().OptionWin[0])

	return c, nil
}

func GetCountF3DCombo(me MethodEngine, o *Order) (int, error) {
	if ok, _ := regexp.MatchString(me.GetOpt().BetReg, o.Content); !ok {
		return 0, errors.New("Number error")
	}
	l := strings.Split(o.Content, "")
	il, _ := S2IList(l)
	for _, i := range il {
		for _, j := range il[1:] {
			if i > j {
				return 0, errors.New("Number error")
			}
		}
	}

	return 1, nil
}
