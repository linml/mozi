package models

import (
	"fmt"
	"github.com/xiuos/mozi/common"
	"strings"
)

type RecordUserLogin struct {
	ID         int    `json:"id"`
	UserID     int    `json:"user_id"`
	Name       string `json:"name"`
	DeviceType int    `json:"device_type"`
	IP         string `json:"ip"`
	UserAgent  string `json:"user_agent"`
	Url        string `json:"url"`
	RecordAt   string `json:"record_at"`
	Status     int    `json:"status"`
	Remark     string `json:"remark"`
}

func (r *RecordUserLogin) TableName() string {
	return "record_user_login"
}

func (r *RecordUserLogin) Field() []string {
	return []string{"id", "user_id", "name", "device_type", "ip", "user_agent", "url", "record_at", "status", "remark"}
}

func (r *RecordUserLogin) FieldItem() []interface{} {
	return []interface{}{&r.ID, &r.UserID, &r.Name, &r.DeviceType, &r.IP, &r.UserAgent, &r.Url, &r.RecordAt, &r.Status, &r.Remark}
}

type RecordAdminLogin struct {
	ID         int    `json:"id"`
	UserID     int    `json:"user_id"`
	Name       string `json:"name"`
	DeviceType int    `json:"device_type"`
	IP         string `json:"ip"`
	UserAgent  string `json:"user_agent"`
	Url        string `json:"url"`
	RecordAt   string `json:"record_at"`
	Status     int    `json:"status"`
	Remark     string `json:"remark"`
}

func (r *RecordAdminLogin) TableName() string {
	return "record_admin_login"
}

func (r *RecordAdminLogin) Field() []string {
	return []string{"id", "user_id", "name", "device_type", "ip", "user_agent", "url", "record_at", "status", "remark"}
}

func (r *RecordAdminLogin) FieldItem() []interface{} {
	return []interface{}{&r.ID, &r.UserID, &r.Name, &r.DeviceType, &r.IP, &r.UserAgent, &r.Url, &r.RecordAt, &r.Status, &r.Remark}
}

func LogRecordUserLogin(r *RecordUserLogin) error {
	createSql := "INSERT INTO " + r.TableName() + " SET  user_id = ?, name = ?, device_type = ?, ip = ?, user_agent = ?, url = ?, record_at = ?, status = ?, remark = ?"
	_, err := common.BaseDb.Exec(createSql, r.UserID, r.Name, r.DeviceType, r.IP, r.UserAgent, r.Url, r.RecordAt, r.Status, r.Remark)
	return err
}

func LogRecordAdminLogin(r *RecordAdminLogin) error {
	createSql := "INSERT INTO " + r.TableName() + " SET  user_id = ?, name = ?, device_type = ?, ip = ?, user_agent = ?, url = ?, record_at = ?, status = ?, remark = ?"
	_, err := common.BaseDb.Exec(createSql, r.UserID, r.Name, r.DeviceType, r.IP, r.UserAgent, r.Url, r.RecordAt, r.Status, r.Remark)
	return err
}

func PageFindUserRecordLogin4Admin(pageParam common.PageParams) (*common.PageResult, *[]RecordUserLogin, error) {
	r := RecordUserLogin{}
	conditionSql := fmt.Sprintf(" FROM %s WHERE 1=1 ", r.TableName())
	countSql := "SELECT COUNT(*) AS count " + conditionSql
	querySql := fmt.Sprintf("SELECT %s ", strings.Join(r.Field(), ",")) + conditionSql
	sqlWhere, args := "", []interface{}{}
	if v, ok := pageParam.Params["user_id"]; ok {
		sqlWhere += " AND user_id = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["name"]; ok {
		sqlWhere += " AND name = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["device_type"]; ok {
		sqlWhere += " AND device_type = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["ip"]; ok {
		sqlWhere += " AND ip = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["url"]; ok {
		sqlWhere += " AND url = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["status"]; ok {
		sqlWhere += " AND status = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["start_at"]; ok {
		sqlWhere += " AND record_at >= ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["end_at"]; ok {
		sqlWhere += " AND record_at <= ?"
		args = append(args, v)
	}
	var pg common.PageResult
	var data []RecordUserLogin
	var err error
	countSql += sqlWhere
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
		r := RecordUserLogin{}
		err = rows.Scan(r.FieldItem()...)
		if err != nil {
			return &pg, &data, err
		}
		data = append(data, r)
		pg.RecordData = append(pg.RecordData, r)
	}
	pg.PageRow = pageParam.PageRow
	pg.CurrentPage = pageParam.CurrentPage
	pg.PageCount = len(data)
	return &pg, &data, err
}

func PageFindRecordAdminLogin(pageParam common.PageParams) (*common.PageResult, *[]RecordAdminLogin, error) {
	r := RecordAdminLogin{}
	conditionSql := fmt.Sprintf(" FROM %s WHERE 1=1 ", r.TableName())
	countSql := "SELECT COUNT(*) AS count " + conditionSql
	querySql := fmt.Sprintf("SELECT %s ", strings.Join(r.Field(), ",")) + conditionSql
	sqlWhere, args := "", []interface{}{}
	if v, ok := pageParam.Params["user_id"]; ok {
		sqlWhere += " AND user_id = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["name"]; ok {
		sqlWhere += " AND name = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["device_type"]; ok {
		sqlWhere += " AND device_type = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["ip"]; ok {
		sqlWhere += " AND ip = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["url"]; ok {
		sqlWhere += " AND url = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["status"]; ok {
		sqlWhere += " AND status = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["start_at"]; ok {
		sqlWhere += " AND record_at >= ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["end_at"]; ok {
		sqlWhere += " AND record_at <= ?"
		args = append(args, v)
	}
	var pg common.PageResult
	var data []RecordAdminLogin
	var err error
	countSql += sqlWhere
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
		r := RecordAdminLogin{}
		err = rows.Scan(r.FieldItem()...)
		if err != nil {
			return &pg, &data, err
		}
		data = append(data, r)
		pg.RecordData = append(pg.RecordData, r)
	}
	pg.PageRow = pageParam.PageRow
	pg.CurrentPage = pageParam.CurrentPage
	pg.PageCount = len(data)
	return &pg, &data, err
}
