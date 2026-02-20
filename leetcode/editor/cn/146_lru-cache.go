package main

//
// 请你设计并实现一个满足
// LRU (最近最少使用) 缓存 约束的数据结构。
//
//
//
// 实现
// LRUCache 类：
//
//
//
//
//
// LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
// int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
// void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组
//key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
//
//
//
//
// 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
//
//
//
// 示例：
//
//
//输入
//["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
//[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
//输出
//[null, null, null, 1, null, -1, null, -1, 3, 4]
//
//解释
//LRUCache lRUCache = new LRUCache(2);
//lRUCache.put(1, 1); // 缓存是 {1=1}
//lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
//lRUCache.get(1);    // 返回 1
//lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
//lRUCache.get(2);    // 返回 -1 (未找到)
//lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
//lRUCache.get(1);    // 返回 -1 (未找到)
//lRUCache.get(3);    // 返回 3
//lRUCache.get(4);    // 返回 4
//
//
//
//
// 提示：
//
//
// 1 <= capacity <= 3000
// 0 <= key <= 10000
// 0 <= value <= 10⁵
// 最多调用 2 * 10⁵ 次 get 和 put
//
//
// Related Topics 设计 哈希表 链表 双向链表 👍 3717 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
type LRUCache struct {
	cache map[int]*DoubleListNode
	//  虚拟的前后指针
	head, tail *DoubleListNode

	capacity int
	length   int
}

// DoubleListNode 双向链表
type DoubleListNode struct {
	Key       int // 添加key字段，用于淘汰时从cache中删除
	Val       int
	Pre, Next *DoubleListNode
}

func Constructor(capacity int) LRUCache {
	head := &DoubleListNode{}
	tail := &DoubleListNode{}
	head.Next = tail
	tail.Pre = head
	return LRUCache{
		capacity: capacity,
		length:   0,
		cache:    make(map[int]*DoubleListNode),
		head:     head,
		tail:     tail,
	}

}

func (this *LRUCache) Get(key int) int {
	node := this.cache[key]
	if node == nil {
		return -1
	}
	// 断开连接
	this.removeNode(node)
	// move to tail
	this.addNodeInTail(node)

	return node.Val
}

// 插入一个新节点到尾部
func (this *LRUCache) addNodeInTail(node *DoubleListNode) {
	this.tail.Pre.Next = node
	node.Pre = this.tail.Pre
	node.Next = this.tail
	this.tail.Pre = node
}

// 断开原来链接
func (this *LRUCache) removeNode(node *DoubleListNode) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

func (this *LRUCache) Put(key int, value int) {
	node := this.cache[key]
	if node == nil {
		node = &DoubleListNode{Key: key, Val: value}
		this.cache[key] = node
		this.length++
	} else {
		node.Val = value
		this.removeNode(node)
	}
	// 尾巴插入一个新的 node
	this.addNodeInTail(node)

	// 淘汰策略
	if this.length > this.capacity {
		// 淘汰最久未使用的节点（头节点的下一个节点）
		toRemove := this.head.Next
		this.removeNode(toRemove)
		delete(this.cache, toRemove.Key)
		this.length--
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
//leetcode submit region end(Prohibit modification and deletion)
