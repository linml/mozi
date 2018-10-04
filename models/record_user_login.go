package models

import "github.com/xiuos/mozi/common"

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

func PageFindUserRecordLogin(pageParam PageParams) (*PageResult, *[]RecordUserLogin, error) {
	// todo
}
