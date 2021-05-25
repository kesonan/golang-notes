package bytedance

import (
	"fmt"
	"strings"
	"testing"
)

//input = ["year", "month", "day"];

//func(input)
//"ymd". "yma", "ymy" .. "rhy"

func TestArray(t *testing.T) {
	ret := listAll([]string{"year", "month", "day"})
	for _, e := range ret {
		fmt.Println(e)
	}
}

func listAll(list []string) []string {
	if len(list) < 2 {
		return list
	}

	var (
		n      = len(list)
		l1, l2 = strings.Split(list[0], ""), []string{}
		i      = 1
	)

	for i < n {
		l2 = strings.Split(list[i], "")
		l1 = listTwoString(l1, l2)
		i += 1
	}

	return l1
}

func listTwoString(s1, s2 []string) []string {
	var ret []string
	for _, e := range s1 {
		for _, ie := range s2 {
			ret = append(ret, e+ie)
		}
	}

	return ret
}
