package models

import (
	"database/sql"
	"fmt"
	"github.com/xiuos/lotto"
)

type Order = lotto.Order

const (
	TableRecordLottoOrder = "record_lotto_order"
)

func CreateRecordLottoOrderTx(tx *sql.Tx, o *lotto.Order) error {
	createSql := fmt.Sprintf("INSERT INTO %s SET  order_no = ?, user_id = ?, username = ?, lotto_id = ?, game_kind = ?, game_type = ?, issue = ?, method_code = ?, count = ?, content = ?, win_content = ?, win_count = ?, draw_number = ?, odds = ?, amount = ?, status = ?, payout = ?, profit = ?, flag = ?, bet_date = ?, calc_date = ?, bet_time = ?, update_time = ?", TableRecordLottoOrder)
	_, err := tx.Exec(createSql, o.OrderNo, o.UserID, o.Username, o.LottoID, o.GameKind, o.GameType, o.Issue, o.MethodCode, o.Count, o.Content, o.WinContent, o.WinCount, o.DrawNumber, o.Odds, o.Amount, o.Status, o.Payout, o.Profit, o.Flag, o.BetDate, o.CalcDate, o.BetTime, o.UpdateTime)
	return err
}
