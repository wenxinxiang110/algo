# package main

# 给你一个正整数 n ，生成一个包含 1 到 n² 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
#
#
#
#  示例 1：
#
#
# 输入：n = 3
# 输出：[[1,2,3],[8,9,4],[7,6,5]]
#
#
#  示例 2：
#
#
# 输入：n = 1
# 输出：[[1]]
#
#
#
#
#  提示：
#
#
#  1 <= n <= 20
#
#
#  Related Topics 数组 矩阵 模拟 👍 1475 👎 0


# leetcode submit region begin(Prohibit modification and deletion)
class Solution(object):
    def generateMatrix(self, n):
        """
        :type n: int
        :rtype: List[List[int]]
        """
        matrix = [[0 for _ in range(n)] for _ in range(n)]
        top, bottom, left, right = 0, n - 1, 0, n - 1
        cal = 1
        while True:
            # top
            for i in range(left, right):
                matrix[top][i] = cal
                cal += 1

            for i in range(top, bottom):
                matrix[i][right] = cal
                cal += 1

            for i in range(right, left, -1):
                matrix[bottom][i] = cal
                cal += 1

            for i in range(bottom, top, -1):
                matrix[i][left] = cal
                cal += 1

            top += 1
            bottom -= 1
            left += 1
            right -= 1

            if top >= bottom or left >= right:
                break

        if n % 2 != 0:
            matrix[int(n / 2)][int(n / 2)] = cal

        return matrix

# leetcode submit region end(Prohibit modification and deletion)
