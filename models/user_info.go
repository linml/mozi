package models

import "github.com/shopspring/decimal"

type UserInfos struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	Status int    `json:"status"`

	Balance decimal.Decimal `json:"balance"`

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
}
