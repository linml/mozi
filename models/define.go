package models

const (
	OperatorTypeSelf  = 0
	OperatorTypeOther = 1
	OperatorTypeAdmin = 2
)

const (
	LottoStatusDisable = 0 // 关闭
	LottoStatusEnable  = 1 // 正常
	LottoStatusPause   = 2 // 停售
)

const (
	ChangeKindLotto    = 1 // 彩票
	ChangeKindDeposit  = 2 // 入款
	ChangeKindWithdraw = 3 // 出款
	ChangeKindTransfer = 4 // 转账
	ChangeKindActivity = 5 // 活动
	ChangeKindWage     = 6 // 工资
	ChangeKindDividend = 7 // 分红
)

const (
	ChangeTypeLottoBet    = 101 // 彩票-下注
	ChangeTypeLottoPayout = 102 // 彩票-派彩
	ChangeTypeLottoCancel = 103 // 彩票-撤单

	ChangeKindDepositManual   = 200 //入款-人工入款
	ChangeKindDepositCencel   = 201 //入款-冲账-取消出款
	ChangeKindDepositRepeate  = 202 //入款-冲账-重复出款
	ChangeKindDepositPayoff   = 203 //入款-优惠(活动)
	ChangeKindDepositActivity = 204 //入款-活动(活动)
	ChangeKindDepositOther    = 205 //入款-其他(活动)
)
