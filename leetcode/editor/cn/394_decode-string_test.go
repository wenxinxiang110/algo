package main

import (
	"testing"
)

func TestDecodeString(t *testing.T) {
	t.Run("示例1：3[a]2[bc]", func(t *testing.T) {
		s := "3[a]2[bc]"
		got := decodeString(s)
		expected := "aaabcbc"

		if got != expected {
			t.Errorf("decodeString(\"%s\") = \"%s\", 期望 \"%s\"", s, got, expected)
		}
	})

	t.Run("示例2：3[a2[c]]", func(t *testing.T) {
		s := "3[a2[c]]"
		got := decodeString(s)
		expected := "accaccacc"

		if got != expected {
			t.Errorf("decodeString(\"%s\") = \"%s\", 期望 \"%s\"", s, got, expected)
		}
	})

	t.Run("示例3：2[abc]3[cd]ef", func(t *testing.T) {
		s := "2[abc]3[cd]ef"
		got := decodeString(s)
		expected := "abcabccdcdcdef"

		if got != expected {
			t.Errorf("decodeString(\"%s\") = \"%s\", 期望 \"%s\"", s, got, expected)
		}
	})

	t.Run("示例4：abc3[cd]xyz", func(t *testing.T) {
		s := "abc3[cd]xyz"
		got := decodeString(s)
		expected := "abccdcdcdxyz"

		if got != expected {
			t.Errorf("decodeString(\"%s\") = \"%s\", 期望 \"%s\"", s, got, expected)
		}
	})

	t.Run("简单嵌套：2[3[a]]", func(t *testing.T) {
		s := "2[3[a]]"
		got := decodeString(s)
		expected := "aaaaaa"

		if got != expected {
			t.Errorf("decodeString(\"%s\") = \"%s\", 期望 \"%s\"", s, got, expected)
		}
	})

	t.Run("多层嵌套：3[a2[b3[c]]]", func(t *testing.T) {
		s := "3[a2[b3[c]]]"
		got := decodeString(s)
		expected := "abcccbcccabcccbcccabcccbccc"

		if got != expected {
			t.Errorf("decodeString(\"%s\") = \"%s\", 期望 \"%s\"", s, got, expected)
		}
	})

	t.Run("无括号：abc", func(t *testing.T) {
		s := "abc"
		got := decodeString(s)
		expected := "abc"

		if got != expected {
			t.Errorf("decodeString(\"%s\") = \"%s\", 期望 \"%s\"", s, got, expected)
		}
	})

	t.Run("单个重复：10[a]", func(t *testing.T) {
		s := "10[a]"
		got := decodeString(s)
		expected := "aaaaaaaaaa"

		if got != expected {
			t.Errorf("decodeString(\"%s\") = \"%s\", 期望 \"%s\"", s, got, expected)
		}
	})

	t.Run("大数字：100[a]", func(t *testing.T) {
		s := "100[a]"
		got := decodeString(s)

		// 验证长度是否正确
		expectedLength := 100
		if len(got) != expectedLength {
			t.Errorf("decodeString(\"%s\") 返回字符串长度 %d，期望 %d", s, len(got), expectedLength)
		}

		// 验证所有字符都是 'a'
		for i, ch := range got {
			if ch != 'a' {
				t.Errorf("decodeString(\"%s\") 在第 %d 个位置返回了错误的字符 '%c'，期望 'a'", s, i, ch)
			}
		}
	})

	t.Run("空字符串", func(t *testing.T) {
		s := ""
		got := decodeString(s)
		expected := ""

		if got != expected {
			t.Errorf("decodeString(\"%s\") = \"%s\", 期望 \"%s\"", s, got, expected)
		}
	})

	t.Run("复杂混合：2[ab]3[cd2[ef]]", func(t *testing.T) {
		s := "2[ab]3[cd2[ef]]"
		got := decodeString(s)
		expected := "ababcd efef cd efef cd efef"

		if got != expected {
			t.Errorf("decodeString(\"%s\") = \"%s\", 期望 \"%s\"", s, got, expected)
		}
	})

	t.Run("验证输出长度限制", func(t *testing.T) {
		// 测试用例保证输出的长度不会超过 10⁵
		s := "300[a300[b]]"
		got := decodeString(s)

		if len(got) > 100000 {
			t.Errorf("decodeString(\"%s\") 返回字符串长度 %d，超过了 10⁵ 的限制", s, len(got))
		}
	})

	t.Run("边界情况：最大数字300", func(t *testing.T) {
		s := "300[a]"
		got := decodeString(s)

		expectedLength := 300
		if len(got) != expectedLength {
			t.Errorf("decodeString(\"%s\") 返回字符串长度 %d，期望 %d", s, len(got), expectedLength)
		}

		for i, ch := range got {
			if ch != 'a' {
				t.Errorf("decodeString(\"%s\") 在第 %d 个位置返回了错误的字符 '%c'，期望 'a'", s, i, ch)
			}
		}
	})

	t.Run("混合字母和数字：a2[b]c3[d]e", func(t *testing.T) {
		s := "a2[b]c3[d]e"
		got := decodeString(s)
		expected := "abbcddde"

		if got != expected {
			t.Errorf("decodeString(\"%s\") = \"%s\", 期望 \"%s\"", s, got, expected)
		}
	})

	t.Run("深度嵌套：2[3[4[a]]]", func(t *testing.T) {
		s := "2[3[4[a]]]"
		got := decodeString(s)
		expectedLength := 2 * 3 * 4 // 24个a

		if len(got) != expectedLength {
			t.Errorf("decodeString(\"%s\") 返回字符串长度 %d，期望 %d", s, len(got), expectedLength)
		}

		for i, ch := range got {
			if ch != 'a' {
				t.Errorf("decodeString(\"%s\") 在第 %d 个位置返回了错误的字符 '%c'，期望 'a'", s, i, ch)
			}
		}
	})
}
