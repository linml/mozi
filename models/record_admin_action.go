package models

import (
	"fmt"
	"github.com/xiuos/mozi/common"
)

type RecordAdminAction struct {
	ID           int    `json:"id"`
	UserID       int    `json:"user_id"`
	Name         string `json:"name"`
	ActionModule int    `json:"action_module"`
	ActionID     int    `json:"action_id"`
	Content      string `json:"content"`
	IP           string `json:"ip"`
	RecordAt     string `json:"record_at"`
	Success      int    `json:"success"`
	Message      string `json:"message"`
	OperatorID   int    `json:"operator_id"`
	OperatorName string `json:"operator_name"`
	OperatorType int    `json:"operator_type"`
}

func (r *RecordAdminAction) TableName() string {
	return "record_log_admin_action"
}

func (r *RecordAdminAction) Field() []string {
	return []string{"id", "user_id", "name", "action_module", "action_id", "content", "ip", "record_at", "success", "message", "operator_id", "operator_name", "operator_type"}
}

func (r *RecordAdminAction) FieldItem() []interface{} {
	return []interface{}{&r.ID, &r.UserID, &r.Name, &r.ActionModule, &r.ActionID, &r.Content, &r.IP, &r.RecordAt, &r.Success, &r.Message, &r.OperatorID, &r.OperatorName, &r.OperatorType}
}

func LogRecordAdminAction(r *RecordAdminAction) error {
	createSql := "INSERT INTO " + r.TableName() + " SET user_id = ?, name = ?, action_module = ?, action_id = ?, content = ?, ip = ?, record_at = ?, success = ?, message = ?, operator_id = ?, operator_name = ?, operator_type = ?"
	_, err := common.BaseDb.Exec(createSql, &r.UserID, &r.Name, &r.ActionModule, &r.ActionID, &r.Content, &r.IP, &r.RecordAt, &r.Success, &r.Message, &r.OperatorID, &r.OperatorName, &r.OperatorType)
	return err
}

func PageFindLogRecordAdminActionList(pageParam common.PageParams) (*common.PageResult, *[]RecordAdminAction, error) {
	r := RecordAdminAction{}
	conditionSql := fmt.Sprintf(" FROM %s WHERE 1=1 ", r.TableName())
	countSql := "SELECT COUNT(*) AS count " + conditionSql
	querySql := "SELECT id,user_id,name,action_module,action_id,content,ip,record_at,success,message,operator_id,operator_name,operator_type " + conditionSql
	sqlWhere, args := "", []interface{}{}
	if v, ok := pageParam.Params["user_id"]; ok {
		sqlWhere += " AND user_id = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["name"]; ok {
		sqlWhere += " AND name = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["action_module"]; ok {
		sqlWhere += " AND action_module = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["action_id"]; ok {
		sqlWhere += " AND action_id = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["ip"]; ok {
		sqlWhere += " AND ip = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["success"]; ok {
		sqlWhere += " AND success = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["operator_id"]; ok {
		sqlWhere += " AND operator_id = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["operator_name"]; ok {
		sqlWhere += " AND operator_name = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["operator_type"]; ok {
		sqlWhere += " AND operator_type = ?"
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
	sqlWhere += " ORDER BY record_at DESC "

	var pg common.PageResult
	var data []RecordAdminAction
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
		d := RecordAdminAction{}
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
