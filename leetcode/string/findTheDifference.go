package string

//给定两个字符串 s 和 t，它们只包含小写字母。
//
//字符串 t 由字符串 s 随机重排，然后在随机位置添加一个字母。
//
//请找出在 t 中被添加的字母。
//
//提示：
//
//0 <= s.length <= 1000
//t.length == s.length + 1
//s 和 t 只包含小写字母

func findTheDifference(s string, t string) byte {
	var ret int32 = 0

	for _, i := range s {
		ret ^= i
	}
	for _, i := range t {
		ret ^= i
	}
	return byte(ret)
}
