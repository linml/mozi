package service

import (
	"github.com/shopspring/decimal"
	"github.com/xiuos/mozi/models"
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