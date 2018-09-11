package service

import (
	"github.com/xiuos/mozi/models"
	"github.com/xiuos/mozi/models/errors"
	"github.com/xiuos/mozi/common"
)

func Bet(o *models.Order) error {
	u, err := models.GetUserByID(o.UserID)
	if err != nil {
		return err
	}

	if u.Status == models.UserStatusDisable{
		return errors.UserDisableErr{}
	}

	uw,err:= models.GetUserWallet(o.UserID)
	if err != nil {
		return err
	}

	if o.Amount.GreaterThan(uw.Balance){
		return errors.New("账户余额不足")
	}

	err = models.MethodCheckBetLegal(o.MethodCode, o)
	if err != nil {
		return err
	}
	o.BetTime = common.GetTimeNowString()
	o.BetDate = common.GetDateNowString()


}
