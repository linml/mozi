package lotto

import (
	"database/sql"
	"fmt"
	"github.com/shopspring/decimal"
)

type Order struct {
	ID         int             `json:"id"`
	OrderID    string          `json:"order_id"`
	UserID     int             `json:"user_id"`
	Name       string          `json:"name"`
	LottoID    int             `json:"lotto_id"`
	LottoType  int             `json:"lotto_type"`
	GameKind   int             `json:"game_kind"`
	GameType   int             `json:"game_type"`
	Issue      string          `json:"issue"`
	MethodCode string          `json:"method_code"`
	PlayCode   string          `json:"play_code"`
	BetCount   int             `json:"bet_count"`
	BetContent string          `json:"bet_content"`
	WinCount   int             `json:"win_count"`
	WinContent string          `json:"win_content"`
	DrawNumber string          `json:"draw_number"`
	Odds       decimal.Decimal `json:"odds"`
	Amount     decimal.Decimal `json:"amount"`
	Status     int             `json:"status"`
	Flag       int             `json:"flag"`
	Payout     decimal.Decimal `json:"payout"`
	Profit     decimal.Decimal `json:"profit"`
	BetDate    string          `json:"bet_date"`
	CalcDate   string          `json:"calc_date"`
	BetTime    string          `json:"bet_time"`
	UpdateTime string          `json:"update_time"`
	IP         int64           `json:"ip"`
}

func (d *Order) TableName() string {
	return "record_lotto_order"
}

func (d *Order) Field() []string {
	return []string{"id", "order_id", "user_id", "name", "lotto_id", "lotto_type", "game_kind", "game_type", "issue", "method_code", "play_code", "bet_count", "bet_content", "win_count", "win_content", "draw_number", "odds", "amount", "status", "flag", "payout", "profit", "bet_date", "calc_date", "bet_time", "update_time", "ip"}
}

func (d *Order) FieldItem() []interface{} {
	return []interface{}{&d.ID, &d.OrderID, &d.UserID, &d.Name, &d.LottoID, &d.LottoType, &d.GameKind, &d.GameType, &d.Issue, &d.MethodCode, &d.PlayCode, &d.BetCount, &d.BetContent, &d.WinCount, &d.WinContent, &d.DrawNumber, &d.Odds, &d.Amount, &d.Status, &d.Flag, &d.Payout, &d.Profit, &d.BetDate, &d.CalcDate, &d.BetTime, &d.UpdateTime, &d.IP}
}

func CreateRecordLottoOrderTx(tx *sql.Tx, o *Order) error {
	createSql := fmt.Sprintf("INSERT INTO %s SET order_id=?, user_id=?, name=?, lotto_id=?, lotto_type=?, game_kind=?, game_type=?, issue=?, method_code=?, play_code=?, bet_count=?, bet_content=?, win_count=?, win_content=?, draw_number=?, odds=?, amount=?, status=?, flag=?, payout=?, profit=?, bet_date=?, calc_date=?, bet_time=?, update_time=?, ip=?", o.TableName())
	_, err := tx.Exec(createSql, &o.OrderID, &o.UserID, &o.Name, &o.LottoID, &o.LottoType, &o.GameKind, &o.GameType, &o.Issue, &o.MethodCode, &o.PlayCode, &o.BetCount, &o.BetContent, &o.WinCount, &o.WinContent, &o.DrawNumber, &o.Odds, &o.Amount, &o.Status, &o.Flag, &o.Payout, &o.Profit, &o.BetDate, &o.CalcDate, &o.BetTime, &o.UpdateTime, &o.IP)
	return err
}
