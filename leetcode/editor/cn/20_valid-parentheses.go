package main

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
// 有效字符串需满足：
//
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号。
//
//
//
//
// 示例 1：
//
//
// 输入：s = "()"
//
//
// 输出：true
//
// 示例 2：
//
//
// 输入：s = "()[]{}"
//
//
// 输出：true
//
// 示例 3：
//
//
// 输入：s = "(]"
//
//
// 输出：false
//
// 示例 4：
//
//
// 输入：s = "([])"
//
//
// 输出：true
//
//
//
// 提示：
//
//
// 1 <= s.length <= 10⁴
// s 仅由括号 '()[]{}' 组成
//
//
// Related Topics 栈 字符串 👍 4561 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	stack := Stack(make([]byte, 0))
	for _, b := range []byte(s) {
		// 左括号
		if b == '(' || b == '{' || b == '[' {
			stack.Push(b)
			continue
		} else {
			left, ok := stack.Pop()
			if !ok {
				return false
			}
			switch left {
			case '(':
				if b != ')' {
					return false
				}
			case '{':
				if b != '}' {
					return false
				}
			case '[':
				if b != ']' {
					return false
				}
			}
		}
	}
	return len(stack) == 0
}

type Stack []byte

func (s *Stack) Push(v byte) {
	*s = append(*s, v)
}

func (s *Stack) Pop() (byte, bool) {
	if len(*s) == 0 {
		return 0, false
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v, true
}

//leetcode submit region end(Prohibit modification and deletion)
