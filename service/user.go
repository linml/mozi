package service

import (
	"github.com/shopspring/decimal"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/models"
	"github.com/xiuos/mozi/models/errors"
)

func GetBalance(uid int) (decimal.Decimal, error) {
	uw, err := models.GetUserWallet(uid)
	return uw.Balance, err
}

func SetUserProfileEmail(uid int, email string) error {
	return models.SetUserProfile(uid, "email", email)
}

func SetUserProfileQQ(uid int, email string) error {
	return models.SetUserProfile(uid, "qq", email)
}

func SetUserProfileWeChat(uid int, weChat string) error {
	return models.SetUserProfile(uid, "wechat", weChat)
}

func SetUserProfileMobile(uid int, email string) error {
	return models.SetUserProfile(uid, "mobile", email)
}

func SetUserProfileNickname(uid int, email string) error {
	return models.SetUserProfile(uid, "nickname", email)
}

func SetUserProfileRealName(uid int, email string) error {
	return models.SetUserProfile(uid, "real_name", email)
}

func AuthLogin(name string, password string, params map[string]string) error {
	u, err := models.GetUserByName(name)

	if err != nil {
		return errors.NameOrPasswordErr{}
	}

	if u.Status == models.UserStatusDisable {
		return errors.UserStatusDisableErr{Name: u.Name}
	}

	if ok := common.CheckBCrypt(password, u.Password); ok == false {
		return errors.NameOrPasswordErr{}
	}

	ip := ""
	deviceType := 0
	userAgent := ""
	url := ""
	if val, ok := params["ip"]; ok {
		ip = val
	}
	if val, ok := params["device_type"]; ok {
		deviceType = common.GetInt(val)
	}
	if val, ok := params["user_agent"]; ok {
		userAgent = val
	}
	if val, ok := params["url"]; ok {
		url = val
	}

	r := models.RecordUserLogin{
		UserID:     u.UserID,
		Name:       u.Name,
		DeviceType: deviceType,
		IP:         ip,
		UserAgent:  userAgent,
		Url:        url,
		RecordAt:   common.GetTimeNowString(),
	}
	models.LogRecordUserLogin(&r)

	return nil
}
