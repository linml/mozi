package models

const (
	UserTypeProxy = 0 // 代理
	UserTypePlay  = 1 // 会员
)

const (
	PowerEnable  = 1
	PowerDisable = 0
)
const (
	OperatorTypeSelf  = 0
	OperatorTypeOther = 1
	OperatorTypeAdmin = 2
)

const (
	OperatorTypeAdminSelf   = 0
	OperatorTypeAdminSystem = 1
)

const (
	LottoStatusDisable = 0 // 关闭
	LottoStatusEnable  = 1 // 正常
	LottoStatusPause   = 2 // 停售
)
