package common

import (
	"fmt"
	"time"
)

func GetChinaSpring(year int) time.Time {
	/*
		立春准确时间的计算方法
		计算公式：[Y*D+C]-L
		公式解读：年数的后2位乘0.2422加3.87取整数减闰年数。21世纪C值=3.87，22世纪C值=4.15。
		举例说明：2058年立春日期的计算步骤[58×.0.2422+3.87]-[（58-1）/4]=17-14=3，则2月3日立春。
	*/
	d := int(float32(year%100)*0.2422+3.87) - int((year%100-1)/4)
	dStr := fmt.Sprintf("%d02%2d000000", year, d)
	r, _ := Str2Time(dStr)
	return r
}

func GetNowChinaSpring() time.Time {
	year := time.Now().Year()
	return GetChinaSpring(year)
}


func GetChinaZodiac() int {
	/*
		子鼠、丑牛、寅虎、卯兔、辰龙、巳蛇、午马、未羊、申猴、酉鸡、戌狗、亥猪。
	*/
	d := time.Now()
	zodiac := d.Year() - 2008

	if d.Before(GetChinaSpring(d.Year())){
		zodiac -= 1
	}
	return zodiac % 12
}

