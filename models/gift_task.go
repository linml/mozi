package models

import (
	"fmt"
	"github.com/xiuos/mozi/common"
	"strings"
)

type GiftTask struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImgUrl      string `json:"img_url"`
	Content     string `json:"content"`
	SortIndex   int    `json:"sort_index"`
	Status      int    `json:"status"`
	ShowBanner  int    `json:"show_banner"`
	StartAt     string `json:"start_at"`
	EndAt       string `json:"end_at"`
}

func (g *GiftTask) TableName() string {
	return "gift_task"
}

func (g *GiftTask) Field() []string {
	return []string{"id", "title", "description", "img_url", "content", "sort_index", "status", "show_banner", "start_at", "end_at"}
}

func (g *GiftTask) FieldItem() []interface{} {
	return []interface{}{&g.ID, &g.Title, &g.Description, &g.ImgUrl, &g.Content, &g.SortIndex, &g.Status, &g.ShowBanner, &g.StartAt, &g.EndAt}
}

func FindGiftTaskList(param map[string]string) (*[]GiftTask, error) {
	arl := []GiftTask{}
	t := GiftTask{}
	querySql := fmt.Sprintf("SELECT %s FROM %s WHERE 1=1", strings.Join(t.Field(), ","), t.TableName())
	sqlWhere, args := "", []interface{}{}
	if v, ok := param["id"]; ok {
		sqlWhere += " AND id = ?"
		args = append(args, v)
	}
	if v, ok := param["status"]; ok {
		sqlWhere += " AND status = ?"
		args = append(args, v)
	}
	if v, ok := param["show_banner"]; ok {
		sqlWhere += " AND show_banner = ?"
		args = append(args, v)
	}
	if v, ok := param["is_show"]; ok {
		sqlWhere += " AND is_show = ?"
		args = append(args, v)
	}

	if v, ok := param["start_at_from"]; ok {
		sqlWhere += " AND start_at >= ?"
		args = append(args, v)
	}
	if v, ok := param["end_at_from"]; ok {
		sqlWhere += " AND end_at <= ?"
		args = append(args, v)
	}

	if v, ok := param["order_by"]; ok {
		f, val := true, ""
		switch v {
		case "sort_index":
			val = " ORDER BY sort_index "
		default:
			f = false
		}
		sqlWhere += val
		if f == true {
			if v, ok := param["sort_type"]; ok {
				if v == "ASC" {
					sqlWhere += " ASC "
				} else {
					sqlWhere += " DESC "
				}
			}
		}
	}

	querySql += sqlWhere
	rows, err := common.BaseDb.Query(querySql, args...)
	if err != nil {
		fmt.Println(err)
		return &arl, err
	}
	for rows.Next() {
		am := GiftTask{}
		err = rows.Scan(am.FieldItem()...)
		if err != nil {
			return &arl, err
		}
		arl = append(arl, am)
	}
	return &arl, nil
}
