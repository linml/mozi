package main

import (
	"fmt"
	"github.com/xiuos/mozi/xorm"
)

type User struct {
	UserID   int
	Username string
	Age      int
	School   string
}

func (u *User) Field() []string {
	return []string{"user_id", "username", "age", "school"}
}

func main() {
	o := xorm.Orm{}
	u := User{}
	str, _ := o.Table("tb_user").Query().Where("username = ? AND age > ?", "test", 18).Do(&u)
	fmt.Println(str)
	fmt.Println(o.Args())
}
