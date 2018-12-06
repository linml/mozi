package lotto

import (
	"database/sql"
	"fmt"
	"github.com/shopspring/decimal"
)

type Order struct {
	ID         int             `json:"id"`
	OrderNo    string          `json:"order_no"`
	UserID     int             `json:"user_id"`
	Username   string          `json:"username"`
	LottoID    int             `json:"lotto_id"`
	GameKind   int             `json:"game_kind"`
	GameType   int             `json:"game_type"`
	Issue      string          `json:"issue"`
	MethodCode string          `json:"method_code"`
	Count      int             `json:"count"`
	Content    string          `json:"content"`
	WinContent string          `json:"win_content"`
	WinCount   int             `json:"win_count"`
	DrawNumber string          `json:"draw_number"`
	Odds       decimal.Decimal `json:"odds"`
	Amount     decimal.Decimal `json:"amount"`
	Status     int             `json:"status"`
	Payout     decimal.Decimal `json:"payout"`
	Profit     decimal.Decimal `json:"profit"`
	Flag       int             `json:"flag"`
	BetDate    string          `json:"bet_date"`
	CalcDate   string          `json:"calc_date"`
	BetTime    string          `json:"bet_time"`
	UpdateTime string          `json:"update_time"`
}

func (d *Order) TableName() string {
	return "record_lotto_order"
}

func (d *Order) Field() []string {
	return []string{"id", "order_no", "user_id", "username", "lotto_id", "game_kind", "game_type", "issue", "method_code", "count", "content", "win_content", "win_count", "draw_number", "odds", "amount", "status", "payout", "profit", "flag", "bet_date", "calc_date", "bet_time", "update_time"}
}

func (d *Order) FieldItem() []interface{} {
	return []interface{}{&d.ID, &d.OrderNo, &d.UserID, &d.Username, &d.LottoID, &d.GameKind, &d.GameType, &d.Issue, &d.MethodCode, &d.Count, &d.Content, &d.WinContent, &d.WinCount, &d.DrawNumber, &d.Odds, &d.Amount, &d.Status, &d.Payout, &d.Profit, &d.Flag, &d.BetDate, &d.CalcDate, &d.BetTime, &d.UpdateTime}
}

func CreateRecordLottoOrderTx(tx *sql.Tx, o *Order) error {
	createSql := fmt.Sprintf("INSERT INTO %s SET  order_no = ?, user_id = ?, username = ?, lotto_id = ?, game_kind = ?, game_type = ?, issue = ?, method_code = ?, count = ?, content = ?, win_content = ?, win_count = ?, draw_number = ?, odds = ?, amount = ?, status = ?, payout = ?, profit = ?, flag = ?, bet_date = ?, calc_date = ?, bet_time = ?, update_time = ?", o.TableName())
	_, err := tx.Exec(createSql, o.OrderNo, o.UserID, o.Username, o.LottoID, o.GameKind, o.GameType, o.Issue, o.MethodCode, o.Count, o.Content, o.WinContent, o.WinCount, o.DrawNumber, o.Odds, o.Amount, o.Status, o.Payout, o.Profit, o.Flag, o.BetDate, o.CalcDate, o.BetTime, o.UpdateTime)
	return err
}
