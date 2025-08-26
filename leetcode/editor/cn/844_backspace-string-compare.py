# package main


# 给定 s 和 t 两个字符串，当它们分别被输入到空白的文本编辑器后，如果两者相等，返回 true 。# 代表退格字符。
# 
#  注意：如果对空文本输入退格字符，文本继续为空。 
# 
#  
# 
#  示例 1： 
# 
#  
# 输入：s = "ab#c", t = "ad#c"
# 输出：true
# 解释：s 和 t 都会变成 "ac"。
#  
# 
#  示例 2： 
# 
#  
# 输入：s = "ab##", t = "c#d#"
# 输出：true
# 解释：s 和 t 都会变成 ""。
#  
# 
#  示例 3： 
# 
#  
# 输入：s = "a#c", t = "b"
# 输出：false
# 解释：s 会变成 "c"，但 t 仍然是 "b"。 
# 
#  
# 
#  提示： 
# 
#  
#  1 <= s.length, t.length <= 200 
#  s 和 t 只含有小写字母以及字符 '#' 
#  
# 
#  
# 
#  进阶： 
# 
#  
#  你可以用 O(n) 的时间复杂度和 O(1) 的空间复杂度解决该问题吗？ 
#  
# 
#  Related Topics 栈 双指针 字符串 模拟 👍 829 👎 0


# leetcode submit region begin(Prohibit modification and deletion)
class Solution:
    def backspaceCompare(self, s: str, t: str) -> bool:
        back = '#'

        s_iter, t_iter = len(s) - 1, len(t) - 1
        s_skip, t_skip = 0, 0

        while s_iter >= 0 or t_iter >= 0:
            # 如果 back 是 #，s--

            while s_iter >= 0:
                # 退格，记录需要退格的数量
                if s[s_iter] == back:
                    s_skip += 1
                    s_iter -= 1
                elif s_skip > 0:
                    # 非退格符，开始生效
                    s_skip -= 1
                    s_iter -= 1
                else:
                    break

            while t_iter >= 0:
                if t[t_iter] == back:
                    t_skip += 1
                    t_iter -= 1
                elif t_skip > 0:
                    t_skip -= 1
                    t_iter -= 1
                else:
                    break
            # s和 t 其中一个 iter < 0，返回 false
            if s_iter >= 0 and t_iter >= 0:
                if s[s_iter] != t[t_iter]:
                    return False
            elif s_iter >= 0 or t_iter >= 0:
                return False
            s_iter -= 1
            t_iter -= 1

        return True

# leetcode submit region end(Prohibit modification and deletion)
