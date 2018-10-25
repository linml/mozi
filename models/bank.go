package models

import "database/sql"

type Bank struct {
	ID       int    `json:"id"`
	BankCode int    `json:"bank_code"`
	BankName string `json:"bank_name"`
	Group    string `json:"group"`
	Status   string `json:"status"`
}

func (b *Bank) TableName() string {
	return "code_bank"
}

func (b *Bank) Field() []string {
	return []string{"id", "bank_code", "bank_name", "group", "status"}
}

func (b *Bank) FieldItem() []interface{} {
	return []interface{}{&b.ID, &b.BankCode, &b.BankName, &b.Group, &b.Status}
}

func CreatebankTx(tx *sql.Tx, b *Bank) (sql.Result, error) {
	createSql := "INSERT INTO " + b.TableName() + " SET bank_code = ?, bank_name = ?, group = ?, status = ?"
	rs, err := tx.Exec(createSql, &b.BankCode, &b.BankName, &b.Group, &b.Status)
	if err != nil {
		return nil, err
	}
	return rs, nil
}
