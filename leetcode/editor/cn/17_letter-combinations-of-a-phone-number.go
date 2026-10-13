package main

//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
//
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
//
//
//
//
//
// 示例 1：
//
//
//输入：digits = "23"
//输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
//
//
// 示例 2：
//
//
//输入：digits = "2"
//输出：["a","b","c"]
//
//
//
//
// 提示：
//
//
// 1 <= digits.length <= 4
// digits[i] 是范围 ['2', '9'] 的一个数字。
//
//
// Related Topics 哈希表 字符串 回溯 👍 3197 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func letterCombinations(digits string) (output []string) {
	if len(digits) == 0 {
		return
	}

	table := map[uint8][]rune{
		2: {'a', 'b', 'c'},
		3: {'d', 'e', 'f'},
		4: {'g', 'h', 'i'},
		5: {'j', 'k', 'l'},
		6: {'m', 'n', 'o'},
		7: {'p', 'q', 'r', 's'},
		8: {'t', 'u', 'v'},
		9: {'w', 'x', 'y', 'z'},
	}

	var backtrace func(depth int, path []rune)

	backtrace = func(depth int, path []rune) {

		if depth == len(digits) {
			output = append(output, string(path))
			return
		}
		for _, r := range table[digits[depth]-'0'] {

			path = append(path, r)

			backtrace(depth+1, path)

			path = path[:len(path)-1]
		}
	}

	backtrace(0, nil)
	return output
}

//leetcode submit region end(Prohibit modification and deletion)
