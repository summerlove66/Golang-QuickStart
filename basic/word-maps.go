package basic

import (
	"strings"
)

//统计单词
func WordCount(s string) map[string]int {
	ss := strings.Fields(s)
	c := make(map[string]int)
	for _, v := range ss {
		ele, m := c[v]
		if m == true {
			c[v] = ele + 1
		} else {
			c[v] = 1
		}
	}
	return c

}
