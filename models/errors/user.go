package errors

import "fmt"

func New(text string) error {
	return &errorString{text}
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}



type NameNotLegal struct {
	Name string
}

func (err NameNotLegal) Error() string {
	return fmt.Sprintf("用户名: %s 不符合规范，用户名以字母开头,长度在6~18之间,只能包含字母、数字和下划线", err.Name)
}

type PasswordNotLegal struct {
	Name string
}

func (err PasswordNotLegal) Error() string {
	return fmt.Sprintf("密码需要6位以上的数字和字符串组合")
}

type WalletPasswordNotLegal struct {
	Name string
}

func (err WalletPasswordNotLegal) Error() string {
	return fmt.Sprintf("资金密码需要6位数字")
}

type NameExist struct {
	Name string
}

func (err NameExist) Error() string {
	return fmt.Sprintf("用户名: %s 已存在", err.Name)
}

type UserNotExist struct {
	Name string
}

func (err UserNotExist) Error() string {
	return fmt.Sprintf("用户不存在")
}

func IsNameExist(err error) bool {
	_, ok := err.(NameExist)
	return ok
}

type UserAddFail struct {
}

func (err UserAddFail) Error() string {
	return fmt.Sprintf("添加用户失败")
}

type UserDisableErr struct {
}

func (err UserDisableErr) Error() string {
	return fmt.Sprintf("账户已禁用")
}

type UserScoreAddFail struct {
}

func (err UserScoreAddFail) Error() string {
	return fmt.Sprintf("添加用户积分失败")
}

type UserWalletAddFail struct {
}

func (err UserWalletAddFail) Error() string {
	return fmt.Sprintf("添加用户钱包失败")
}

type UserRelationAddFail struct {
}

func (err UserRelationAddFail) Error() string {
	return fmt.Sprintf("添加用户关系失败")
}

type UserProfileAddFail struct {
}

func (err UserProfileAddFail) Error() string {
	return fmt.Sprintf("添加用户详情失败")
}

type RegisterUserFail struct {
}

func (err RegisterUserFail) Error() string {
	return fmt.Sprintf("注册用户事务失败")
}

type UserStatusDisableErr struct {
	Name string
}

func (err UserStatusDisableErr) Error() string {
	return fmt.Sprintf("用户: %s 处于状态不能使用", err.Name)
}

type NameOrPasswordErr struct {
	Name string
}

func (err NameOrPasswordErr) Error() string {
	return fmt.Sprintf("用户名或密码错误")
}

type PasswordErr struct {
	Name string
}

func (err PasswordErr) Error() string {
	return fmt.Sprintf("密码错误")
}

type RealNameExistErr struct {
	Name string
}

func (err RealNameExistErr) Error() string {
	return fmt.Sprintf("真实姓名已存在")
}

type Unauthorized struct {
	Name string
}

func (err Unauthorized) Error() string {
	return fmt.Sprintf("请先登录")
}

type WalletPasswordAlreadyInitErr struct {
}

func (err WalletPasswordAlreadyInitErr) Error() string {
	return fmt.Sprintf("资金密码已初始化无需重复操作")
}
