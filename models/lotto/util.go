package lotto

import (
	"sort"
	"strconv"
)

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

func SumInt(l []int) int {
	sum := 0
	for _, v := range l {
		sum += v
	}
	return sum
}

func GetInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

func MinVal(l []int) int {
	m := l[0]
	for i, _ := range l {
		if l[i] < m {
			m = l[i]
		}
	}
	return m
}

func MaxVal(l []int) int {
	m := l[0]
	for i, _ := range l {
		if l[i] > m {
			m = l[i]
		}
	}
	return m
}

func CheckInInt(dis []int, k int) bool {
	for i, _ := range dis {
		if dis[i] == k {
			return true
		}
	}
	return false
}

func CheckInStr(dis []string, k string) bool {
	for i, _ := range dis {
		if dis[i] == k {
			return true
		}
	}
	return false
}

func SliceRemoveDuplicates(slice []string) []string {

	sort.Strings(slice)

	i := 0

	var j int

	for {

		if i >= len(slice)-1 {

			break

		}

		for j = i + 1; j < len(slice) && slice[i] == slice[j]; j++ {

		}

		slice = append(slice[:i+1], slice[j:]...)

		i++

	}

	return slice

}

func Combination(m, n int) int {
	if n > m-n {
		n = m - n
	}
	c := 1
	for i := 0; i < n; i++ {
		c *= m - i
		c /= i + 1
	}
	return c
}
