package service

import (
	"github.com/shopspring/decimal"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/models"
	"github.com/xiuos/mozi/models/errors"
)

func GetUserIDByName(name string) (int, error) {
	u, err := models.GetUserByName(name)
	if err != nil {
		return 0, err
	}
	return u.UserID, nil
}

func GetNameByUserID(uid int) (string, error) {
	u, err := models.GetUserByID(uid)
	if err != nil {
		return "", err
	}
	return u.Name, nil
}

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

func GetInfos(uid int) (*models.UserInfos, error) {
	var ui models.UserInfos
	u, err := models.GetUserByID(uid)
	if err != nil {
		return &ui, err
	}
	up, err := models.GetUserProfile(uid)
	if err != nil {
		return &ui, err
	}

	uw, err := models.GetUserWallet(uid)
	if err != nil {
		return &ui, err
	}
	ui.UserID = u.UserID
	ui.Name = u.Name
	ui.Status = u.Status
	ui.Balance = uw.Balance
	ui.RealName = up.RealName
	ui.Nickname = up.Nickname
	ui.Email = up.Email
	ui.IsEmailVerified = up.IsEmailVerified
	ui.Mobile = up.Mobile
	ui.IsMobileVerified = up.IsMobileVerified
	ui.QQ = up.QQ
	ui.Wechat = up.Wechat
	ui.RegisterIp = up.RegisterIp
	ui.Registered = up.Registered
	return &ui, err
}

// 重置密码，区分admin/user不同操作者
func ResetUserPassword(uid int, oldPass string, newPass string, operatorType int) error {
	if err := CheckPasswordLegal(newPass); err != nil {
		return err
	}

	if operatorType != models.OperatorTypeAdmin {
		u, err := models.GetUserByID(uid)

		if err != nil {
			return errors.UserNotExist{}
		}

		if ok := common.CheckBCrypt(oldPass, u.Password); ok == false {
			return errors.PasswordErr{}
		}
	}

	err := models.SetPassword(uid, newPass)
	return err

}

func InitUserWalletPassword(uid int, password string) error {
	err := CheckWalletPasswordLegal(password)
	if err != nil {
		return err
	}

	uw, err := models.GetUserWallet(uid)

	if uw.Status == models.WalletStatusUnActivate {
		models.SetWalletPassword(uid, password)
		models.SetWalletStatus(uid, models.WalletStatusEnable)

	} else {
		return errors.WalletPasswordAlreadyInitErr{}
	}

	return nil
}

func RestUserWalletPassword(uid int, oldPass string, newPass string, operatorType int) error {

	err := CheckWalletPasswordLegal(newPass)
	if err != nil {
		return err
	}

	if operatorType != models.OperatorTypeAdmin {
		uw, err := models.GetUserWallet(uid)
		if err != nil {
			return err
		}
		if b := common.CheckBCrypt(oldPass, uw.Password); b == false {
			return errors.PasswordErr{}
		}
	}

	err = models.SetWalletPassword(uid, newPass)
	return err

}

func GetUserWalletPasswordStatus(uid int) (int, error) {
	uw, err := models.GetUserWallet(uid)
	if err != nil {
		return 0, err
	}
	return uw.Status, nil

}

func GenUserRef(refInfo models.UserLinks) error {
	uid := refInfo.UserID
	uRelation, err := models.GetUserRelation(uid)
	if err != nil {
		return err
	}

	if uRelation.UserType == models.UserTypePlay {
		return errors.New("暂无权限")
	}

	tx, err := common.BaseDb.Begin()

	if err != nil {
		return err
	}

	refInfo.Ref = common.RandLowerString(6)
	err = models.CreateUserLinksTx(tx, &refInfo)
	if err != nil {
		tx.Rollback()
		return errors.New("添加连接失败")
	}
	err = tx.Commit()
	return err
}
