package models

import (
	"fmt"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/models/lotto"
)

const (
	TableIssueFactory = "issue_factory"
)

type IssueFactory = lotto.IssueFactory

func FindIssueFactory(params map[string]string) (*[]IssueFactory, error) {
	var ifl []IssueFactory
	querySql := fmt.Sprintf("SELECT  lotto_id, status, count, issue_interval, "+
		"iss_bit, block_sec, start_time, end_time, issue_type, offset, extra_info FROM %s WHERE 1=1 ", TableIssueFactory)

	rows, err := common.BaseDb.Query(querySql)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		i := IssueFactory{}
		err = rows.Scan(&i.LottoID, &i.Status, &i.Count, &i.IssueInterval, &i.IssBit, &i.BlockSec, &i.StartTime, &i.EndTime, &i.IssueType, &i.Offset, &i.ExtraInfo)
		if err != nil {
			return nil, err
		}
		ifl = append(ifl, i)
	}
	return &ifl, nil
}
