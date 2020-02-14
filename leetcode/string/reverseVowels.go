package string

//345. 反转字符串中的元音字母
/*编写一个函数，以字符串作为输入，反转该字符串中的元音字母。

示例 1:

输入: "hello"
输出: "holle"
示例 2:

输入: "leetcode"
输出: "leotcede"
说明:
元音字母不包含字母"y"。
链接：https://leetcode-cn.com/problems/reverse-vowels-of-a-string  */
func reverseVowels(s string) string {
	i, j := 0, len(s)-1

	//var temp uint8

	for i != j {

		for i < j && !IsVowel(s[i]) {
			i++
		}

		for i < j && !IsVowel(s[j]) {
			j++
		}

		//temp = s[j]
		//s[j] = s[i]
		//s[i] = temp
	}

	return s
}

func IsVowel(s uint8) bool {
	if s == 'a' || s == 'e' || s == 'i' || s == 'o' || s == 'u' {
		return true
	}
	return false
}
