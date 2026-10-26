package main

//给定一个链表的头节点 head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
//
// 如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到
//链表中的位置（索引从 0 开始）。如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
//
// 不允许修改 链表。
//
//
//
//
//
//
// 示例 1：
//
//
//
//
//输入：head = [3,2,0,-4], pos = 1
//输出：返回索引为 1 的链表节点
//解释：链表中有一个环，其尾部连接到第二个节点。
//
//
// 示例 2：
//
//
//
//
//输入：head = [1,2], pos = 0
//输出：返回索引为 0 的链表节点
//解释：链表中有一个环，其尾部连接到第一个节点。
//
//
// 示例 3：
//
//
//
//
//输入：head = [1], pos = -1
//输出：返回 null
//解释：链表中没有环。
//
//
//
//
// 提示：
//
//
// 链表中节点的数目范围在范围 [0, 10⁴] 内
// -10⁵ <= Node.val <= 10⁵
// pos 的值为 -1 或者链表中的一个有效索引
//
//
//
//
// 进阶：你是否可以使用 O(1) 空间解决此题？
//
// Related Topics 哈希表 链表 双指针 👍 2630 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	return detectCycleDoublePtr(head)
}

// detectCycleDoublePtr 使用双指针（快慢指针）解法，空间复杂度 O(1)，满足进阶要求。
//
// 思路分两步：
//
// 第一步：判断是否有环，并找到快慢指针的「第一次相遇点」。
//   - 慢指针 slow 每次走 1 步，快指针 fast 每次走 2 步。
//   - 若链表无环，fast 会先到达链表末尾（fast 或 fast.Next 为 nil），循环结束返回 nil。
//   - 若有环，由于 fast 比 slow 快，两者终会在环内相遇（fast 在环里「套圈」slow）。
//
// 第二步：推导环的入口节点。这是本解法的核心，下面给出数学证明：
//   - 设：链表头到环入口的距离为 a；环入口到相遇点的距离为 b；相遇点再绕回环入口的距离为 c。
//     则环的总长为 b + c。
//   - 相遇时，慢指针走过的总路程 = a + b（在环里可能绕了若干圈，但慢指针第一次被追上时最多绕不到一圈，
//     这里按 a + b 计算即可，绕圈部分可被下面的等式消掉）。
//   - 快指针走过的总路程 = a + b + n*(b + c)，其中 n 是快指针在环里多绕的圈数（n >= 1）。
//   - 又因为 fast 速度是 slow 的 2 倍，所以「快指针路程 = 2 * 慢指针路程」：
//     a + b + n*(b + c) = 2*(a + b)
//     化简得：
//     a = n*(b + c) - b = (n - 1)*(b + c) + c
//   - 这个式子的含义是：从链表头走 a 步，等于「从相遇点走 c 步，再绕环 (n-1) 圈」。
//     无论从相遇点怎么绕完整的环，最终都会回到相遇点后再走 c 步到达环入口。
//
// 因此，让两个指针分别从「链表头 head」和「相遇点 slow」出发，每次各走 1 步，
// 它们必然在「环入口」第一次相遇（此时 idx1 == idx2 且恰好走了 a 步 / c+(n-1)圈 步），
// 该节点就是环的入口，直接返回即可。
func detectCycleDoublePtr(head *ListNode) *ListNode {
	var fast, slow = head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			// 已确认有环且 slow、fast 在环内相遇，
			// 分别从头节点和相遇点同步前进，二者必在环入口相遇。
			var idx1, idx2 = head, slow
			for idx1 != idx2 {
				idx1 = idx1.Next
				idx2 = idx2.Next
			}
			return idx1
		}
	}
	return nil
}

// hash表，第一个出现的就是环的入口
func detectCycleHash(head *ListNode) *ListNode {
	var hash = make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := hash[head]; ok {
			return head
		}
		hash[head] = struct{}{}
	}
	return nil
}

//leetcode submit region end(Prohibit modification and deletion)
