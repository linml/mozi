package service

import (
	"github.com/xiuos/mozi/models"
	"github.com/xiuos/mozi/models/lotto"
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
