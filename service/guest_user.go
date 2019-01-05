package service

import (
	"database/sql"
	"github.com/shopspring/decimal"
	"github.com/xiuos/mozi/models"
	"github.com/xiuos/mozi/models/errors"
)

func GetGuestBalance(uid int) (decimal.Decimal, error) {
	uw, err := models.GetGuestUser(uid)
	return uw.Balance, err
}

func GetGuestInfos(uid int) (*models.UserInfos, error) {
	var ui models.UserInfos
	u, err := models.GetGuestUser(uid)
	if err != nil {
		if err == sql.ErrNoRows {
			return &ui, errors.New("账号不存在")
		}
		return &ui, err
	}

	ui.UserID = u.UserID
	ui.Name = u.Name
	ui.Status = 1
	ui.Balance = u.Balance
	ui.RealName = ""
	ui.Nickname = ""
	ui.Email = ""
	ui.IsEmailVerified = 0
	ui.Mobile = ""
	ui.IsMobileVerified = 0
	ui.QQ = ""
	ui.Wechat = ""
	ui.RegisterIp = ""
	ui.Registered = u.CreateAt
	return &ui, err
}
