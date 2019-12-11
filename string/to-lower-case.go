package string

/*
unicode编码表:https://blog.csdn.net/tgvincent/article/details/93884725
字符编码:https://www.cnblogs.com/notbecoder/p/4840783.html
wiki:https://en.wikipedia.org/wiki/Unicode
rune:转成unicode的code point
单引号创建 code point
https://studygolang.com/articles/13792

*/
func toLowerCase(str string) string {
	unicode := []rune(str)
	diff := 'a' - 'A'
	for i := 0; i < len(str); i++ {
		if unicode[i] >= 'A' && unicode[i] <= 'Z' {
			unicode[i] += diff
		}
	}
	return string(unicode)
}
