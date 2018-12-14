package service

import (
	"fmt"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/models"
	"github.com/xiuos/mozi/models/lotto"
)

func RefreshLottoDayCount(countDate string) error {

	lottoDayCountList, err := lotto.FindLottoOrderDayCount(countDate)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	for i := range *lottoDayCountList {
		relationInfo, err := models.GetUserRelation((*lottoDayCountList)[i].UserID)
		if err != nil {
			fmt.Println(err)
		}
		rp := models.ReportLottoDayCount{
			CountDate:   (*lottoDayCountList)[i].CountDate,
			UserID:      (*lottoDayCountList)[i].UserID,
			Name:        (*lottoDayCountList)[i].Name,
			ParentID:    relationInfo.ParentID,
			Parents:     relationInfo.Parents,
			LottoID:     (*lottoDayCountList)[i].LottoID,
			GameKind:    (*lottoDayCountList)[i].GameKind,
			GameType:    (*lottoDayCountList)[i].GameType,
			TotalCount:  (*lottoDayCountList)[i].TotalCount,
			TotalBet:    (*lottoDayCountList)[i].TotalBet,
			TotalPayout: (*lottoDayCountList)[i].TotalPayout,
			TotalProfit: (*lottoDayCountList)[i].TotalProfit,
			UpdateTime:  common.GetTimeNowString(),
		}
		models.ReplaceLottoDayCount(&rp)
	}

	return nil
}

func PageFindAgentLottoReportList4Admin(params common.PageParams) (*common.PageResult, *[]models.AgentLottoReport, error) {
	if v, ok := params.Params["parent_name"]; ok {
		info, err := models.GetUserByName(v)
		if err == nil {
			params.Params["parent_id"] = fmt.Sprintf("%d", info.UserID)
		}
	}
	if v, ok := params.Params["name"]; ok {
		info, err := models.GetUserByName(v)
		if err == nil {
			params.Params["user_id"] = fmt.Sprintf("%d", info.UserID)
		}
	}

	return models.PageFindAgentLottoReportList(params)
}
