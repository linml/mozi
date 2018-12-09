package lotto

import (
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"sort"
	"strconv"
	"strings"
)

var (
	SscNumberPoolJUNKO = []string{"012", "123", "234", "345", "456", "567", "678", "789", "089", "019"} // 时时彩 顺子

)

func FindWinDig(me MethodEngine, o *Order, drawNumber []string) (*[]WinInfo, error) {
	n := drawNumber[me.GetOpt().Bit[0]]
	var w []WinInfo
	if o.BetContent == n {
		w = append(w, WinInfo{Info: o.BetContent, Odds: o.Odds, Count: 1})
		return &w, nil
	}
	return &[]WinInfo{}, nil
}

// 时时彩 大,小,单,双
func FindWinSscTwoSide(me MethodEngine, o *Order, drawNumber []string) (*[]WinInfo, error) {
	var w []WinInfo
	num, _ := strconv.Atoi(drawNumber[me.GetOpt().Bit[0]])

	switch o.BetContent {
	case Big:
		if num > 4 {
			w = append(w, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	case Small:
		if num < 5 {
			w = append(w, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	case Odd:
		if num%2 != 0 {
			w = append(w, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	case Even:
		if num%2 == 0 {
			w = append(w, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	default:
		return &w, errors.New("投注内容异常")
	}

	return &w, nil
}

// 时时彩 总和 大,小,单,双,龙,虎,和
func FindWinSscTwoSideSum(me MethodEngine, o *Order, drawNumber []string) (*[]WinInfo, error) {
	var w []WinInfo
	drawNumberIntList, _ := S2IList(drawNumber)
	drawNumberSum := SumInt(drawNumberIntList)
	switch o.BetContent {
	case Big:
		if drawNumberSum > 22 {
			w = append(w, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	case Small:
		if drawNumberSum <= 22 {
			w = append(w, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	case Odd:
		if drawNumberSum%2 != 0 {
			w = append(w, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	case Even:
		if drawNumberSum%2 == 0 {
			w = append(w, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	case Dragon:
		if drawNumberIntList[0] > drawNumberIntList[4] {
			w = append(w, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	case Tiger:
		if drawNumberIntList[0] < drawNumberIntList[4] {
			w = append(w, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	case He:
		if drawNumberIntList[0] == drawNumberIntList[4] {
			w = append(w, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	default:
		return &w, errors.New("投注内容异常")
	}

	return &w, nil
}

// 时时彩 豹子 > 顺子 > 对子 > 半顺 > 杂六
func FindWinSscExtra(me MethodEngine, o *Order, drawNumber []string) (*[]WinInfo, error) {
	var w []WinInfo
	drawNumberIntList, _ := S2IList(drawNumber)
	num := []int{}
	for _, v := range me.GetOpt().Bit {
		num = append(num, drawNumberIntList[v])
	}
	switch o.BetContent {
	case BaoZi:
		if num[0] == num[1] && num[2] == num[1] {
			w = append(w, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	case ShunZi:
		sort.Ints(num)
		s := fmt.Sprintf("%d%d%d", num[0], num[1], num[2])
		for k, _ := range SscNumberPoolJUNKO {
			if SscNumberPoolJUNKO[k] == s {
				w = append(w, WinInfo{
					Info:  o.BetContent,
					Odds:  o.Odds,
					Count: 1,
				})
			}
		}
	case DuiZi:
		if (num[0] == num[1] && num[0] != num[2]) || (num[0] == num[2] && num[0] != num[1]) || (num[1] == num[2] && num[1] != num[0]) {
			w = append(w, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	case BanShun:
		for k, _ := range num {
			_v := (num[k] + 1) % 10
			for j, _ := range num {
				if _v == num[j] {
					w = append(w, WinInfo{
						Info:  o.BetContent,
						Odds:  o.Odds,
						Count: 1,
					})
				}
			}
		}
	case ZaLiu:
		flag := 0
		if num[0] == num[1] && num[2] == num[1] {
			flag = 1
		}
		if (num[0] == num[1] && num[0] != num[2]) || (num[0] == num[2] && num[0] != num[1]) || (num[1] == num[2] && num[1] != num[0]) {
			flag = 1
		}
		for k, _ := range num {
			_v := (num[k] + 1) % 10
			for j, _ := range num {
				if _v == num[j] {
					flag = 1
				}
			}
		}
		sort.Ints(num)
		s := fmt.Sprintf("%d%d%d", num[0], num[1], num[2])
		for k, _ := range SscNumberPoolJUNKO {
			if SscNumberPoolJUNKO[k] == s {
				flag = 1
			}
		}

		if flag == 0 {
			w = append(w, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}

	default:
		return &w, errors.New("投注内容异常")
	}

	return &w, nil

}

// 11选5 大,小,单,双 (开11退本金)
func FindWinSyxwTwoSide(me MethodEngine, o *Order, drawNumber []string) (*[]WinInfo, error) {
	var w []WinInfo
	num, _ := strconv.Atoi(drawNumber[me.GetOpt().Bit[0]])

	if num == 11 {
		w = append(w, WinInfo{
			Info:  o.BetContent,
			Odds:  decimal.NewFromFloat(1),
			Count: 1,
		})
	} else {
		switch o.BetContent {
		case Big:
			if num > 5 {
				w = append(w, WinInfo{
					Info:  o.BetContent,
					Odds:  o.Odds,
					Count: 1,
				})
			}
		case Small:
			if num < 6 {
				w = append(w, WinInfo{
					Info:  o.BetContent,
					Odds:  o.Odds,
					Count: 1,
				})
			}
		case Odd:
			if num%2 != 0 {
				w = append(w, WinInfo{
					Info:  o.BetContent,
					Odds:  o.Odds,
					Count: 1,
				})
			}
		case Even:
			if num%2 == 0 {
				w = append(w, WinInfo{
					Info:  o.BetContent,
					Odds:  o.Odds,
					Count: 1,
				})
			}
		default:
			return &w, errors.New("投注内容异常")
		}

	}

	return &w, nil
}

// 11选5 总和大,小,单,双,尾大,尾小,龙,虎 (买大小，如果开30不计输赢)
func FindWinSyxwTwoSideSum(me MethodEngine, o *Order, drawNumber []string) (*[]WinInfo, error) {
	var w []WinInfo
	drawNumberIntList, _ := S2IList(drawNumber)
	num := SumInt(drawNumberIntList)

	switch o.BetContent {
	case Big:
		if num == 30 {
			w = append(w, WinInfo{
				Info:  o.BetContent,
				Odds:  decimal.NewFromFloat(1),
				Count: 1,
			})
		} else {
			if num > 30 {
				w = append(w, WinInfo{
					Info:  o.BetContent,
					Odds:  o.Odds,
					Count: 1,
				})
			}
		}
	case Small:
		if num == 30 {
			w = append(w, WinInfo{
				Info:  o.BetContent,
				Odds:  decimal.NewFromFloat(1),
				Count: 1,
			})
		} else {
			if num < 30 {
				w = append(w, WinInfo{
					Info:  o.BetContent,
					Odds:  o.Odds,
					Count: 1,
				})
			}
		}
	case Odd:
		if num%2 != 0 {
			w = append(w, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	case Even:
		if num%2 == 0 {
			w = append(w, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	case WeiDa:
		if num%10 > 4 {
			w = append(w, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	case WeiXiao:
		if num%10 < 5 {
			w = append(w, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	case Dragon:
		if drawNumberIntList[0] > drawNumberIntList[4] {
			w = append(w, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	case Tiger:
		if drawNumberIntList[0] < drawNumberIntList[4] {
			w = append(w, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	default:
		return &w, errors.New("投注内容异常")
	}

	return &w, nil
}

// 11选5 任选N中Y
func FindWinSyxwOptionWin(me MethodEngine, o *Order, drawNumber []string) (*[]WinInfo, error) {
	var w []WinInfo
	betList := strings.Split(o.BetContent, ",")

	//_N := mi.Option.OptionWin[0]
	_Y := me.GetOpt().OptionWin[1]

	y := 0
	for _, i := range drawNumber {
		for _, j := range betList {
			if i == j {
				y += 1
			}
		}
	}

	if y == _Y {
		w = append(w, WinInfo{
			Info:  o.BetContent,
			Odds:  o.Odds,
			Count: 1,
		})
	}

	return &w, nil
}

// 11选5 组选
func FindWinSyxwGroupWin(me MethodEngine, o *Order, drawNumber []string) (*[]WinInfo, error) {
	var w []WinInfo
	betList := strings.Split(o.BetContent, ",")

	num := []string{}
	for _, v := range me.GetOpt().Bit {
		num = append(num, drawNumber[v])
	}

	//_N := mi.Option.OptionWin[0]
	_Y := me.GetOpt().OptionWin[0]

	y := 0
	for _, i := range num {
		for _, j := range betList {
			if i == j {
				y += 1
			}
		}
	}

	if y >= _Y {
		w = append(w, WinInfo{
			Info:  o.BetContent,
			Odds:  o.Odds,
			Count: 1,
		})
	}

	return &w, nil
}

// 11选5 直选
func FindWinSyxwZhixuan(me MethodEngine, o *Order, drawNumber []string) (*[]WinInfo, error) {
	var w []WinInfo
	betList := strings.Split(o.BetContent, ",")

	num := []string{}
	for _, v := range me.GetOpt().Bit {
		num = append(num, drawNumber[v])
	}

	flag := true
	for _, i := range num {
		for _, j := range betList {
			if i != j {
				flag = flag
			}
		}
	}

	if flag {
		w = append(w, WinInfo{
			Info:  o.BetContent,
			Odds:  o.Odds,
			Count: 1,
		})
	}

	return &w, nil
}

// 快3 总和点数
func FindWinKuai3Sum(me MethodEngine, o *Order, drawNumber []string) (*[]WinInfo, error) {
	var winInfoList []WinInfo
	drawNumberIntList, _ := S2IList(drawNumber)
	num := SumInt(drawNumberIntList)
	betNum, err := strconv.Atoi(o.BetContent)
	if err == nil {
		if num == betNum {
			winInfoList = append(winInfoList, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	}

	return &winInfoList, nil
}

// 快3 三军
func FindWinKuai33Jun(me MethodEngine, o *Order, drawNumber []string) (*[]WinInfo, error) {
	var winInfoList []WinInfo
	for _, v := range drawNumber {
		if v == o.BetContent {
			winInfoList = append(winInfoList, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
			break
		}
	}
	return &winInfoList, nil
}

// 快3 大小
func FindWinKuai3TwoSide(me MethodEngine, o *Order, drawNumber []string) (*[]WinInfo, error) {
	var winInfoList []WinInfo
	drawNumberIntList, _ := S2IList(drawNumber)
	num := SumInt(drawNumberIntList)

	if drawNumber[0] == drawNumber[1] && drawNumber[1] == drawNumber[2] {
		return &winInfoList, nil
	}

	switch o.BetContent {
	case Big:
		if num >= 11 && num <= 17 {
			winInfoList = append(winInfoList, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	case Small:
		if num >= 4 && num <= 10 {
			winInfoList = append(winInfoList, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	}

	return &winInfoList, nil
}

// 快3 围骰
func FindWinKuai3Weitou(me MethodEngine, o *Order, drawNumber []string) (*[]WinInfo, error) {
	var winInfoList []WinInfo
	if o.BetContent == strings.Join(drawNumber, "") {
		winInfoList = append(winInfoList, WinInfo{
			Info:  o.BetContent,
			Odds:  o.Odds,
			Count: 1,
		})
	}

	return &winInfoList, nil
}

// 快3 全骰
func FindWinKuai3Quantou(me MethodEngine, o *Order, drawNumber []string) (*[]WinInfo, error) {
	var winInfoList []WinInfo
	if drawNumber[0] == drawNumber[1] && drawNumber[1] == drawNumber[2] {
		winInfoList = append(winInfoList, WinInfo{
			Info:  o.BetContent,
			Odds:  o.Odds,
			Count: 1,
		})
	}

	return &winInfoList, nil
}

// 快3 长牌
func FindWinKuai3Chang(me MethodEngine, o *Order, drawNumber []string) (*[]WinInfo, error) {
	var winInfoList []WinInfo

	if CheckInStr(drawNumber, o.BetContent[0:1]) && CheckInStr(drawNumber, o.BetContent[1:2]) {
		winInfoList = append(winInfoList, WinInfo{
			Info:  o.BetContent,
			Odds:  o.Odds,
			Count: 1,
		})
	}

	return &winInfoList, nil
}

// 快3 短牌
func FindWinKuai3Duan(me MethodEngine, o *Order, drawNumber []string) (*[]WinInfo, error) {
	var winInfoList []WinInfo

	if strings.Count(strings.Join(drawNumber, ""), o.BetContent) > 0 {
		winInfoList = append(winInfoList, WinInfo{
			Info:  o.BetContent,
			Odds:  o.Odds,
			Count: 1,
		})
	}

	return &winInfoList, nil
}

// pk10 大小单双龙虎
func FindWinPK10TwoSide(me MethodEngine, o *Order, drawNumber []string) (*[]WinInfo, error) {
	var winInfoList []WinInfo
	num, _ := strconv.Atoi(drawNumber[me.GetOpt().Bit[0]])
	num1, _ := strconv.Atoi(drawNumber[9-me.GetOpt().Bit[0]])

	switch o.BetContent {
	case Big:
		if num > 5 {
			winInfoList = append(winInfoList, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	case Small:
		if num < 6 {
			winInfoList = append(winInfoList, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	case Odd:
		if num%2 != 0 {
			winInfoList = append(winInfoList, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	case Even:
		if num%2 == 0 {
			winInfoList = append(winInfoList, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	case Dragon:
		if num > num1 {
			winInfoList = append(winInfoList, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	case Tiger:
		if num < num1 {
			winInfoList = append(winInfoList, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	default:
		return &winInfoList, errors.New("投注内容异常")
	}

	return &winInfoList, nil
}

// pk10 冠亚大小单双
func FindWinPK10GYTwoSide(me MethodEngine, o *Order, drawNumber []string) (*[]WinInfo, error) {
	var winInfoList []WinInfo
	num, _ := strconv.Atoi(drawNumber[me.GetOpt().Bit[0]])
	num1, _ := strconv.Atoi(drawNumber[9-me.GetOpt().Bit[1]])
	num = num + num1

	switch o.BetContent {
	case Big:
		if num > 11 {
			winInfoList = append(winInfoList, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	case Small:
		if num < 12 {
			winInfoList = append(winInfoList, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	case Odd:
		if num%2 != 0 {
			winInfoList = append(winInfoList, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	case Even:
		if num%2 == 0 {
			winInfoList = append(winInfoList, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	default:
		return &winInfoList, errors.New("投注内容异常")
	}

	return &winInfoList, nil
}

// pk10 冠亚
func FindWinPK10GY(me MethodEngine, o *Order, drawNumber []string) (*[]WinInfo, error) {
	var winInfoList []WinInfo
	num, _ := strconv.Atoi(drawNumber[me.GetOpt().Bit[0]])
	num1, _ := strconv.Atoi(drawNumber[me.GetOpt().Bit[1]])
	betNum, _ := strconv.Atoi(o.BetContent)
	num = num + num1
	if num == betNum {
		winInfoList = append(winInfoList, WinInfo{
			Info:  o.BetContent,
			Odds:  o.Odds,
			Count: 1,
		})
	}
	return &winInfoList, nil
}

// pk10 组合
func FindWinF3DCombo(me MethodEngine, o *Order, drawNumber []string) (*[]WinInfo, error) {
	var winInfoList []WinInfo
	betList := strings.Split(o.BetContent, "")

	num := []string{}
	for _, v := range me.GetOpt().Bit {
		num = append(num, drawNumber[v])
	}

	//_N := mi.Option.OptionWin[0]
	_Y := me.GetOpt().OptionWin[0]

	y := 0
	for _, i := range betList {
		for jj, j := range num {
			if i == j {
				y += 1
				num[jj] = "pass"
				break
			}
		}
	}

	if y >= _Y {
		winInfoList = append(winInfoList, WinInfo{
			Info:  o.BetContent,
			Odds:  o.Odds,
			Count: 1,
		})
	}

	return &winInfoList, nil
}

// 3d
func FindWinF3DNumberMultiple(me MethodEngine, o *Order, drawNumber []string) (*[]WinInfo, error) {
	var winInfoList []WinInfo
	num := []string{}
	for _, v := range me.GetOpt().Bit {
		num = append(num, drawNumber[v])
	}

	betNum := strings.Split(o.BetContent, "")

	flag := true

	for _, i := range num {
		for _, j := range betNum {
			if i != j {
				flag = false
			}
		}
	}

	if flag {
		winInfoList = append(winInfoList, WinInfo{
			Info:  o.BetContent,
			Odds:  o.Odds,
			Count: 1,
		})
	}
	return &winInfoList, nil

}

// 3d
func FindWinF3DSum(me MethodEngine, o *Order, drawNumber []string) (*[]WinInfo, error) {
	var winInfoList []WinInfo
	numList := []int{}
	drawNumberIntList, _ := S2IList(drawNumber)

	for _, v := range me.GetOpt().Bit {
		numList = append(numList, drawNumberIntList[v])
	}

	num := SumInt(numList)

	betNum, err := strconv.Atoi(o.BetContent)

	if err == nil {
		if betNum == num {
			winInfoList = append(winInfoList, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	}
	return &winInfoList, nil

}

//3d 尾数
func FindWinF3DSumWei(me MethodEngine, o *Order, drawNumber []string) (*[]WinInfo, error) {
	var winInfoList []WinInfo
	numList := []int{}
	drawNumberIntList, _ := S2IList(drawNumber)

	for _, v := range me.GetOpt().Bit {
		numList = append(numList, drawNumberIntList[v])
	}

	num := SumInt(numList) % 10

	betNum, err := strconv.Atoi(o.BetContent)

	if err == nil {
		if betNum == num {
			winInfoList = append(winInfoList, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	}
	return &winInfoList, nil

}

// 3d 跨度
func FindWinF3DKuaidu(me MethodEngine, o *Order, drawNumber []string) (*[]WinInfo, error) {
	var winInfoList []WinInfo
	numList := []int{}
	drawNumberIntList, _ := S2IList(drawNumber)

	for _, v := range me.GetOpt().Bit {
		numList = append(numList, drawNumberIntList[v])
	}

	num := MaxVal(numList) - MinVal(numList)

	betNum, err := strconv.Atoi(o.BetContent)

	if err == nil {
		if betNum == num {
			winInfoList = append(winInfoList, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	}
	return &winInfoList, nil

}

// 3d
func FindWinF3DTwoSide(me MethodEngine, o *Order, drawNumber []string) (*[]WinInfo, error) {
	var winInfoList []WinInfo
	drawNumberIntList, _ := S2IList(drawNumber)
	num := drawNumberIntList[me.GetOpt().Bit[0]]

	switch o.BetContent {
	case Big:
		if num > 4 {
			winInfoList = append(winInfoList, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	case Small:
		if num < 5 {
			winInfoList = append(winInfoList, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	case Odd:
		if num%2 != 0 {
			winInfoList = append(winInfoList, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	case Even:
		if num%2 == 0 {
			winInfoList = append(winInfoList, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	case ZhiShu:
		if CheckInInt([]int{1, 2, 3, 5, 7}, num) {
			winInfoList = append(winInfoList, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	case HeShu:
		if CheckInInt([]int{0, 4, 6, 8, 9}, num) {
			winInfoList = append(winInfoList, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	default:
		return &winInfoList, errors.New("投注内容异常")
	}

	return &winInfoList, nil

}

func FindWinF3DTwoSide1(me MethodEngine, o *Order, drawNumber []string) (*[]WinInfo, error) {
	var winInfoList []WinInfo
	drawNumberIntList, _ := S2IList(drawNumber)
	numList := []int{}
	for _, v := range me.GetOpt().Bit {
		numList = append(numList, drawNumberIntList[v])
	}
	num := SumInt(numList)

	switch o.BetContent {
	case Odd:
		if num%2 != 0 {
			winInfoList = append(winInfoList, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	case Even:
		if num%2 == 0 {
			winInfoList = append(winInfoList, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	default:
		return &winInfoList, errors.New("投注内容异常")
	}

	return &winInfoList, nil

}

func FindWinF3DTwoSide2(me MethodEngine, o *Order, drawNumber []string) (*[]WinInfo, error) {
	var winInfoList []WinInfo
	drawNumberIntList, _ := S2IList(drawNumber)
	numList := []int{}
	for _, v := range me.GetOpt().Bit {
		numList = append(numList, drawNumberIntList[v])
	}
	num := SumInt(numList)

	switch o.BetContent {
	case Big:
		if num > 13 {
			winInfoList = append(winInfoList, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	case Small:
		if num < 14 {
			winInfoList = append(winInfoList, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	case Odd:
		if num%2 != 0 {
			winInfoList = append(winInfoList, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	case Even:
		if num%2 == 0 {
			winInfoList = append(winInfoList, WinInfo{
				Info:  o.BetContent,
				Odds:  o.Odds,
				Count: 1,
			})
		}
	default:
		return &winInfoList, errors.New("投注内容异常")
	}

	return &winInfoList, nil

}
