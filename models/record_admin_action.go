package models

import "github.com/xiuos/mozi/common"

type RecordAdminAction struct {
	ID           int
	UserID       int
	ActionModule int
	ActionID     int
	Content      string
	IP           string
	RecordAt     string
	Success      int
	Message      string
	OperatorID   int
	OperatorType int
}

func (r *RecordAdminAction) TableName() string {
	return "record_log_admin_action"
}

func (r *RecordAdminAction) Field() []string {
	return []string{"id", "user_id", "action_module", "action_id", "content", "ip", "record_at", "success", "message", "operator_id", "operator_type"}
}

func (r *RecordAdminAction) FieldItem() []interface{} {
	return []interface{}{&r.ID, &r.UserID, &r.ActionModule, &r.ActionID, &r.Content, &r.IP, &r.RecordAt, &r.Success, &r.Message, &r.OperatorID, &r.OperatorType}
}

func LogRecordAdminAction(r *RecordAdminAction) error {
	createSql := "INSERT INTO " + r.TableName() + " SET user_id = ?, action_module = ?, action_id = ?, content = ?, ip = ?, record_at = ?, success = ?, message = ?, operator_id = ?, operator_type = ?"
	_, err := common.BaseDb.Exec(createSql, &r.UserID, &r.ActionModule, &r.ActionID, &r.Content, &r.IP, &r.RecordAt, &r.Success, &r.Message, &r.OperatorID, &r.OperatorType)
	return err
}
