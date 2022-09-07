package utils

import "strings"

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
