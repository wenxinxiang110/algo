package main

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	t.Run("示例测试", func(t *testing.T) {
		// 输入: ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
		// [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
		// 输出: [null, null, null, 1, null, -1, null, -1, 3, 4]

		cache := Constructor(2)

		// put(1, 1)
		cache.Put(1, 1)

		// put(2, 2)
		cache.Put(2, 2)

		// get(1) 应该返回 1
		if got := cache.Get(1); got != 1 {
			t.Errorf("Get(1) = %v, want 1", got)
		}

		// put(3, 3) - 这会使得关键字2作废
		cache.Put(3, 3)

		// get(2) 应该返回 -1
		if got := cache.Get(2); got != -1 {
			t.Errorf("Get(2) = %v, want -1", got)
		}

		// put(4, 4) - 这会使得关键字1作废
		cache.Put(4, 4)

		// get(1) 应该返回 -1
		if got := cache.Get(1); got != -1 {
			t.Errorf("Get(1) = %v, want -1", got)
		}

		// get(3) 应该返回 3
		if got := cache.Get(3); got != 3 {
			t.Errorf("Get(3) = %v, want 3", got)
		}

		// get(4) 应该返回 4
		if got := cache.Get(4); got != 4 {
			t.Errorf("Get(4) = %v, want 4", got)
		}
	})

	t.Run("容量为1的缓存", func(t *testing.T) {
		cache := Constructor(1)

		cache.Put(1, 10)
		if got := cache.Get(1); got != 10 {
			t.Errorf("Get(1) = %v, want 10", got)
		}

		// 添加新元素，应该淘汰旧元素
		cache.Put(2, 20)
		if got := cache.Get(1); got != -1 {
			t.Errorf("Get(1) = %v, want -1", got)
		}
		if got := cache.Get(2); got != 20 {
			t.Errorf("Get(2) = %v, want 20", got)
		}
	})

	t.Run("更新已存在的键值", func(t *testing.T) {
		cache := Constructor(2)

		cache.Put(1, 10)
		cache.Put(2, 20)

		// 更新键1的值
		cache.Put(1, 100)

		if got := cache.Get(1); got != 100 {
			t.Errorf("Get(1) = %v, want 100", got)
		}
		if got := cache.Get(2); got != 20 {
			t.Errorf("Get(2) = %v, want 20", got)
		}

		// 添加新元素，应该淘汰最久未使用的键2
		cache.Put(3, 30)
		if got := cache.Get(2); got != -1 {
			t.Errorf("Get(2) = %v, want -1", got)
		}
		if got := cache.Get(1); got != 100 {
			t.Errorf("Get(1) = %v, want 100", got)
		}
		if got := cache.Get(3); got != 30 {
			t.Errorf("Get(3) = %v, want 30", got)
		}
	})

	t.Run("访问顺序影响淘汰策略", func(t *testing.T) {
		cache := Constructor(3)

		cache.Put(1, 10)
		cache.Put(2, 20)
		cache.Put(3, 30)

		// 访问键1，使其成为最近使用的
		cache.Get(1)

		// 添加新元素，应该淘汰最久未使用的键2
		cache.Put(4, 40)

		if got := cache.Get(2); got != -1 {
			t.Errorf("Get(2) = %v, want -1", got)
		}
		if got := cache.Get(1); got != 10 {
			t.Errorf("Get(1) = %v, want 10", got)
		}
		if got := cache.Get(3); got != 30 {
			t.Errorf("Get(3) = %v, want 30", got)
		}
		if got := cache.Get(4); got != 40 {
			t.Errorf("Get(4) = %v, want 40", got)
		}
	})

	t.Run("边界值测试", func(t *testing.T) {
		// 容量为0的缓存
		cache := Constructor(0)

		cache.Put(1, 10)
		if got := cache.Get(1); got != -1 {
			t.Errorf("Get(1) = %v, want -1", got)
		}

		// 容量为3000（最大容量）
		cache = Constructor(3000)
		for i := 0; i < 3000; i++ {
			cache.Put(i, i*10)
		}

		// 验证所有值都能正确获取
		for i := 0; i < 3000; i++ {
			if got := cache.Get(i); got != i*10 {
				t.Errorf("Get(%d) = %v, want %d", i, got, i*10)
			}
		}

		// 添加第3001个元素，应该淘汰第一个元素
		cache.Put(3000, 30000)
		if got := cache.Get(0); got != -1 {
			t.Errorf("Get(0) = %v, want -1", got)
		}
		if got := cache.Get(3000); got != 30000 {
			t.Errorf("Get(3000) = %v, want 30000", got)
		}
	})
}

// TestLRUCacheFrameworkExample 展示如何使用框架处理复杂的测试场景
func TestLRUCacheFrameworkExample(t *testing.T) {
	runner := LRUCacheTestRunner()

	// 从LeetCode格式直接创建测试用例
	methods := []string{"Constructor", "put", "put", "get", "put", "get", "put", "get", "get", "get"}
	inputs := [][]int{
		{2},    // Constructor
		{1, 1}, // put
		{2, 2}, // put
		{1},    // get
		{3, 3}, // put
		{2},    // get
		{4, 4}, // put
		{1},    // get
		{3},    // get
		{4},    // get
	}
	expected := []interface{}{nil, nil, nil, 1, nil, -1, nil, -1, 3, 4}

	testCase := CreateTestCaseFromLeetCode("LeetCode标准示例", methods, inputs, expected)

	// 打印测试用例信息（调试用）
	PrintTestCase(testCase)

	runner.RunTestCase(t, testCase)
}
