package service

import (
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/models"
	"github.com/xiuos/mozi/models/errors"
	"fmt"
)

func GetAdminUserIDByName(name string) (int, error) {
	u, err := models.GetAdminUserByName(name)
	if err != nil {
		return 0, err
	}
	return u.UserID, nil
}

func AuthAdminLogin(name string, password string, params map[string]string) error {
	u, err := models.GetAdminUserByName(name)

	if err != nil {
		fmt.Println(err)
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

	r := models.RecordAdminUserLogin{
		UserID:     u.UserID,
		Name:       u.Name,
		DeviceType: deviceType,
		IP:         ip,
		UserAgent:  userAgent,
		Url:        url,
		RecordAt:   common.GetTimeNowString(),
	}
	models.LogRecordAdminUserLogin(&r)

	return nil
}
