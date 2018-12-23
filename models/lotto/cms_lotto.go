package lotto

type CMSHallLotto struct {
	LottoID   int `json:"lotto_id"`
	Status    int `json:"status"`
	SortIndex int `json:"sort_index"`
}

func FindCMSHallLotto(params map[string]string) (*[]CMSHallLotto, error) {
	dl := []CMSHallLotto{}
	lottoInfoList, err := FindCodeLottoList(map[string]string{
		"status":    "1",
		"order_by":  "extra_1_sort_index",
		"sort_type": "ASC",
	})
	if err != nil {
		return &dl, err
	}

	for i := range *lottoInfoList {
		dl = append(dl, CMSHallLotto{
			LottoID:   (*lottoInfoList)[i].LottoID,
			Status:    (*lottoInfoList)[i].Status,
			SortIndex: (*lottoInfoList)[i].Extra1SortIndex,
		})
	}

	return &dl, err
}

type CMSHomeLotto struct {
	LottoID   int `json:"lotto_id"`
	Status    int `json:"status"`
	SortIndex int `json:"sort_index"`
}

func FindCMSHomeLotto(params map[string]string) (*[]CMSHomeLotto, error) {
	dl := []CMSHomeLotto{}
	lottoInfoList, err := FindCodeLottoList(map[string]string{
		"status":    "1",
		"order_by":  "extra_2_sort_index",
		"sort_type": "ASC",
	})
	if err != nil {
		return &dl, err
	}

	for i := range *lottoInfoList {
		dl = append(dl, CMSHomeLotto{
			LottoID:   (*lottoInfoList)[i].LottoID,
			Status:    (*lottoInfoList)[i].Status,
			SortIndex: (*lottoInfoList)[i].Extra1SortIndex,
		})
	}

	return &dl, err
}
