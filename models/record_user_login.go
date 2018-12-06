package models

import (
	"fmt"
	"github.com/xiuos/mozi/common"
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

type RecordAdminUserLogin struct {
	ID         int
	UserID     int
	Name       string
	DeviceType int
	IP         string
	UserAgent  string
	Url        string
	RecordAt   string
	Remark     string
}

func (r *RecordAdminUserLogin) TableName() string {
	return "record_admin_user_login"
}

func LogRecordUserLogin(r *RecordUserLogin) error {
	createSql := "INSERT INTO " + r.TableName() + " SET  user_id = ?, name = ?, device_type = ?, ip = ?, user_agent = ?, url = ?, record_at = ?, status = ?, remark = ?"
	_, err := common.BaseDb.Exec(createSql, r.UserID, r.Name, r.DeviceType, r.IP, r.UserAgent, r.Url, r.RecordAt, r.Status, r.Remark)
	return err
}

func LogRecordAdminUserLogin(r *RecordAdminUserLogin) error {
	createSql := "INSERT INTO " + r.TableName() + " SET  user_id = ?, name = ?, device_type = ?, ip = ?, user_agent = ?, url = ?, record_at = ?, remark = ?"
	_, err := common.BaseDb.Exec(createSql, r.UserID, r.Name, r.DeviceType, r.IP, r.UserAgent, r.Url, r.RecordAt, r.Remark)
	return err
}

func PageFindUserRecordLogin4Admin(pageParam common.PageParams) (*common.PageResult, *[]RecordUserLogin, error) {
	conditionSql := " FROM record_user_login WHERE 1=1 "
	countSql := "SELECT COUNT(*) AS count " + conditionSql
	querySql := "SELECT id,user_id,name,device_type,ip,user_agent,url,record_at,status,remark " + conditionSql
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
		d := RecordUserLogin{}
		err = rows.Scan(&d.ID, &d.UserID, &d.Name, &d.DeviceType, &d.IP, &d.UserAgent, &d.Url, &d.RecordAt, &d.Status, &d.Remark)
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
