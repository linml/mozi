package models

import (
	"fmt"
	"github.com/xiuos/mozi/common"
	"strings"
)

const (
	NoticeClassRoll = 1 // 滚动公告
	NoticeClassPop  = 2 // 弹窗公告
)

type Notice struct {
	ID        int    `json:"id"`
	ClassID   int    `json:"class_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	SortIndex int    `json:"sort_index"`
	Status    int    `json:"status"`
	StartAt   string `json:"start_at"`
	EndAt     string `json:"end_at"`
	UpdateAt  string `json:"update_at"`
}

func (n *Notice) TableName() string {
	return "notice"
}

func (n *Notice) Field() []string {
	return []string{"id", "class_id", "title", "content", "sort_index", "status", "start_at", "end_at", "update_at"}
}

func (n *Notice) FieldItem() []interface{} {
	return []interface{}{&n.ID, &n.ClassID, &n.Title, &n.Content, &n.SortIndex, &n.Status, &n.StartAt, &n.EndAt, &n.UpdateAt}
}

func AddNotice(n *Notice) error {
	createSql := fmt.Sprintf("INSERT INTO %s SET class_id=?, title=?, content=?, sort_index=?, status=?, start_at=?, end_at=?, update_at=?", n.TableName())
	_, err := common.BaseDb.Exec(createSql, n.ClassID, n.Title, n.Content, n.SortIndex, n.Status, n.StartAt, n.EndAt, n.UpdateAt)
	return err
}

func SetNotice(n *Notice) error {
	updateSql := fmt.Sprintf("UPDATE %s SET class_id=?, title=?, content=?, sort_index=?, status=?, start_at=?, end_at=?, update_at=? WHERE id=?", n.TableName())
	_, err := common.BaseDb.Exec(updateSql, n.ClassID, n.Title, n.Content, n.SortIndex, n.Status, n.StartAt, n.EndAt, n.UpdateAt, n.ID)
	return err
}

func DelNotice(id int) error {
	n := Notice{}
	createSql := fmt.Sprintf("DELETE FROM %s WHERE id=?", n.TableName())
	_, err := common.BaseDb.Exec(createSql, id)
	return err
}

func GetNotice(id int) (*Notice, error) {
	n := Notice{}
	querySql := fmt.Sprintf("SELECT %s FROM %s WHERE id=?", strings.Join(n.Field(), ","), n.TableName())
	err := common.BaseDb.QueryRow(querySql, id).Scan(n.FieldItem()...)
	return &n, err
}

func FindNoticeList(param map[string]string) (*[]Notice, error) {
	n := Notice{}
	querySql := fmt.Sprintf("SELECT %s FROM %s WHERE 1=1", strings.Join(n.Field(), ","), n.TableName())
	sqlWhere, args := "", []interface{}{}
	if v, ok := param["id"]; ok {
		sqlWhere += " AND id = ?"
		args = append(args, v)
	}
	if v, ok := param["class_id"]; ok {
		sqlWhere += " AND class_id = ?"
		args = append(args, v)
	}
	if v, ok := param["status"]; ok {
		sqlWhere += " AND status = ?"
		args = append(args, v)
	}

	querySql += sqlWhere
	querySql += "ORDER BY sort_id "

	nl := []Notice{}
	rows, err := common.BaseDb.Query(querySql, args...)
	if err != nil {
		return &nl, err
	}
	for rows.Next() {
		n := Notice{}
		err = rows.Scan(n.FieldItem()...)
		if err != nil {
			return &nl, err
		}
		nl = append(nl, n)
	}
	return &nl, nil
}

func PageFindNoticeList(pageParam common.PageParams) (*common.PageResult, *[]Notice, error) {
	t := Notice{}
	conditionSql := fmt.Sprintf(" FROM %s WHERE 1=1 ", t.TableName())
	countSql := "SELECT COUNT(*) AS count " + conditionSql
	querySql := fmt.Sprintf("SELECT %s ", strings.Join(t.Field(), ",")) + conditionSql
	sqlWhere, args := "", []interface{}{}
	if v, ok := pageParam.Params["id"]; ok {
		sqlWhere += " AND id = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["class_id"]; ok {
		sqlWhere += " AND class_id = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["status"]; ok {
		sqlWhere += " AND status = ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["start_at_from"]; ok {
		sqlWhere += " AND start_at >= ?"
		args = append(args, v)
	}
	if v, ok := pageParam.Params["end_at_to"]; ok {
		sqlWhere += " AND end_at <= ?"
		args = append(args, v)
	}

	var pg common.PageResult
	var data []Notice
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
		d := Notice{}
		err = rows.Scan(d.FieldItem()...)
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
