package errors

import "fmt"

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

func IsNameExist(err error) bool {
	_, ok := err.(NameExist)
	return ok
}

type UserAddFail struct {
}

func (err UserAddFail) Error() string {
	return fmt.Sprintf("添加用户失败")
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
