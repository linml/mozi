package service

import (
	"fmt"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/models"
	"github.com/xiuos/mozi/models/errors"
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

func UpdateRoleMenu(roleID int, newMenu []int) error {
	oldRoleList, err := models.FindRoleMenu(map[string]string{"role_id": fmt.Sprintf("%d", roleID)})
	if err != nil {
		return err
	}
	for i, _ := range *oldRoleList {
		r := models.AdminRoleMenu{
			RoleID: roleID,
			MenuID: (*oldRoleList)[i].MenuID,
		}
		models.DelRoleMenu(&r)
	}

	for _, mID := range newMenu {
		r := models.AdminRoleMenu{
			RoleID: roleID,
			MenuID: mID,
		}
		models.AddRoleMenu(&r)
	}
	return nil
}

func SetAdminRole(uid int, roleID int) error {
	return models.SetAdminInfo(uid, "role", roleID)
}

func SetAdminStatus(uid int, status int) error {
	return models.SetAdminInfo(uid, "status", status)
}

func AddAdmin(name string, password string, role int) error {
	hashPassword, err := common.HashPassword(password)
	if err != nil {
		return err
	}
	u := models.AdminUser{
		Name:               name,
		Password:           hashPassword,
		Role:               role,
		Status:             1,
		CreatedAt:          common.GetTimeNowString(),
		GoogleSecretStatus: 0,
	}
	_, err = models.CreateAdminUser(&u)
	return err
}
