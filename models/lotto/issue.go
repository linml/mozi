package lotto

import (
	"errors"
	"fmt"
	"github.com/nosixtools/solarlunar"
	"strings"
	"time"
)

const (
	IssueType1 = 1 // 每天期号在同一天, 时间间隔相同
	IssueType2 = 2 // 期号夸天生成, 时间间隔相同
	IssueType3 = 3 // 期号夸天生成, 时间间隔不相同
	IssueType4 = 4 // 每天期号在同一天, 时间间隔相同, 期号永久累计自增
	IssueType5 = 5 // 每天一期, 春节不开奖
	IssueType6 = 6 // 每天期号在同一天, 时间间隔不相同 重庆彩
)

type IssueFactory struct {
	LottoID       int    `json:"lotto_id"`       // 彩票编号
	Status        int    `json:"status"`         // 状态
	Count         int    `json:"count"`          // 期数
	IssueInterval int    `json:"issue_interval"` // 单位秒
	IssBit        int    `json:"iss_bit"`        // 期号位数，去除时间
	BlockSec      int    `json:"block_sec"`      // 封单时间秒
	StartTime     string `json:"start_time"`     // 开售起始时间
	EndTime       string `json:"end_time"`       // 结束时间
	IssueType     int    `json:"issue_type"`     // 期号类型
	Offset        int    `json:"offset"`         // 期号封单时间偏移
	ExtraInfo     string `json:"extra_info"`     // 额外信息
}

type IssueInfo struct {
	LottoID    int    `json:"lotto_id"`
	Issue      string `json:"issue"`
	DrawNumber string `json:"draw_number"`
	Status     int    `json:"status"`
	StartTime  string `json:"start_time"`
	CloseTime  string `json:"close_time"`
	EndTime    string `json:"end_time"`
	DrawTime   string `json:"draw_time"`
}

