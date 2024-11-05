package kmp

// Next kmp前缀表,所有值-1
func Next(s string) []int {
	b := []byte(s)
	n := len(s)
	table := make([]int, n)

	//j指向前缀末尾位置，同时表示该位置上的最长前缀后缀长度
	var j = -1
	// 第一个位置可以省略查找的过程
	table[0] = j

	//i指向后缀末尾位置
	for i := 1; i < n; i++ {
		for j >= 0 && b[i] != b[j+1] {
			j = table[j]
		}
		if b[i] == b[j+1] { // 找到相同的前后缀
			j++
		}
		table[i] = j // 将j（前缀的长度）赋给next[i]
	}

	return table
}

// First kmp匹配,返回第一个匹配位置
func First(haystack, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}
	b := []byte(haystack)
	n := []byte(needle)
	next := Next(needle)
	j := -1
	for i := 0; i < len(haystack); i++ {
		for j >= 0 && b[i] != n[j+1] {
			j = next[j]
		}
		if b[i] == n[j+1] {
			j++
		}
		if j == len(needle)-1 {
			return i - j
		}

	}
	return -1
}
