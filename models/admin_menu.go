package models

import (
	"fmt"
	"github.com/xiuos/mozi/common"
	"strconv"
	"strings"
)

type AdminMenu struct {
	ID         int    `json:"id"`
	Sort       int    `json:"sort"`
	PID        int    `json:"pid"`
	Open       int    `json:"open"`
	Text       string `json:"text"`
	Icon       string `json:"icon"`
	Url        string `json:"url"`
	TargetType string `json:"target_type"`
}

func (am *AdminMenu) TableName() string {
	return "admin_menu"
}

func (am *AdminMenu) Field() []string {
	return []string{"id", "sort", "pid", "open", "text", "icon", "url", "target_type"}
}

func (am *AdminMenu) FieldItem() []interface{} {
	return []interface{}{&am.ID, &am.Sort, &am.PID, &am.Open, &am.Text, &am.Icon, &am.Url, &am.TargetType}
}

type TreePowerMap struct {
	ID         int            `json:"id"`
	Sort       int            `json:"sort"`
	PID        int            `json:"pid"`
	IsOpen     int            `json:"isOpen"`
	Text       string         `json:"text"`
	Icon       string         `json:"icon"`
	Url        string         `json:"url"`
	TargetType string         `json:"targetType"`
	Child      []TreePowerMap `json:"children"`
}

func S2IList(l []string) ([]int, error) {
	_l := make([]int, len(l))

	for i, v := range l {
		b, err := strconv.Atoi(v)
		if err != nil {
			return _l, err
		} else {
			_l[i] = b
		}
	}
	return _l, nil
}

func CheckInInt(dis []int, k int) bool {
	for i, _ := range dis {
		if dis[i] == k {
			return true
		}
	}
	return false
}

func GetTreeMenu(mp *TreePowerMap, pid int, dataList *[]AdminMenu) {
	for _, v := range *dataList {
		if pid == v.PID {
			d := TreePowerMap{ID: v.ID, Sort: v.Sort, PID: v.PID, IsOpen: v.Open, Text: v.Text, Icon: v.Icon, Url: v.Url, TargetType: v.TargetType, Child: []TreePowerMap{}}
			GetTreeMenu(&d, v.ID, dataList)
			mp.Child = append(mp.Child, d)
		}
	}
}

func FindAdminMenu(params map[string]string) (*[]AdminMenu, error) {
	am := AdminMenu{}
	aml := []AdminMenu{}
	querySql := fmt.Sprintf("SELECT %s FROM %s WHERE 1=1 ", strings.Join(am.Field(), ","), am.TableName())
	rows, err := common.BaseDb.Query(querySql)
	if err != nil {
		return &aml, err
	}

	for rows.Next() {
		am := AdminMenu{}
		err = rows.Scan(am.FieldItem()...)
		if err != nil {
			return &aml, err
		}
		aml = append(aml, am)
	}
	return &aml, err
}

func GetAdminMenuTree(uid int) (*TreePowerMap, error) {
	t := TreePowerMap{}
	u, err := GetAdminUserByID(uid)
	if err != nil {
		return &t, err
	}
	roleInfo, err := GetAdminRole(u.Role)
	mIdList := strings.Split(roleInfo.Menu, ",")
	idList, _ := S2IList(mIdList)
	m, err := FindAdminMenu(map[string]string{})
	if err != nil {
		fmt.Println(err)
		return &t, err
	}
	var mList []AdminMenu
	for i, _ := range *m {
		if CheckInInt(idList, (*m)[i].ID) {
			mList = append(mList, (*m)[i])
		}
	}

	GetTreeMenu(&t, 0, &mList)
	return &t, err
}
