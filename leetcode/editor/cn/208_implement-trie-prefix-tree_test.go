package main

import (
	"testing"
)

func TestTrie(t *testing.T) {
	tests := []struct {
		name       string
		operations []string
		params     []string
		expected   []interface{}
	}{
		{
			name:       "题目示例",
			operations: []string{"Trie", "insert", "search", "search", "startsWith", "insert", "search"},
			params:     []string{"", "apple", "apple", "app", "app", "app", "app"},
			expected:   []interface{}{nil, nil, true, false, true, nil, true},
		},
		{
			name:       "空字符串处理",
			operations: []string{"Trie", "insert", "search", "startsWith"},
			params:     []string{"", "", "", ""},
			expected:   []interface{}{nil, nil, true, true},
		},
		{
			name:       "单个字符操作",
			operations: []string{"Trie", "insert", "search", "startsWith", "insert", "search"},
			params:     []string{"", "a", "a", "a", "b", "b"},
			expected:   []interface{}{nil, nil, true, true, nil, true},
		},
		{
			name:       "前缀匹配测试",
			operations: []string{"Trie", "insert", "startsWith", "insert", "startsWith"},
			params:     []string{"", "hello", "hel", "hell", "he"},
			expected:   []interface{}{nil, nil, true, nil, true},
		},
		{
			name:       "不存在单词搜索",
			operations: []string{"Trie", "search", "startsWith"},
			params:     []string{"", "nonexistent", "pre"},
			expected:   []interface{}{nil, false, false},
		},
		{
			name:       "多个单词插入和搜索",
			operations: []string{"Trie", "insert", "insert", "insert", "search", "search", "search", "startsWith"},
			params:     []string{"", "cat", "car", "card", "cat", "car", "card", "ca"},
			expected:   []interface{}{nil, nil, nil, nil, true, true, true, true},
		},
		{
			name:       "重叠前缀测试",
			operations: []string{"Trie", "insert", "insert", "search", "startsWith"},
			params:     []string{"", "application", "apple", "app", "app"},
			expected:   []interface{}{nil, nil, nil, false, true},
		},
		{
			name:       "长字符串测试",
			operations: []string{"Trie", "insert", "search", "startsWith"},
			params:     []string{"", "abcdefghijklmnopqrstuvwxyz", "abcdefghijklmnopqrstuvwxyz", "abc"},
			expected:   []interface{}{nil, nil, true, true},
		},
		{
			name:       "重复插入相同单词",
			operations: []string{"Trie", "insert", "insert", "search"},
			params:     []string{"", "test", "test", "test"},
			expected:   []interface{}{nil, nil, nil, true},
		},
		{
			name:       "边界长度测试",
			operations: []string{"Trie", "insert", "search", "startsWith"},
			params:     []string{"", "a", "a", "a"},
			expected:   []interface{}{nil, nil, true, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var trie *Trie = &Trie{}

			for i, op := range tt.operations {
				switch op {
				case "Trie":
					*trie = Constructor()
				case "insert":
					trie.Insert(tt.params[i])
				case "search":
					result := trie.Search(tt.params[i])
					if result != tt.expected[i] {
						t.Errorf("Search(%s) = %v, want %v", tt.params[i], result, tt.expected[i])
					}
				case "startsWith":
					result := trie.StartsWith(tt.params[i])
					if result != tt.expected[i] {
						t.Errorf("StartsWith(%s) = %v, want %v", tt.params[i], result, tt.expected[i])
					}
				}
			}
		})
	}
}

// 单独测试构造函数
func TestConstructor(t *testing.T) {
	//trie := Constructor()
	//if trie == nil {
	//	t.Error("Constructor() should return a non-nil Trie instance")
	//}
}

// 性能测试：多次操作
func TestTriePerformance(t *testing.T) {
	trie := Constructor()

	// 插入多个单词
	words := []string{"hello", "world", "test", "go", "programming", "algorithm", "data", "structure"}
	for _, word := range words {
		trie.Insert(word)
	}

	// 验证插入的单词都能搜索到
	for _, word := range words {
		if !trie.Search(word) {
			t.Errorf("Search(%s) should return true after insertion", word)
		}
	}

	// 验证前缀匹配
	prefixes := []string{"he", "wo", "te", "g", "pro", "alg", "dat", "stru"}
	for _, prefix := range prefixes {
		if !trie.StartsWith(prefix) {
			t.Errorf("StartsWith(%s) should return true for existing words", prefix)
		}
	}

	// 验证不存在的单词
	if trie.Search("nonexistent") {
		t.Error("Search(\"nonexistent\") should return false")
	}

	if trie.StartsWith("xyz") {
		t.Error("StartsWith(\"xyz\") should return false")
	}
}
