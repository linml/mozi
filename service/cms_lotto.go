package service

import (
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/models"
	"github.com/xiuos/mozi/models/lotto"
	"strings"
)

type HallLotto struct {
	LottoID int    `json:"lotto_id"`
	Name    string `json:"name"`
	Issue   string `json:"issue"`
	ImgUrl  string `json:"img_url"`
}

type HomeLotto struct {
	LottoID int    `json:"lotto_id"`
	Name    string `json:"name"`
	Issue   string `json:"issue"`
	ImgUrl  string `json:"img_url"`
}

func CMSFindHallLottoList() (*[]HallLotto, error) {
	dataList := []HallLotto{}
	hl, err := lotto.FindCMSHallLotto(map[string]string{"status": "1"})
	if err != nil {
		return &dataList, err
	}
	for i := range *hl {
		lid := (*hl)[i].LottoID
		lottoInfo, err := lotto.GetLotto(lid)
		if err != nil {
			continue
		}
		issueInfo, err := lotto.GetCurIssue(lid)
		if err != nil {

		}
		data := HallLotto{
			LottoID: lottoInfo.LottoID,
			Name:    lottoInfo.Name,
			Issue:   issueInfo.Issue,
			ImgUrl:  lottoInfo.ImgUrl,
		}
		dataList = append(dataList, data)
	}
	return &dataList, err
}

//获取首页初始化数据
// 返回
// banner_list 轮播图列表
// box_notice_list 弹窗
// lotto_list 彩票列表
func CMSHomeInit() (interface{}, error) {
	ret := map[string]interface{}{}
	dataList := []HomeLotto{}
	hl, err := lotto.FindCMSHomeLotto(map[string]string{"status": "1"})
	if err != nil {
		return &dataList, err
	}
	for i := range *hl {
		lid := (*hl)[i].LottoID
		lottoInfo, err := lotto.GetLotto(lid)
		if err != nil {
			continue
		}
		issueInfo, err := lotto.GetCurIssue(lid)
		if err != nil {

		}
		data := HomeLotto{
			LottoID: lottoInfo.LottoID,
			Name:    lottoInfo.Name,
			Issue:   issueInfo.Issue,
			ImgUrl:  lottoInfo.ImgUrl,
		}
		dataList = append(dataList, data)
	}

	ret["lotto_list"] = dataList

	bannerList := []string{}
	gtl, err := models.FindGiftTaskList(map[string]string{"status": "1"})
	for i := range *gtl {
		if (*gtl)[i].ShowBanner == 1 {
			bannerList = append(bannerList, (*gtl)[i].ImgUrl)
		}
	}
	ret["banner_list"] = bannerList

	return &ret, err
}

// 彩票列表
func CMSCodeLottoList() (*[]HomeLotto, error) {
	dataList := []HomeLotto{}
	hl, err := lotto.FindCMSHomeLotto(map[string]string{"status": "1"})
	if err != nil {
		return &dataList, err
	}
	for i := range *hl {
		lid := (*hl)[i].LottoID
		lottoInfo, err := lotto.GetLotto(lid)
		if err != nil {
			continue
		}
		issueInfo, err := lotto.GetCurIssue(lid)
		if err != nil {

		}
		data := HomeLotto{
			LottoID: lottoInfo.LottoID,
			Name:    lottoInfo.Name,
			Issue:   issueInfo.Issue,
			ImgUrl:  lottoInfo.ImgUrl,
		}
		dataList = append(dataList, data)
	}
	return &dataList, err
}

type outGroup struct {
	ID          int                  `json:"id"`
	LottoID     int                  `json:"lotto_id"`
	GroupID     int                  `json:"group_id"`
	GroupName   string               `json:"group_name"`
	GroupAlias  string               `json:"group_alias"`
	SortIndex   int                  `json:"sort_index"`
	Status      int                  `json:"status"`
	MethodItems []outMethodGroupPlay `json:"method_items"`
}

