package service

import (
	"fmt"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/mozi/models"
	"github.com/xiuos/mozi/models/errors"
	"regexp"
	"strconv"
)

func CheckNameLegal(name string) error {
	if m, _ := regexp.MatchString("^[a-zA-Z]\\w{5,17}$", name); m {
		return nil
	}
	return errors.NameNotLegal{Name: name}
}

func CheckPasswordLegal(password string) error {
	if m, _ := regexp.MatchString("^[a-zA-Z\\d]\\w{5,17}$", password); m {
		return nil
	}
	return errors.PasswordNotLegal{}
}

func CheckWalletPasswordLegal(username string) error {
	if m, _ := regexp.MatchString("^\\d{6}$", username); m {
		return nil
	}
	return errors.WalletPasswordNotLegal{}
}

// 注册用户入口
func RegisterUser(name string, password string, params map[string]string) error {
	if err := CheckNameLegal(name); err != nil {
		return err
	}
	if err := CheckPasswordLegal(password); err != nil {
		return err
	}

	isExist, err := models.IsUserExist(name)
	if err != nil {
		return err
	} else if isExist {
		return errors.NameExist{Name: name}
	}
	pRelation := &models.UserRelation{}
	if ref, ok := params["ref"]; ok {
		refID, _ := strconv.Atoi(ref)
		pRelation, err = models.GetUserRelation(refID)
		if err != nil {
			return errors.RefNotFound{Ref: refID}
		}
	} else {
		sysSettings, err := models.GetSysSettings(models.SysDefaultParentID)
		if err != nil {
			return err
		}
		pRelation, err = models.GetUserRelation(common.GetInt(sysSettings.SysValue))
		if err != nil {
			return err
		}
	}

	encbPassword, err := common.HashPassword(password)

	if err != nil {
		return err
	}

	user := models.User{Name: name, Password: encbPassword, Status: models.UserStatusEnable}

	tx, err := common.BaseDb.Begin()

	if err != nil {
		return err
	}

	// 新增用户
	rs, err := models.CreateUserTx(tx, &user)
	if err != nil {
		tx.Rollback()
		return errors.UserAddFail{}
	}

	userID, err := rs.LastInsertId()
	if err != nil {
		return err
	}

	uid := int(userID)

	// 新增用户 end

	// 添加积分信息
	userScore := models.UserScore{UserID: uid}
	err = models.CreateUserScoreTx(tx, &userScore)
	if err != nil {
		tx.Rollback()
		common.Logger.Error(err)
		return errors.UserScoreAddFail{}
	}
	// 添加积分信息 end

	// 添加钱包
	userWallet := models.UserWallet{UserID: uid, Status: models.WalletStatusUnActivate}
	err = models.CreateUserWalletTx(tx, &userWallet)
	if err != nil {
		tx.Rollback()
		common.Logger.Error(err)
		return errors.UserWalletAddFail{}
	}
	// 添加钱包 end

	// 添加关系信息
	userRelation := models.UserRelation{
		UserID:   uid,
		ParentID: pRelation.UserID,
		Parents:  fmt.Sprintf("%s,%d", pRelation.Parents, uid),
	}
	err = models.CreateUserRelationTx(tx, &userRelation)
	if err != nil {
		tx.Rollback()
		common.Logger.Error(err)
		return errors.UserRelationAddFail{}
	}
	// 添加关系信息 end

	// profile
	registerIp := ""
	if ip, ok := params["register_ip"]; ok {
		registerIp = ip
	}
	userProfile := models.UserProfile{UserID: uid, RegisterIp: registerIp, Registered: common.GetTimeNowString()}
	err = models.CreateUserProfileTx(tx, &userProfile)
	if err != nil {
		tx.Rollback()
		return errors.UserProfileAddFail{}
	}
	// profile end

	err = tx.Commit()
	if err != nil {
		common.Logger.Error("添加用户事务失败! 详情:", err)
		return errors.RegisterUserFail{}
	}

	return nil

}
