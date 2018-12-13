package lotto

import (
	"errors"
	"regexp"
	"strings"
)

var (
	NumberErr = errors.New("开奖号码不合法")
)

var DrawNumberLegalFuncMap = map[int]func(numbers string) error{
	CQSSC:   CheckDrawNumberSSC,
	XJSSC:   CheckDrawNumberSSC,
	TJSSC:   CheckDrawNumberSSC,
	MZ1F:    CheckDrawNumberSSC,
	MZ2F:    CheckDrawNumberSSC,
	MZ3F:    CheckDrawNumberSSC,
	MZ5F:    CheckDrawNumberSSC,
	SD11X5:  CheckDrawNumber11x5,
	JX11X5:  CheckDrawNumber11x5,
	GD11X5:  CheckDrawNumber11x5,
	JS11X5:  CheckDrawNumber11x5,
	AH11X5:  CheckDrawNumber11x5,
	SX11X5:  CheckDrawNumber11x5,
	SH11X5:  CheckDrawNumber11x5,
	MZ11X5:  CheckDrawNumber11x5,
	JSK3:    CheckDrawNumberKuai3,
	AHK3:    CheckDrawNumberKuai3,
	HBK3:    CheckDrawNumberKuai3,
	HNK3:    CheckDrawNumberKuai3,
	JSTB:    CheckDrawNumberKuai3,
	MZK3:    CheckDrawNumberKuai3,
	BJPK10:  CheckDrawNumberPK10,
	XYFT:    CheckDrawNumberPK10,
	MZPK10:  CheckDrawNumberPK10,
	FC3D:    CheckDrawNumber3D,
	PL3:     CheckDrawNumberPL3,
	PL5:     CheckDrawNumberPL5,
	GDKLSF:  CheckDrawNumberNongChang,
	CQKLSF:  CheckDrawNumberNongChang,
	TJKLSF:  CheckDrawNumberNongChang,
	MZLHC:   CheckDrawNumberMarkSix,
	MZ10LHC: CheckDrawNumberMarkSix,
}

/* 时时彩 */
func CheckDrawNumberSSC(numbers string) error {
	if m, _ := regexp.MatchString("^([0-9],){4}[0-9]$", numbers); !m {
		return NumberErr
	}
	return nil
}

/* 3D */
func CheckDrawNumber3D(numbers string) error {
	if m, _ := regexp.MatchString("^([0-9],){2}[0-9]$", numbers); !m {
		return NumberErr
	}
	return nil
}

/* 排列3 */
func CheckDrawNumberPL3(numbers string) error {
	if m, _ := regexp.MatchString("^([0-9],){2}[0-9]$", numbers); !m {
		return NumberErr
	}
	return nil
}

/* 排列5 */
func CheckDrawNumberPL5(numbers string) error {
	if m, _ := regexp.MatchString("^([0-9],){4}[0-9]$", numbers); !m {
		return NumberErr
	}
	return nil
}

/* PK10 */
func CheckDrawNumberPK10(numbers string) error {
	_numbers := strings.Split(numbers, ",")

	for _i, i := range _numbers {
		for _, j := range _numbers[_i+1:] {
			if i == j {
				return NumberErr
			}
		}
	}

	if m, _ := regexp.MatchString("^((0[0-9]|[0-9]|10),){9}(0[0-9]|[0-9]|10)$", numbers); !m {
		return NumberErr
	}

	m := map[string]string{}
	l := strings.Split(numbers, ",")
	for _, v := range l {
		m[v] = v
	}

	if len(m) != 10 {
		return NumberErr
	}

	return nil
}

/* MarkSix */
func CheckDrawNumberMarkSix(numbers string) error {
	_numbers := strings.Split(numbers, ",")

	for _i, i := range _numbers {
		for _, j := range _numbers[_i+1:] {
			if i == j {
				return NumberErr
			}
		}
	}

	if m, _ := regexp.MatchString("^(([1-9]|1[0-9]|2[0-9]|3[0-9]|4[0-9]),){6}([1-9]|1[0-9]|2[0-9]|3[0-9]|4[0-9])$", numbers); !m {
		return NumberErr
	}

	m := map[string]string{}
	l := strings.Split(numbers, ",")
	for _, v := range l {
		m[v] = v
	}

	if len(m) != 7 {
		return NumberErr
	}
	return nil
}

/* Lucky28 */
func CheckDrawNumberLucky28(numbers string) error {
	if m, _ := regexp.MatchString("^([0-9],){2}[0-9]$", numbers); !m {
		return NumberErr
	}
	return nil
}

/* 快3 */
func CheckDrawNumberKuai3(numbers string) error {
	if m, _ := regexp.MatchString("^([1-6],){2}[1-6]$", numbers); !m {
		return NumberErr
	}
	return nil
}

/* 农场 */
func CheckDrawNumberNongChang(numbers string) error {
	_numbers := strings.Split(numbers, ",")

	for _i, i := range _numbers {
		for _, j := range _numbers[_i+1:] {
			if i == j {
				return NumberErr
			}
		}
	}

	if m, _ := regexp.MatchString("^(([1-9]|0[1-9]|1[0-9]|20),){7}([1-9]|0[1-9]|1[0-9]|20)$", numbers); !m {
		return NumberErr
	}
	m := map[string]string{}
	l := strings.Split(numbers, ",")
	for _, v := range l {
		m[v] = v
	}

	if len(m) != 8 {
		return NumberErr
	}
	return nil
}

/* 11x5 */
func CheckDrawNumber11x5(numbers string) error {
	_numbers := strings.Split(numbers, ",")

	for _i, i := range _numbers {
		for _, j := range _numbers[_i+1:] {
			if i == j {
				return NumberErr
			}
		}
	}

	if m, _ := regexp.MatchString("^((0[1-9]|1[0-1]),){4}(0[1-9]|1[0-1])$", numbers); !m {
		return NumberErr
	}

	m := map[string]string{}
	l := strings.Split(numbers, ",")
	for _, v := range l {
		m[v] = v
	}

	if len(m) != 5 {
		return NumberErr
	}
	return nil
}

func CheckDrawNumberLegal(lottoID int, numbers string) error {
	if f, ok := DrawNumberLegalFuncMap[lottoID]; ok {
		err := f(numbers)
		return err
	}
	return errors.New("找不到对应彩种")
}
