package models

import "github.com/xiuos/mozi/common"

const (
	ActionModuleUser     = 1 // 用户模块
	ActionModuleMoney    = 2 // 资金模块
	ActionModuleActivity = 3 // 活动模块
	ActionModuleSystem   = 4 // 系统模块
)

const (
	ActionUserRegister = 1 // 用户注册
	ActionUserLogin    = 2 // 用户登入
	ActionCreateRef    = 3 // 生成推广码
)

type RecordUserAction struct {
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

func (r *RecordUserAction) TableName() string {
	return "record_user_action"
}

func (r *RecordUserAction) Field() []string {
	return []string{"id", "user_id", "action_module", "action_id", "content", "ip", "record_at", "success", "message", "operator_id", "operator_type"}
}

func (r *RecordUserAction) FieldItem() []interface{} {
	return []interface{}{&r.ID, &r.UserID, &r.ActionModule, &r.ActionID, &r.Content, &r.IP, &r.RecordAt, &r.Success, &r.Message, &r.OperatorID, &r.OperatorType}
}

func LogRecordUserAction(r *RecordUserAction) error {
	createSql := "INSERT INTO " + r.TableName() + " SET user_id = ?, action_module = ?, action_id = ?, content = ?, ip = ?, record_at = ?, success = ?, message = ?, operator_id = ?, operator_type = ?"
	_, err := common.BaseDb.Exec(createSql, &r.UserID, &r.ActionModule, &r.ActionID, &r.Content, &r.IP, &r.RecordAt, &r.Success, &r.Message, &r.OperatorID, &r.OperatorType)
	return err
}
