package models

import (
	"database/sql"
	"fmt"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/xorm"
)

const (
	TableUserProfile = "user_profile"
)

type UserProfile struct {
	UserID           int
	RealName         string
	Nickname         string
	Email            string
	IsEmailVerified  int
	Mobile           string
	IsMobileVerified int
	QQ               string
	Wechat           string
	RegisterIp       string
	Registered       string
}

func (up *UserProfile) Field() []string {
	return []string{"user_id", "real_name", "nickname", "email", "is_email_verified", "mobile", "is_mobile_verified", "qq", "wechat", "register_ip", "registered"}
}

func (up *UserProfile) FieldItem() []interface{} {
	return []interface{}{&up.UserID, &up.RealName, &up.Nickname, &up.Email, &up.IsEmailVerified, &up.Mobile, &up.IsMobileVerified, &up.QQ, &up.Wechat, &up.RegisterIp, &up.Registered}
}

func CreateUserProfileTx(tx *sql.Tx, up *UserProfile) error {
	o := xorm.Orm{}
	createSql, err := o.Table(TableUserProfile).Create().Do(up)
	if err != nil {
		return err
	}
	_, err = tx.Exec(createSql, up.FieldItem()...)
	return err
}

func GetUserProfile(uid int) (*UserProfile, error) {
	o := xorm.Orm{}
	up := UserProfile{}
	querySql, err := o.Table(TableUserProfile).Query().Where("user_id = ?", uid).Do(&up)
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
