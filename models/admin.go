package models

import (
	"database/sql"
	"fmt"
	"github.com/xiuos/mozi/common"
	"github.com/xiuos/xorm"
)

type AdminUser struct {
	UserID             int    `json:"user_id"`
	Name               string `json:"name"`
	Password           string `json:"password"`
	GoogleSecret       string `json:"google_secret"`
	GoogleSecretStatus int    `json:"google_secret_status"`
	Role               int    `json:"role"`
	Status             int    `json:"status"`
	CreatedAt          string `json:"created_at"`
}

func (u *AdminUser) TableName() string {
	return "admin"
}

func (u *AdminUser) Field() []string {
	return []string{"user_id", "name", "password", "google_secret", "google_secret_status", "role", "status", "created_at"}
}

func (u *AdminUser) FieldItem() []interface{} {
	return []interface{}{&u.UserID, &u.Name, &u.Password, &u.GoogleSecret, &u.GoogleSecretStatus, &u.Role, &u.Status, &u.CreatedAt}
}

type AdminUser4Admin struct {
	UserID             int    `json:"user_id"`
	Name               string `json:"name"`
	GoogleSecretStatus int    `json:"google_secret_status"`
	Role               int    `json:"role"`
	Status             int    `json:"status"`
	CreatedAt          string `json:"created_at"`
}

func CreateAdminUserTx(tx *sql.Tx, u *AdminUser) (sql.Result, error) {
	createSql := "INSERT INTO " + u.TableName() + " SET name = ?, password = ?, google_secret = ?, google_secret_status = ?, role = ?, status = ?, created_at = ?"
	rs, err := tx.Exec(createSql, &u.Name, &u.Password, &u.GoogleSecret, &u.GoogleSecretStatus, &u.Role, &u.Status, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return rs, nil
}

func CreateAdminUser(u *AdminUser) (sql.Result, error) {
	createSql := "INSERT INTO " + u.TableName() + " SET name = ?, password = ?, google_secret = ?, google_secret_status = ?, role = ?, status = ?, created_at = ?"
	rs, err := common.BaseDb.Exec(createSql, &u.Name, &u.Password, &u.GoogleSecret, &u.GoogleSecretStatus, &u.Role, &u.Status, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return rs, nil
}

func GetAdminUserByID(uid int) (*AdminUser, error) {
	o := xorm.Orm{}
	u := AdminUser{}
	querySql, err := o.Table(u.TableName()).Query().Where("user_id = ?", uid).Do(&u)
	if err != nil {
		return &u, err
	}
	err = common.BaseDb.QueryRow(querySql, o.Args()...).Scan(u.FieldItem()...)
	return &u, err
}

func SetAdminInfo(uid int, filed string, val interface{}) error {
	u := AdminUser{}
	updateSql := fmt.Sprintf("UPDATE %s SET %s=? WHERE user_id=?", u.TableName(), filed)
	_, err := common.BaseDb.Exec(updateSql, val, uid)
	return err

}

func GetAdminUserByName(name string) (*AdminUser, error) {
	o := xorm.Orm{}
	u := AdminUser{}
	querySql, err := o.Table(u.TableName()).Query().Where("name = ?", name).Do(&u)
	if err != nil {
		return &u, err
	}
	err = common.BaseDb.QueryRow(querySql, o.Args()...).Scan(u.FieldItem()...)
	return &u, err
}

func PageFindAdmin(pageParam PageParams) (*PageResult, *[]AdminUser4Admin, error) {
	conditionSql := " FROM admin WHERE 1=1 "
	countSql := "SELECT COUNT(*) AS count " + conditionSql
	querySql := "SELECT user_id,name,google_secret_status,role,status,created_at " + conditionSql
	sqlWhere, args := "", []interface{}{}
	if v, ok := pageParam.Params["user_id"]; ok {
		sqlWhere += " AND user_id = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["name"]; ok {
		sqlWhere += " AND name = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["role"]; ok {
		sqlWhere += " AND role = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["status"]; ok {
		sqlWhere += " AND status = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["start_at"]; ok {
		sqlWhere += " AND record_at >= ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["end_at"]; ok {
		sqlWhere += " AND record_at <= ?"
		args = append(args, v)
	}
	var pg PageResult
	var data []AdminUser4Admin
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
		d := AdminUser4Admin{}
		err = rows.Scan(&d.UserID, &d.Name, &d.GoogleSecretStatus, &d.Role, &d.Status, &d.CreatedAt)
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
