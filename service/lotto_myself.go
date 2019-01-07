package service

import (
	"errors"
	"fmt"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/models/lotto"
	"math/rand"
	"strings"
	"time"
)

func Shuffle(vals []string) []string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ret := make([]string, len(vals))
	perm := r.Perm(len(vals))
	for i, randIndex := range perm {
		ret[i] = vals[randIndex]
	}
	return ret
}

func GenerateRangeNum(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(max-min) + min
	return randNum
}

func GenLotto(ballList []int, count int, resultType int) ([]int, error) {
	// 此函数用于生成连续期号的开奖号码
	// resultType=>0:结果可以重复, 1: 结果不能重复
	var ret []int
	salt := time.Now().UnixNano()

	minBall := ballList[0]
	maxBall := ballList[len(ballList)-1]

	if maxBall < minBall {
		return []int{}, errors.New("号码范围错误")
	}
	if (maxBall-minBall) < count-1 && resultType == 1 {
		return []int{}, errors.New("无法得到不重复的结果")
	}
	for {
		salt = salt << 1
		salt = salt + time.Now().UnixNano()
		salt = salt + 1
		seed := rand.NewSource(salt)
		rng := rand.New(seed)
		v := rng.Int()%maxBall + minBall
		if 1 == resultType {
			flag := false
			for _, j := range ret {
				if j == v {
					flag = true
					break
				}
			}
			if flag == true {
				continue
			}
		}
		ret = append(ret, v)
		if len(ret) == count {
			break
		}
	}
	return ret, nil
}

func formatNumbers(lottoType int, numbers []int) []string {
	retNumbers := []string{}
	formatStr := "%d"
	switch lottoType {
	case lotto.LottoTypeSSC:
		formatStr = "%d"
	case lotto.LottoType11X5:
		formatStr = "%02d"
	case lotto.LottoTypeK3:
		formatStr = "%d"
	case lotto.LottoTypePK10:
		formatStr = "%02d"
	case lotto.LottoTypePL3:
		formatStr = "%d"
	case lotto.LottoTypePL5:
		formatStr = "%d"
	case lotto.LottoTypeKLSF:
		formatStr = "%02d"
	case lotto.LottoTypeKENO:
		formatStr = "%02d"
	case lotto.LottoTypeLHC:
		formatStr = "%02d"
	default:
		formatStr = "%d"
	}
	for i := range numbers {
		retNumbers = append(retNumbers, fmt.Sprintf(formatStr, numbers[i]))
	}
	return retNumbers
}

func DoPreCalc(lottoID int, issue string) ([]int, error) {
	// todo 自开彩票校验
	lottoInfo, err := lotto.GetLotto(lottoID)
	var lottoBall []int
	ballSSC := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	ball11X5 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	ballK3 := []int{1, 2, 3, 4, 5, 6}
	ballPK10 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ballPL3 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	ballPL5 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	ballKL10 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	ballKENO := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80}
	ballLHC := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49}

	orderList, err := lotto.FindLottoOrderList(map[string]string{
		"lotto_id": fmt.Sprintf("%d", lottoID),
		"issue":    issue})
	fmt.Println("orderList:", len(*orderList))

	for {
		switch lottoInfo.LottoType {
		case lotto.LottoTypeSSC:
			lottoBall, err = GenLotto(ballSSC, 5, 0)
		case lotto.LottoType11X5:
			lottoBall, err = GenLotto(ball11X5, 5, 1)
		case lotto.LottoTypeK3:
			lottoBall, err = GenLotto(ballK3, 3, 0)
		case lotto.LottoTypePK10:
			lottoBall, err = GenLotto(ballPK10, 10, 1)
		case lotto.LottoTypePL3:
			lottoBall, err = GenLotto(ballPL3, 3, 0)
		case lotto.LottoTypePL5:
			lottoBall, err = GenLotto(ballPL5, 5, 0)
		case lotto.LottoTypeKLSF:
			lottoBall, err = GenLotto(ballKL10, 8, 1)
		case lotto.LottoTypeKENO:
			lottoBall, err = GenLotto(ballKENO, 20, 1)
		case lotto.LottoTypeLHC:
			lottoBall, err = GenLotto(ballLHC, 7, 1)
		default:
			return nil, errors.New(fmt.Sprintf("未找到合适的开奖器, 彩票编号: %d", lottoID))
		}
		return lottoBall, err
	}

	if err != nil {
		return []int{}, errors.New("找不到彩票信息")
	}
	return lottoBall, nil
}

func sendResult(lottoID int, issue string, drawNumber string) bool {
	url, err := common.Conf.String("url", "lotto_calc_myself")
	fmt.Println("URL:>", url)
	post := map[string]string{}
	post["lotto_id"] = fmt.Sprintf("%d", lottoID)
	post["issue"] = issue
	post["numbers"] = drawNumber

	body, err := common.HttpPost(url, "application/x-www-form-urlencoded;charset=UTF-8", post)
	if err == nil {
		fmt.Println(lottoID, drawNumber, string(body))
	}
	return true
}

func SelfCalcLottoV1(lottoID int) {
	lottoInfo, err := lotto.GetLotto(lottoID)
	if err != nil {
		fmt.Println("没有找到此彩票", err)
	}
	for {
		lottoIssueList, err := lotto.FindLottoResultList(lottoID, map[string]string{
			"start_time_to": common.GetTimeNowString(),
			"status":        "0",
			"order_by":      "issue",
			"sort_type":     "ASC",
		})
		if err != nil {
			fmt.Println(err)
			continue
		}
		for i := range *lottoIssueList {
			issue := (*lottoIssueList)[i].Issue
			resultTime := (*lottoIssueList)[i].ResultTime

			// 判断是否到达开奖时间
			if common.Time2Str(time.Now()) < resultTime {
				continue
			}

			// 计算号码
			lottoResultList, err := DoPreCalc(lottoID, issue)
			if err != nil {
				fmt.Println(err)
			}
			lottoResultStrList := formatNumbers(lottoInfo.LottoType, lottoResultList)
			lottoResultStr := strings.Join(lottoResultStrList, ",")

			// 发送号码
			sendResult(lottoID, issue, lottoResultStr)
		}
		// 计算下期开奖时间
		curIssueInfo, err := lotto.GetCurIssue(lottoID)
		nextDrawTime, err := common.Str2Time(curIssueInfo.ResultTime)
		if err != nil {
			nextDrawTime = time.Now()
		}
		d, _ := time.ParseDuration(fmt.Sprintf("%ds", GenerateRangeNum(2, 5)))
		nextDrawTime = nextDrawTime.Add(d)
		var subSeconds time.Duration
		subSeconds = time.Duration(nextDrawTime.Sub(time.Now()).Seconds())
		offsetSec := time.Duration(rand.Intn(5))
		if subSeconds > 0 {
			time.Sleep(subSeconds*time.Second + offsetSec*time.Second)
		} else {
			time.Sleep(time.Second * 3)
		}
	}
}