type outMethodGroupPlay struct {
	ID          int           `json:"id"`
	LottoID     int           `json:"lotto_id"`
	MethodCode  string        `json:"method_code"`
	MethodName  string        `json:"method_name"`
	MethodAlias string        `json:"method_alias"`
	SortIndex   int           `json:"sort_index"`
	Status      int           `json:"status"`
	PlayItems   []outOddsInfo `json:"play_items"`
}

type outOddsInfo struct {
	MethodCode string          `json:"method_code"`
	MethodName string          `json:"method_name"`
	PlayCode   string          `json:"play_code"`
	PlayName   string          `json:"play_name"`
	Odds       decimal.Decimal `json:"odds"`
	BetMin     decimal.Decimal `json:"bet_min"`
	BetMax     decimal.Decimal `json:"bet_max"`
}

func CMSBetPlayInfo(uid int, lid int) (interface{}, error) {
	ret := map[string]interface{}{}

	curIssueInfo, _ := lotto.GetCurIssue(lid)
	lastIssueInfo, _ := lotto.GetLastIssue(lid)
	lottoInfo, _ := lotto.GetLotto(lid)
	methodGroup, err := lotto.FindCMSLottoMethodGroupList(map[string]string{"lotto_id": fmt.Sprintf("%d", lid), "status": "1"})
	if err != nil {
		fmt.Println(err)
	}

	outgl := []outGroup{}

	for i := range *methodGroup {
		outg := outGroup{
			ID:         (*methodGroup)[i].ID,
			LottoID:    (*methodGroup)[i].LottoID,
			GroupID:    (*methodGroup)[i].GroupID,
			GroupName:  (*methodGroup)[i].GroupName,
			GroupAlias: (*methodGroup)[i].GroupAlias,
			SortIndex:  (*methodGroup)[i].SortIndex,
			Status:     (*methodGroup)[i].Status,
		}
		mgpl, _ := lotto.FindCMSLottoMethodGroupPlayList(map[string]string{
			"lotto_id": fmt.Sprintf("%d", lid),
			"group_id": fmt.Sprintf("%d", (*methodGroup)[i].GroupID),
			"status":   "1"})
		outmgpl := []outMethodGroupPlay{}
		for j := range *mgpl {
			outmgp := outMethodGroupPlay{
				ID:          (*mgpl)[j].ID,
				LottoID:     (*mgpl)[j].LottoID,
				MethodCode:  (*mgpl)[j].MethodCode,
				MethodName:  (*mgpl)[j].MethodName,
				MethodAlias: (*mgpl)[j].MethodAlias,
				SortIndex:   (*mgpl)[j].SortIndex,
				Status:      (*mgpl)[j].Status,
			}
			oddsInfoList, _ := lotto.FindLottoOddsList(map[string]string{
				"lotto_id":    fmt.Sprintf("%d", lid),
				"method_code": outmgp.MethodCode,
				"order_by":    "sort_index",
				"sort_type":   "ASC",
			})
			outoil := []outOddsInfo{}
			for k := range *oddsInfoList {
				outoi := outOddsInfo{
					MethodCode: (*oddsInfoList)[k].MethodCode,
					MethodName: (*oddsInfoList)[k].MethodName,
					PlayCode:   (*oddsInfoList)[k].PlayCode,
					PlayName:   (*oddsInfoList)[k].PlayName,
					Odds:       (*oddsInfoList)[k].Odds,
					BetMin:     (*oddsInfoList)[k].BetMin,
					BetMax:     (*oddsInfoList)[k].BetMax,
				}
				outoil = append(outoil, outoi)
			}
			outmgp.PlayItems = outoil
			outmgpl = append(outmgpl, outmgp)
		}
		outg.MethodItems = outmgpl
		outgl = append(outgl, outg)
	}

	num := []string{}
	if lastIssueInfo.DrawNumber != "" {
		num = strings.Split(lastIssueInfo.DrawNumber, ",")
	}

	ret["server_time"] = common.GetTimeNowString()
	ret["curr_issue"] = &curIssueInfo
	ret["lotto_info"] = &lottoInfo
	ret["method_group"] = &methodGroup
	ret["method_group_play"] = outgl
	ret["last_issue"] = map[string]interface{}{
		"issue": lastIssueInfo.Issue,
		"num":   num,
	}

	return ret, nil
}
