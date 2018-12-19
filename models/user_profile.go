package models

import (
	"database/sql"
	"fmt"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/xorm"
)

const (
	TableUserProfile = "user_profile"
)

type UserProfile struct {
	UserID           int    `json:"user_id"`
	RealName         string `json:"real_name"`
	Nickname         string `json:"nickname"`
	Email            string `json:"email"`
	IsEmailVerified  int    `json:"is_email_verified"`
	Mobile           string `json:"mobile"`
	IsMobileVerified int    `json:"is_mobile_verified"`
	QQ               string `json:"qq"`
	Wechat           string `json:"wechat"`
	RegisterIp       string `json:"register_ip"`
	Registered       string `json:"registered"`
	LastLoginAt      string `json:"last_login_at"`
	LastLoginIp      string `json:"last_login_ip"`
}

func (up *UserProfile) TableName() string {
	return "user_profile"
}

func (up *UserProfile) Field() []string {
	return []string{"user_id", "real_name", "nickname", "email", "is_email_verified", "mobile", "is_mobile_verified", "qq", "wechat", "register_ip", "registered", "last_login_at", "last_login_ip"}
}

func (up *UserProfile) FieldItem() []interface{} {
	return []interface{}{&up.UserID, &up.RealName, &up.Nickname, &up.Email, &up.IsEmailVerified, &up.Mobile, &up.IsMobileVerified, &up.QQ, &up.Wechat, &up.RegisterIp, &up.Registered, &up.LastLoginAt, &up.LastLoginIp}
}

func CreateUserProfileTx(tx *sql.Tx, up *UserProfile) error {
	o := xorm.Orm{}
	createSql, err := o.Table(up.TableName()).Create().Do(up)
	if err != nil {
		return err
	}
	_, err = tx.Exec(createSql, up.FieldItem()...)
	return err
}

func GetUserProfile(uid int) (*UserProfile, error) {
	o := xorm.Orm{}
	up := UserProfile{}
	querySql, err := o.Table(up.TableName()).Query().Where("user_id = ?", uid).Do(&up)
	if err != nil {
		return &up, err
	}
	err = common.BaseDb.QueryRow(querySql, o.Args()...).Scan(up.FieldItem()...)
	return &up, err
}

func SetUserProfile(uid int, filed string, value interface{}) error {
	updateSql := fmt.Sprintf("UPDATE %s SET %s = ? WHERE user_id = ?", TableUserProfile, filed)
	_, err := common.BaseDb.Exec(updateSql, value, uid)
	if err != nil {
		return err
	}
	return nil
}
