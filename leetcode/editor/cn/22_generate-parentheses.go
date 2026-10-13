package main

//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
//
//
// 示例 1：
//
//
//输入：n = 3
//输出：["((()))","(()())","(())()","()(())","()()()"]
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：["()"]
//
//
//
//
// 提示：
//
//
// 1 <= n <= 8
//
//
// Related Topics 字符串 动态规划 回溯 👍 4025 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func generateParenthesis(n int) (output []string) {
	if n == 0 {
		return
	}

	const left, right = '(', ')'

	visit := map[uint8]int{
		left:  n,
		right: n,
	}
	var backtrace func(depth int, path []rune)
	backtrace = func(depth int, path []rune) {
		if depth == 2*n {
			output = append(output, string(path))
			return
		}
		if visit[left] > 0 {
			path = append(path, left)
			visit[left]--
			backtrace(depth+1, path)
			visit[left]++
			path = path[:len(path)-1]
		}

		if visit[right] > visit[left] {
			path = append(path, right)
			visit[right]--
			backtrace(depth+1, path)
			visit[right]++
			path = path[:len(path)-1]
		}

	}

	backtrace(0, nil)
	return
}

//leetcode submit region end(Prohibit modification and deletion)
