package models

import (
	"database/sql"
	"github.com/shopspring/decimal"
)

const (
	TableRecordMonyeChange = "record_money_change"
)

type RecordMoneyChange struct {
	ID            int             `json:"id"`
	Link          string          `json:"link"`
	UserID        int             `json:"user_id"`
	Username      string          `json:"username"`
	GameKind      int             `json:"game_kind"`
	GameType      int             `json:"game_type"`
	ChangeKind    int             `json:"change_kind"`
	ChangeType    int             `json:"change_type"`
	BeforeBalance decimal.Decimal `json:"before_balance"`
	Amount        decimal.Decimal `json:"amount"`
	AfterBalance  decimal.Decimal `json:"after_balance"`
	RecordDate    string          `json:"record_date"`
	RecordAt      string          `json:"record_at"`
	OperateType   int             `json:"operate_type"`
	OperatorID    int             `json:"operator_id"`
	Remark        string          `json:"remark"`
}

func (r *RecordMoneyChange) Field() []string {
	return []string{"id", "link", "user_id", "username", "game_kind", "game_type", "change_kind", "change_type", "before_balance", "amount", "after_balance", "record_date", "record_at", "operate_type", "operator_id", "remark"}
}

func (r *RecordMoneyChange) FieldItem() []interface{} {
	return []interface{}{&r.ID, &r.Link, &r.UserID, &r.Username, &r.GameKind, &r.GameType, &r.ChangeKind, &r.ChangeType, &r.BeforeBalance, &r.Amount, &r.AfterBalance, &r.RecordDate, &r.RecordAt, &r.OperateType, &r.OperatorID, &r.Remark}
}

func CreateRecordMoneyChangeTx(tx *sql.Tx, r *RecordMoneyChange) (sql.Result, error) {
	createSql := "INSERT INTO " + TableRecordMonyeChange + " SET link = ?, user_id = ?, username = ?, game_kind = ?, game_type = ?, change_kind = ?, change_type = ?, before_balance = ?, amount = ?, after_balance = ?, record_date = ?, record_at = ?, operate_type = ?, operator_id = ?, remark = ?"
	rs, err := tx.Exec(createSql, r.Link, r.UserID, r.Username, r.GameKind, r.GameType, r.ChangeKind, r.ChangeType, r.BeforeBalance, r.Amount, r.AfterBalance, r.RecordDate, r.RecordAt, r.OperateType, r.OperatorID, r.Remark)
	if err != nil {
		return nil, err
	}
	return rs, nil
}
