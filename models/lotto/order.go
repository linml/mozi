package lotto

import (
	"database/sql"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/xiuos/mozi/common"
	"strings"
	"time"
)

const (
	OrderStatusUnsettle   = 0 // 订单未结算
	OrderStatusWinning    = 1 // 订单中奖
	OrderStatusNotWinning = 2 // 订单未中奖
	OrderStatusCancel     = 3 // 撤单
	OrderStatusAbnormal   = 4 // 异常
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

type LottoOrderCount struct {
	UserID      int             `json:"user_id"`
	Name        string          `json:"name"`
	LottoID     int             `json:"lotto_id"`
	LottoType   int             `json:"lotto_type"`
	GameKind    int             `json:"game_kind"`
	GameType    int             `json:"game_type"`
	TotalCount  int             `json:"total_count"`
	TotalBet    decimal.Decimal `json:"total_bet"`
	TotalPayout decimal.Decimal `json:"total_payout"`
	TotalProfit decimal.Decimal `json:"total_profit"`
	CountDate   string          `json:"count_date"`
}

type Winning struct {
	Name      string          `json:"name"`
	LottoID   int             `json:"lotto_id"`
	LottoName string          `json:"lotto_name"`
	Amount    decimal.Decimal `json:"amount"`
}

func (d *LottoOrderCount) Field() []string {
	return []string{"user_id", "name", "lotto_id", "lotto_type", "game_kind", "game_type", "total_count", "total_bet", "total_payout", "total_profit", "count_date"}
}

func (d *LottoOrderCount) FieldItem() []interface{} {
	return []interface{}{&d.UserID, &d.Name, &d.LottoID, &d.LottoType, &d.GameKind, &d.GameType, &d.TotalCount, &d.TotalBet, &d.TotalPayout, &d.TotalProfit, &d.CountDate}
}

func (d *Order) TableName(prefix string) string {
	return fmt.Sprintf("%srecord_lotto_order", prefix)
}

func (d *Order) Field() []string {
	return []string{"id", "order_id", "user_id", "name", "lotto_id", "lotto_type", "game_kind", "game_type", "issue", "method_code", "play_code", "bet_count", "bet_content", "win_count", "win_content", "draw_number", "odds", "amount", "status", "flag", "payout", "profit", "bet_date", "calc_date", "bet_time", "update_time", "ip"}
}

func (d *Order) FieldItem() []interface{} {
	return []interface{}{&d.ID, &d.OrderID, &d.UserID, &d.Name, &d.LottoID, &d.LottoType, &d.GameKind, &d.GameType, &d.Issue, &d.MethodCode, &d.PlayCode, &d.BetCount, &d.BetContent, &d.WinCount, &d.WinContent, &d.DrawNumber, &d.Odds, &d.Amount, &d.Status, &d.Flag, &d.Payout, &d.Profit, &d.BetDate, &d.CalcDate, &d.BetTime, &d.UpdateTime, &d.IP}
}

func CreateRecordLottoOrderTx(tx *sql.Tx, o *Order) error {
	createSql := fmt.Sprintf("INSERT INTO %s SET order_id=?, user_id=?, name=?, lotto_id=?, lotto_type=?, game_kind=?, game_type=?, issue=?, method_code=?, play_code=?, bet_count=?, bet_content=?, win_count=?, win_content=?, draw_number=?, odds=?, amount=?, status=?, flag=?, payout=?, profit=?, bet_date=?, calc_date=?, bet_time=?, update_time=?, ip=?", o.TableName(""))
	_, err := tx.Exec(createSql, &o.OrderID, &o.UserID, &o.Name, &o.LottoID, &o.LottoType, &o.GameKind, &o.GameType, &o.Issue, &o.MethodCode, &o.PlayCode, &o.BetCount, &o.BetContent, &o.WinCount, &o.WinContent, &o.DrawNumber, &o.Odds, &o.Amount, &o.Status, &o.Flag, &o.Payout, &o.Profit, &o.BetDate, &o.CalcDate, &o.BetTime, &o.UpdateTime, &o.IP)
	return err
}

func FindOrder4Settle(lastID int, lottoID int, issue string) (*[]Order, error) {
	t := Order{}
	querySql := fmt.Sprintf("SELECT %s FROM %s WHERE lotto_id=? AND issue=? AND id >? ORDER BY id ASC LIMIT 100", strings.Join(t.Field(), ","), t.TableName(""))

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

func FindLottoOrderList(params map[string]string) (*[]Order, error) {
	t := Order{}
	conditionSql := fmt.Sprintf(" FROM %s WHERE 1=1 ", t.TableName(""))
	querySql := fmt.Sprintf("SELECT %s ", strings.Join(t.Field(), ",")) + conditionSql
	sqlWhere, args := "", []interface{}{}
	if v, ok := params["id"]; ok {
		sqlWhere += " AND id = ?"
		args = append(args, v)
	}
	if v, ok := params["order_id"]; ok {
		sqlWhere += " AND order_id = ?"
		args = append(args, v)
	}
	if v, ok := params["user_id"]; ok {
		sqlWhere += " AND user_id = ?"
		args = append(args, v)
	}
	if v, ok := params["name"]; ok {
		sqlWhere += " AND name = ?"
		args = append(args, v)
	}
	if v, ok := params["lotto_id"]; ok {
		sqlWhere += " AND lotto_id = ?"
		args = append(args, v)
	}
	if v, ok := params["lotto_type"]; ok {
		sqlWhere += " AND lotto_type = ?"
		args = append(args, v)
	}
	if v, ok := params["game_kind"]; ok {
		sqlWhere += " AND game_kind = ?"
		args = append(args, v)
	}
	if v, ok := params["game_type"]; ok {
		sqlWhere += " AND game_type = ?"
		args = append(args, v)
	}
	if v, ok := params["issue"]; ok {
		sqlWhere += " AND issue = ?"
		args = append(args, v)
	}
	if v, ok := params["method_code"]; ok {
		sqlWhere += " AND method_code = ?"
		args = append(args, v)
	}
	if v, ok := params["play_code"]; ok {
		sqlWhere += " AND play_code = ?"
		args = append(args, v)
	}
	if v, ok := params["status"]; ok {
		sqlWhere += " AND status = ?"
		args = append(args, v)
	}
	if v, ok := params["flag"]; ok {
		sqlWhere += " AND flag = ?"
		args = append(args, v)
	}
	if v, ok := params["bet_date"]; ok {
		sqlWhere += " AND bet_date = ?"
		args = append(args, v)
	}
	if v, ok := params["calc_date"]; ok {
		sqlWhere += " AND calc_date = ?"
		args = append(args, v)
	}
	if v, ok := params["bet_time"]; ok {
		sqlWhere += " AND bet_time = ?"
		args = append(args, v)
	}
	if v, ok := params["update_time"]; ok {
		sqlWhere += " AND update_time = ?"
		args = append(args, v)
	}
	if v, ok := params["ip"]; ok {
		sqlWhere += " AND ip = ?"
		args = append(args, v)
	}

	if v, ok := params["amount_min"]; ok {
		sqlWhere += " AND amount >= ?"
		args = append(args, v)
	}
	if v, ok := params["amount_max"]; ok {
		sqlWhere += " AND amount <= ?"
		args = append(args, v)
	}
	if v, ok := params["payout_min"]; ok {
		sqlWhere += " AND payout >= ?"
		args = append(args, v)
	}
	if v, ok := params["payout_max"]; ok {
		sqlWhere += " AND payout <= ?"
		args = append(args, v)
	}
	if v, ok := params["profit_min"]; ok {
		sqlWhere += " AND profit >= ?"
		args = append(args, v)
	}
	if v, ok := params["profit_max"]; ok {
		sqlWhere += " AND profit <= ?"
		args = append(args, v)
	}

	if v, ok := params["bet_date_from"]; ok {
		sqlWhere += " AND bet_date >= ?"
		args = append(args, v)
	}
	if v, ok := params["bet_date_to"]; ok {
		sqlWhere += " AND bet_date <= ?"
		args = append(args, v)
	}
	if v, ok := params["calc_date_from"]; ok {
		sqlWhere += " AND calc_date >= ?"
		args = append(args, v)
	}
	if v, ok := params["calc_date_to"]; ok {
		sqlWhere += " AND calc_date <= ?"
		args = append(args, v)
	}
	if v, ok := params["bet_time_from"]; ok {
		sqlWhere += " AND bet_time >= ?"
		args = append(args, v)
	}
	if v, ok := params["bet_time_to"]; ok {
		sqlWhere += " AND bet_time <= ?"
		args = append(args, v)
	}
	if v, ok := params["update_time_from"]; ok {
		sqlWhere += " AND update_time >= ?"
		args = append(args, v)
	}

	if v, ok := params["order_by"]; ok {
		if v == "issue" {
			sqlWhere += " ORDER BY issue "
		} else if v == "amount" {
			sqlWhere += " ORDER BY amount "
		} else if v == "payout" {
			sqlWhere += " ORDER BY payout "
		} else if v == "profit" {
			sqlWhere += " ORDER BY profit "
		} else {
			sqlWhere += " ORDER BY id  "
		}
	} else {
		sqlWhere += " ORDER BY id "
	}

	if v, ok := params["sort_type"]; ok {
		if v == "ASC" {
			sqlWhere += " ASC "
		} else {
			sqlWhere += " DESC "
		}
	} else {
		sqlWhere += " DESC "
	}

	var data []Order
	var err error

	querySql += sqlWhere
	rows, err := common.BaseDb.Query(querySql, args...)
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

func PageFindLottoOrderList(pageParam common.PageParams) (*common.PageResult, *[]Order, error) {
	t := Order{}
	conditionSql := fmt.Sprintf(" FROM %s WHERE 1=1 ", t.TableName(""))
	countSql := "SELECT COUNT(*) AS count " + conditionSql
	querySql := fmt.Sprintf("SELECT %s ", strings.Join(t.Field(), ",")) + conditionSql
	sqlWhere, args := "", []interface{}{}
	if v, ok := pageParam.Params["id"]; ok {
		sqlWhere += " AND id = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["order_id"]; ok {
		sqlWhere += " AND order_id = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["user_id"]; ok {
		sqlWhere += " AND user_id = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["name"]; ok {
		sqlWhere += " AND name = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["lotto_id"]; ok {
		sqlWhere += " AND lotto_id = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["lotto_type"]; ok {
		sqlWhere += " AND lotto_type = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["game_kind"]; ok {
		sqlWhere += " AND game_kind = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["game_type"]; ok {
		sqlWhere += " AND game_type = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["issue"]; ok {
		sqlWhere += " AND issue = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["method_code"]; ok {
		sqlWhere += " AND method_code = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["play_code"]; ok {
		sqlWhere += " AND play_code = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["status"]; ok {
		sqlWhere += " AND status = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["flag"]; ok {
		sqlWhere += " AND flag = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["bet_date"]; ok {
		sqlWhere += " AND bet_date = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["calc_date"]; ok {
		sqlWhere += " AND calc_date = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["bet_time"]; ok {
		sqlWhere += " AND bet_time = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["update_time"]; ok {
		sqlWhere += " AND update_time = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["ip"]; ok {
		sqlWhere += " AND ip = ?"
		args = append(args, v)
	}

	if v, ok := pageParam.Params["amount_min"]; ok {
		sqlWhere += " AND amount >= ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["amount_max"]; ok {
		sqlWhere += " AND amount <= ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["payout_min"]; ok {
		sqlWhere += " AND payout >= ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["payout_max"]; ok {
		sqlWhere += " AND payout <= ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["profit_min"]; ok {
		sqlWhere += " AND profit >= ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["profit_max"]; ok {
		sqlWhere += " AND profit <= ?"
		args = append(args, v)
	}

	if v, ok := pageParam.Params["bet_date_from"]; ok {
		sqlWhere += " AND bet_date >= ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["bet_date_to"]; ok {
		sqlWhere += " AND bet_date <= ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["calc_date_from"]; ok {
		sqlWhere += " AND calc_date >= ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["calc_date_to"]; ok {
		sqlWhere += " AND calc_date <= ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["bet_time_from"]; ok {
		sqlWhere += " AND bet_time >= ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["bet_time_to"]; ok {
		sqlWhere += " AND bet_time <= ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["update_time_from"]; ok {
		sqlWhere += " AND update_time >= ?"
		args = append(args, v)
	}

	if v, ok := pageParam.Params["order_by"]; ok {
		if v == "issue" {
			sqlWhere += " ORDER BY issue "
		} else if v == "amount" {
			sqlWhere += " ORDER BY amount "
		} else if v == "payout" {
			sqlWhere += " ORDER BY payout "
		} else if v == "profit" {
			sqlWhere += " ORDER BY profit "
		} else {
			sqlWhere += " ORDER BY id  "
		}
	} else {
		sqlWhere += " ORDER BY id "
	}

	if v, ok := pageParam.Params["sort_type"]; ok {
		if v == "ASC" {
			sqlWhere += " ASC "
		} else {
			sqlWhere += " DESC "
		}
	} else {
		sqlWhere += " DESC "
	}

	var pg common.PageResult
	var data []Order
	var err error
	countSql += sqlWhere
	pg.PageRow = pageParam.PageRow
	pg.CurrentPage = pageParam.CurrentPage
	err = common.BaseDb.QueryRow(countSql, args...).Scan(&pg.TotalCount)
	if err != nil {
		return &pg, &data, err
	}

	if pg.TotalCount < 1 {
		return &pg, &data, err
	}

	querySql += sqlWhere
	querySql += fmt.Sprintf(" LIMIT %d,%d", (pageParam.CurrentPage-1)*pageParam.PageRow, pageParam.PageRow)
	rows, err := common.BaseDb.Query(querySql, args...)
	if err != nil {
		return &pg, &data, err
	}
	for rows.Next() {
		d := Order{}
		err = rows.Scan(d.FieldItem()...)
		if err != nil {
			return &pg, &data, err
		}
		data = append(data, d)
		pg.RecordData = append(pg.RecordData, d)
	}
	pg.PageRow = pageParam.PageRow
	pg.CurrentPage = pageParam.CurrentPage
	pg.PageCount = len(data)
	return &pg, &data, err
}

func FindLottoOrderDayCount(countDate string) (*[]LottoOrderCount, error) {
	data := []LottoOrderCount{}
	o := Order{}
	totalSql := fmt.Sprintf("SELECT user_id,name,lotto_id,lotto_type,game_kind,game_type,COALESCE(SUM(bet_count), 0) AS total_count,COALESCE(SUM(amount), 0) AS total_bet,COALESCE(SUM(payout), 0) AS total_payout,COALESCE(SUM(profit), 0) AS total_profit,calc_date AS count_date FROM %s WHERE flag = 1 AND calc_date=? GROUP BY user_id,lotto_id", o.TableName(""))
	rows, err := common.BaseDb.Query(totalSql, countDate)
	if err != nil {
		return &data, err
	}
	for rows.Next() {
		d := LottoOrderCount{}
		err = rows.Scan(d.FieldItem()...)
		if err != nil {
			return &data, err
		}
		data = append(data, d)
	}
	return &data, err
}

func SetLottoOrderInfo(id int, filed string, val interface{}) error {
	u := Order{}
	updateSql := fmt.Sprintf("UPDATE %s SET %s=? WHERE id=?", u.TableName(""), filed)
	_, err := common.BaseDb.Exec(updateSql, val, id)
	return err
}

func SettleUpdateLottoOrderTx(tx *sql.Tx, o *Order) error {
	createSql := fmt.Sprintf("UPDATE %s SET order_id=?, user_id=?, name=?, lotto_id=?, lotto_type=?, game_kind=?, game_type=?, issue=?, method_code=?, play_code=?, bet_count=?, bet_content=?, win_count=?, win_content=?, draw_number=?, odds=?, amount=?, status=?, flag=?, payout=?, profit=?, bet_date=?, calc_date=?, bet_time=?, update_time=?, ip=? WHERE id=?", o.TableName(""))
	_, err := tx.Exec(createSql, &o.OrderID, &o.UserID, &o.Name, &o.LottoID, &o.LottoType, &o.GameKind, &o.GameType, &o.Issue, &o.MethodCode, &o.PlayCode, &o.BetCount, &o.BetContent, &o.WinCount, &o.WinContent, &o.DrawNumber, &o.Odds, &o.Amount, &o.Status, &o.Flag, &o.Payout, &o.Profit, &o.BetDate, &o.CalcDate, &o.BetTime, &o.UpdateTime, &o.IP, &o.ID)
	return err
}

func FinWinningList() (*[]Winning, error) {
	t := Order{}
	d, _ := time.ParseDuration("-240h")
	day := time.Now()
	day = day.Add(d)
	bDay := common.Date2Str(day)

	querySql := fmt.Sprintf("SELECT name, lotto_id, profit FROM %s WHERE calc_date >= %s GROUP BY name ORDER BY profit DESC  LIMIT 0, 20 ", t.TableName(""), bDay)

	var data []Winning

	rows, err := common.BaseDb.Query(querySql)
	if err != nil {
		return &data, err
	}
	for rows.Next() {
		d := Winning{}
		err = rows.Scan(&d.Name, &d.LottoID, &d.Amount)
		if err != nil {
			return &data, err
		}

		d.Name = common.Mosaics(d.Name, 2, 2)
		lInfo, err := GetLotto(d.LottoID)
		if err != nil {
			d.LottoName = "彩票"
		} else {
			d.LottoName = lInfo.Name
		}

		data = append(data, d)
	}
	return &data, err
}
