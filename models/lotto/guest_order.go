package lotto

import (
	"database/sql"
	"fmt"
	"github.com/xiuos/mozi/common"
	"strings"
)

const GuestHead = "guest_"

func CreateGuestRecordLottoOrderTx(tx *sql.Tx, o *Order) error {
	createSql := fmt.Sprintf("INSERT INTO %s SET order_id=?, user_id=?, name=?, lotto_id=?, lotto_type=?, game_kind=?, game_type=?, issue=?, method_code=?, play_code=?, bet_count=?, bet_content=?, win_count=?, win_content=?, draw_number=?, odds=?, amount=?, status=?, flag=?, payout=?, profit=?, bet_date=?, calc_date=?, bet_time=?, update_time=?, ip=?", o.TableName(GuestHead))
	_, err := tx.Exec(createSql, &o.OrderID, &o.UserID, &o.Name, &o.LottoID, &o.LottoType, &o.GameKind, &o.GameType, &o.Issue, &o.MethodCode, &o.PlayCode, &o.BetCount, &o.BetContent, &o.WinCount, &o.WinContent, &o.DrawNumber, &o.Odds, &o.Amount, &o.Status, &o.Flag, &o.Payout, &o.Profit, &o.BetDate, &o.CalcDate, &o.BetTime, &o.UpdateTime, &o.IP)
	return err
}

func FindGuestOrder4Settle(lastID int, lottoID int, issue string) (*[]Order, error) {
	t := Order{}
	querySql := fmt.Sprintf("SELECT %s FROM %s WHERE lotto_id=? AND issue=? AND id >? ORDER BY id ASC LIMIT 100", strings.Join(t.Field(), ","), t.TableName(GuestHead))

	var data []Order

	rows, err := common.BaseDb.Query(querySql, lottoID, issue, lastID)
	if err != nil {
		return &data, err
	}
	for rows.Next() {
		d := Order{}
		err = rows.Scan(d.FieldItem()...)
		if err != nil {
			return &data, err
		}
		data = append(data, d)
	}
	return &data, err
}
