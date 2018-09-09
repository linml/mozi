package models

import "github.com/xiuos/mozi/common"

const (
	TableRecordUserLogin = "record_user_login"
)

type RecordUserLogin struct {
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

func LogRecordUserLogin(r *RecordUserLogin) error {
	createSql := "INSERT INTO " + TableRecordUserLogin + " SET  user_id = ?, name = ?, device_type = ?, ip = ?, user_agent = ?, url = ?, record_at = ?, remark = ?"
	_, err := common.BaseDb.Exec(createSql, r.UserID, r.Name, r.DeviceType, r.IP, r.UserAgent, r.Url, r.RecordAt, r.Remark)
	return err
}
