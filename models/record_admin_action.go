package models

import "github.com/xiuos/mozi/common"

const (
	ActionAdminModuleUser     = 1 // 用户模块
	ActionAdminModuleMoney    = 2 // 资金模块
	ActionAdminModuleActivity = 3 // 活动模块
	ActionAdminModuleSystem   = 4 // 系统模块
)

const (
	ActionAdminAddRole     = 1 // 添加角色
	ActionAdminDeleteRole  = 2 // 删除角色
	ActionAdminUpdateeRole = 3 // 设置角色

	ActionAdminSetRoleMenu = 4 // 设置角色菜单
)

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
