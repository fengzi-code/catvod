package utils

import (
	"fmt"
	"regexp"
	"strings"
)

//取文本中间,str原始文本,start文本前面的字符,end文本后面的字符
func GetBetweenStr(str, start, end string) string {
	n := strings.Index(str, start)
	if n == -1 {
		n = 0
		return ""
	} else {
		n = n + len(start) // 增加了else，不加的会把start带上
	}
	str = string([]byte(str)[n:])
	m := strings.Index(str, end)
	if m == -1 {
		m = len(str)
		return ""
	}
	str = string([]byte(str)[:m])
	return str
}

//批量取文本中间,str原始文本,start文本前面的字符,end文本后面的字符
func GetBetweenStrs(str, start, end string) []string {
	var strs []string
	mc := start + `(.*?)` + end
	r := regexp.MustCompile(mc)
	tmpStrs := r.FindAllString(str, -1)
	for _, s := range tmpStrs {
		s = strings.Replace(s, start, "", -1)
		s = strings.Replace(s, end, "", -1)
		strs = append(strs, s)
	}
	fmt.Println(mc, len(strs))
	return strs
}
