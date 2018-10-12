package models

import (
	"database/sql"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/xorm"
)

const (
	TableUser = "users"
)

const (
	UserStatusDisable      = 0 // 用户状态禁用
	UserStatusEnable       = 1 // 用户状态启用
	USER_STATUS_UNACTIVATE = 2 // 用户状态未激活
)

type User struct {
	UserID   int
	Name     string
	Password string
	Status   int
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) Field() []string {
	return []string{"user_id", "name", "password", "status"}
}

func (u *User) FieldItem() []interface{} {
	return []interface{}{&u.UserID, &u.Name, &u.Password, &u.Status}
}

// 管理后台查询用户列表
type User4Admin struct {
	UserID      int             `json:"user_id"`
	Name        string          `json:"name"`
	Balance     decimal.Decimal `json:"balance"`
	Status      int             `json:"status"`
	Registered  string          `json:"registered"`
	LastLoginAt string          `json:"last_login_at"`
	LastLoginIP string          `json:"last_login_ip"`
	UserType    int             `json:"user_type"`
}

func CreateUserTx(tx *sql.Tx, u *User) (sql.Result, error) {
	createSql := "INSERT INTO " + u.TableName() + " SET name = ?, password = ?, status = ?"
	rs, err := tx.Exec(createSql, u.Name, &u.Password, &u.Status)
	if err != nil {
		return nil, err
	}
	return rs, nil
}

func CreateUser(u *User) error {
	createSql := "INSERT INTO " + u.TableName() + " SET name = ?, password = ?, status = ?"
	_, err := common.BaseDb.Exec(createSql, u.Name, &u.Password, &u.Status)
	return err
}

func GetUserByID(uid int) (*User, error) {
	o := xorm.Orm{}
	u := User{}
	querySql, err := o.Table(u.TableName()).Query().Where("user_id = ?", uid).Do(&u)
	if err != nil {
		return &u, err
	}
	err = common.BaseDb.QueryRow(querySql, o.Args()...).Scan(u.FieldItem()...)
	return &u, err
}

func GetUserByName(name string) (*User, error) {
	o := xorm.Orm{}
	u := User{}
	querySql, err := o.Table(u.TableName()).Query().Where("name = ?", name).Do(&u)
	if err != nil {
		return &u, err
	}
	err = common.BaseDb.QueryRow(querySql, o.Args()...).Scan(u.FieldItem()...)
	return &u, err
}

func IsUserExist(name string) (bool, error) {
	querySql := fmt.Sprintf("SELECT COUNT(*) as count FROM %s Where name = ?", TableUser)
	var c int
	err := common.BaseDb.QueryRow(querySql, name).Scan(&c)
	if err != nil {
		return false, err
	}
	if c > 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func SetPassword(uid int, password string) error {
	hashedPassword, err := common.HashPassword(password)

	if err != nil {
		return err
	}
	updateSql := fmt.Sprintf("UPDATE %s SET password = ? WHERE user_id = ?", TableUser)
	_, err = common.BaseDb.Exec(updateSql, hashedPassword, uid)
	return err
}

func PageFindUserList(pageParam PageParams) (*PageResult, *[]User4Admin, error) {
	conditionSql := " FROM users a LEFT JOIN user_wallet b ON a.user_id=b.user_id LEFT JOIN user_profile c ON a.user_id=c.user_id LEFT JOIN user_relation d ON a.user_id=d.user_id WHERE 1=1 "
	countSql := "SELECT COUNT(*) AS count " + conditionSql
	querySql := "SELECT a.user_id,a.name,a.status,b.balance,c.registered,c.last_login_at,c.last_login_ip,d.user_type " + conditionSql
	sqlWhere, args := "", []interface{}{}
	if v, ok := pageParam.Params["user_id"]; ok {
		sqlWhere += " AND a.user_id = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["name"]; ok {
		sqlWhere += " AND a.name = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["status"]; ok {
		sqlWhere += " AND a.status = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["parent_id"]; ok {
		sqlWhere += " AND d.parent_id = ?"
		args = append(args, v)
	}
	var pg PageResult
	var data []User4Admin
	var err error
	countSql += sqlWhere
	err = common.BaseDb.QueryRow(countSql, args...).Scan(&pg.TotalCount)
	if err != nil {
		return &pg, &data, err
	}

	if pg.TotalCount < 1 {
		return &pg, &data, err
	}

	querySql += sqlWhere
	querySql += fmt.Sprintf(" LIMIT %d,%d", (pageParam.CurrentPage-1)*pageParam.PageRow, pageParam.PageRow)
	rows, err := common.BaseDb.Query(querySql, args...)
	if err != nil {
		return &pg, &data, err
	}
	for rows.Next() {
		d := User4Admin{}
		err = rows.Scan(&d.UserID, &d.Name, &d.Status, &d.Balance, &d.Registered, &d.LastLoginAt, &d.LastLoginIP, &d.UserType)
		if err != nil {
			return &pg, &data, err
		}
		data = append(data, d)
		pg.RecordData = append(pg.RecordData, d)
	}
	pg.PageRow = pageParam.PageRow
	pg.CurrentPage = pageParam.CurrentPage
	pg.PageCount = len(data)
	return &pg, &data, err
}