func (issFac *IssueFactory) GenIssue(t time.Time) (*[]IssueInfo, error) {
	var lrl []IssueInfo
	switch issFac.IssueType {
	case IssueType1:
		headIssue := t.Format("20060102")
		d, _ := time.ParseDuration("-24h")
		Index, _ := time.Parse("20060102150405", headIssue+issFac.StartTime)
		StartTime, _ := time.Parse("20060102150405", headIssue+issFac.EndTime)
		d, _ = time.ParseDuration("-24h")
		StartTime = StartTime.Add(d)
		for i := 1; i < issFac.Count+1; i++ {

			d, _ = time.ParseDuration(fmt.Sprintf("%ds", issFac.IssueInterval))
			Index = Index.Add(d)

			EndTime := Index
			LockTime := Index

			d, _ = time.ParseDuration(fmt.Sprintf("-%ds", issFac.BlockSec))
			LockTime = Index.Add(d)

			iPre := "%s%0" + fmt.Sprintf("%d", issFac.IssBit) + "d"
			issue := fmt.Sprintf(iPre, headIssue, i)
			startTimeStr := StartTime.Format("20060102150405")
			lockTimeStr := LockTime.Format("20060102150405")
			endTimeStr := EndTime.Format("20060102150405")
			lrl = append(lrl, IssueInfo{
				LottoID:   issFac.LottoID,
				Issue:     issue,
				StartTime: startTimeStr,
				CloseTime: lockTimeStr,
				EndTime:   endTimeStr,
			})
			StartTime = EndTime
		}
		return &lrl, nil

	case IssueType2:
		headIssue := t.Format("20060102")
		d, _ := time.ParseDuration("-24h")
		Index, _ := time.Parse("20060102150405", headIssue+issFac.StartTime)
		StartTime, _ := time.Parse("20060102150405", headIssue+issFac.EndTime)
		d, _ = time.ParseDuration("-24h")
		//StartTime = StartTime.Add(d)
		for i := 1; i < issFac.Count+1; i++ {

			d, _ = time.ParseDuration(fmt.Sprintf("%ds", issFac.IssueInterval))
			Index = Index.Add(d)

			EndTime := Index
			LockTime := Index

			d, _ = time.ParseDuration(fmt.Sprintf("-%ds", issFac.BlockSec))
			LockTime = Index.Add(d)

			iPre := "%s%0" + fmt.Sprintf("%d", issFac.IssBit) + "d"
			issue := fmt.Sprintf(iPre, headIssue, i)
			startTimeStr := StartTime.Format("20060102150405")
			lockTimeStr := LockTime.Format("20060102150405")
			endTimeStr := EndTime.Format("20060102150405")
			lrl = append(lrl, IssueInfo{
				LottoID:   issFac.LottoID,
				Issue:     issue,
				StartTime: startTimeStr,
				CloseTime: lockTimeStr,
				EndTime:   endTimeStr,
			})
			StartTime = EndTime
		}
		return &lrl, nil

	case IssueType3:

	case IssueType4:
		extraInfo := issFac.ExtraInfo
		extraIssueInfo := strings.Split(extraInfo, "|")
		if len(extraIssueInfo) != 2 {
			return &lrl, errors.New("检查额外信息是否正确")
		}

		pointIssue := GetInt(extraIssueInfo[0])
		pointIssueTime := extraIssueInfo[1]

		preTime, _ := time.Parse("20060102150405", pointIssueTime)
		preIssue := pointIssue
		diffT := int(t.Sub(preTime).Hours()/24) + 1

		headIssue := t.Format("20060102")
		d, _ := time.ParseDuration("-24h")

		Index, _ := time.Parse("20060102150405", headIssue+issFac.StartTime)
		StartTime, _ := time.Parse("20060102150405", headIssue+issFac.EndTime)
		d, _ = time.ParseDuration("-24h")
		StartTime = StartTime.Add(d)
		for i := 1; i < issFac.Count+1; i++ {

			d, _ = time.ParseDuration(fmt.Sprintf("%ds", issFac.IssueInterval))
			Index = Index.Add(d)

			EndTime := Index
			LockTime := Index

			d, _ = time.ParseDuration(fmt.Sprintf("-%ds", issFac.BlockSec))
			LockTime = Index.Add(d)

			issue := fmt.Sprintf("%d", preIssue+issFac.Count*(diffT-1)+i)
			startTimeStr := StartTime.Format("20060102150405")
			lockTimeStr := LockTime.Format("20060102150405")
			endTimeStr := EndTime.Format("20060102150405")
			lrl = append(lrl, IssueInfo{
				LottoID:   issFac.LottoID,
				Issue:     issue,
				StartTime: startTimeStr,
				CloseTime: lockTimeStr,
				EndTime:   endTimeStr,
			})
			StartTime = EndTime
		}
		return &lrl, nil
	case IssueType5:

		headIssue := t.Format("20060102")
		d, _ := time.ParseDuration("-24h")

		lunarDate, f := solarlunar.SolarToLuanr(t.Format("2006-01-02"))

		_lunarDate := lunarDate[:4] + "-01-01"

		solarDate := solarlunar.LunarToSolar(_lunarDate, f)

		springBegin, _ := time.Parse("20060102", solarDate)
		springFestival, _ := time.ParseDuration("168h")
		springEnd := springBegin.Add(springFestival)

		if (t.Before(springEnd) && t.After(springBegin)) || t.Equal(springEnd) || t.Equal(springBegin) {
			return &lrl, errors.New("春节不开售")
		}

		StartTime, _ := time.Parse("20060102150405", headIssue+issFac.EndTime)
		d, _ = time.ParseDuration("-24h")
		StartTime = StartTime.Add(d)

		EndTime, _ := time.Parse("20060102150405", headIssue+issFac.EndTime)

		d, _ = time.ParseDuration(fmt.Sprintf("-%ds", issFac.BlockSec))
		LockTime := EndTime.Add(d)

		issue := ""
		if t.After(springEnd) {
			toyearb, _ := time.Parse("20060102", t.Format("2006")+"0101")
			issue = fmt.Sprintf("%s%03d", headIssue[:4], int(t.Sub(toyearb).Hours()/24)-int(springEnd.Sub(springBegin).Hours()/24)+1)
		} else {
			toyearb, _ := time.Parse("20060102", t.Format("2006")+"0101")
			issue = fmt.Sprintf("%s%03d", headIssue[:4], int(t.Sub(toyearb).Hours()/24)+1)
		}

		lrl = append(lrl, IssueInfo{
			LottoID:   issFac.LottoID,
			Issue:     issue,
			StartTime: StartTime.Format("20060102150405"),
			CloseTime: LockTime.Format("20060102150405"),
			EndTime:   EndTime.Format("20060102150405"),
		})
		return &lrl, nil
	case IssueType6:
		headIssue := t.Format("20060102")
		d, _ := time.ParseDuration("-24h")
		Index, _ := time.Parse("20060102150405", headIssue+issFac.StartTime)
		StartTime, _ := time.Parse("20060102150405", headIssue+issFac.EndTime)

		for i := 1; i < issFac.Count+1; i++ {
			if i < 24 {
				d, _ = time.ParseDuration(fmt.Sprintf("%ds", 300))
			} else if i == 24 {
				d, _ = time.ParseDuration(fmt.Sprintf("%ds", 29100))
			} else if i < 97 {
				d, _ = time.ParseDuration(fmt.Sprintf("%ds", 600))
			} else {
				d, _ = time.ParseDuration(fmt.Sprintf("%ds", 300))
			}

			Index = Index.Add(d)

			EndTime := Index
			LockTime := Index

			if i < 24 || i > 96 {
				d, _ = time.ParseDuration(fmt.Sprintf("-%ds", issFac.BlockSec/2))
			} else {
				d, _ = time.ParseDuration(fmt.Sprintf("-%ds", issFac.BlockSec))
			}
			LockTime = Index.Add(d)

			iPre := "%s%0" + fmt.Sprintf("%d", issFac.IssBit) + "d"
			issue := fmt.Sprintf(iPre, headIssue, i)
			startTimeStr := StartTime.Format("20060102150405")
			lockTimeStr := LockTime.Format("20060102150405")
			endTimeStr := EndTime.Format("20060102150405")
			lrl = append(lrl, IssueInfo{
				LottoID:   issFac.LottoID,
				Issue:     issue,
				StartTime: startTimeStr,
				CloseTime: lockTimeStr,
				EndTime:   endTimeStr,
			})
			StartTime = EndTime
		}
		return &lrl, nil
	default:
		return &lrl, nil
	}
	return &lrl, nil
}
