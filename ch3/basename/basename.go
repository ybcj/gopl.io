package basename

//basename 移除路径部分和.后缀
//e.g,. a +> a, a.go => a, a/b/c.go => c, a/b.c.go => a.c

import "strings"

func Basename1(s string) string {
	// 不使用任何库函数
	// 将最后一个‘/’和之前的部分都舍弃
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	// 保留最后一个‘.’之前的所有内容
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func Basename2(s string) string {
	// 使用strings包
	slash := strings.LastIndex(s, "/") // 如果未找到“/”，则LastIndex返回-1
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
