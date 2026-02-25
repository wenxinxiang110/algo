package main

//你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
//
// 在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表
//示如果要学习课程 ai 则 必须 先学习课程 bi 。
//
//
// 例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
//
//
// 请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。
//
//
//
// 示例 1：
//
//
//输入：numCourses = 2, prerequisites = [[1,0]]
//输出：true
//解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。
//
// 示例 2：
//
//
//输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
//输出：false
//解释：总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。
//
//
//
// 提示：
//
//
// 1 <= numCourses <= 2000
// 0 <= prerequisites.length <= 5000
// prerequisites[i].length == 2
// 0 <= ai, bi < numCourses
// prerequisites[i] 中的所有课程对 互不相同
//
//
// Related Topics 深度优先搜索 广度优先搜索 图 拓扑排序 👍 2302 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func canFinish(numCourses int, prerequisites [][]int) bool {

	// 记录出度
	inDegree := make([]int, numCourses)

	// 构建邻接表
	var dependOn = make(map[int][]int)
	for _, prerequisite := range prerequisites {
		depender := prerequisite[0]
		dependent := prerequisite[1]
		dependOn[dependent] = append(dependOn[dependent], depender)
		inDegree[depender]++
	}

	queue := make([]int, 0)

	// 填充入度为0的课，这些课可以直接选
	for course, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, course)
		}
	}

	var selected = 0
	for len(queue) > 0 {
		//pop queue elements
		selected += len(queue)
		elems := queue[:]
		queue = []int{}

		for _, elem := range elems {
			toEnqueue := dependOn[elem]
			for _, course := range toEnqueue {
				inDegree[course]--
				if inDegree[course] == 0 {
					queue = append(queue, course)
				}
			}

		}
	}
	return selected == numCourses
}

//leetcode submit region end(Prohibit modification and deletion)
