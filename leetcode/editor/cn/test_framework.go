package main

import (
	"fmt"
	"reflect"
	"testing"
)

// TestCase 表示一个测试用例
type TestCase struct {
	Name     string
	Methods  []string
	Inputs   [][]interface{}
	Expected []interface{}
}

// TestRunner 通用的测试运行器
type TestRunner struct {
	Constructor func(int) interface{}
	MethodMap   map[string]func(interface{}, ...interface{}) interface{}
}

// NewTestRunner 创建新的测试运行器
func NewTestRunner(constructor func(int) interface{}, methodMap map[string]func(interface{}, ...interface{}) interface{}) *TestRunner {
	return &TestRunner{
		Constructor: constructor,
		MethodMap:   methodMap,
	}
}

// RunTestCase 执行单个测试用例
func (tr *TestRunner) RunTestCase(t *testing.T, tc TestCase) {
	t.Run(tc.Name, func(t *testing.T) {
		var instance interface{}

		for i, method := range tc.Methods {
			expected := tc.Expected[i]

			if method == "Constructor" {
				// 构造函数调用
				if len(tc.Inputs[i]) > 0 {
					capacity, ok := tc.Inputs[i][0].(int)
					if !ok {
						t.Fatalf("Constructor参数类型错误，期望int，得到%T", tc.Inputs[i][0])
					}
					instance = tr.Constructor(capacity)
				}
				// 构造函数期望输出为null
				if expected != nil {
					t.Errorf("Constructor期望输出为null，得到%v", expected)
				}
			} else {
				// 方法调用
				if instance == nil {
					t.Fatal("实例未初始化")
				}

				methodFunc, exists := tr.MethodMap[method]
				if !exists {
					t.Fatalf("未知方法: %s", method)
				}

				// 调用方法
				result := methodFunc(instance, tc.Inputs[i]...)

				// 验证结果
				if expected == nil {
					// 期望输出为null
					if result != nil {
						t.Errorf("%s期望输出为null，得到%v", method, result)
					}
				} else {
					if !reflect.DeepEqual(result, expected) {
						t.Errorf("%s期望输出%v，得到%v", method, expected, result)
					}
				}
			}
		}
	})
}

// RunTestCases 批量执行测试用例
func (tr *TestRunner) RunTestCases(t *testing.T, testCases []TestCase) {
	for _, tc := range testCases {
		tr.RunTestCase(t, tc)
	}
}

// LRUCacheTestRunner LRU缓存的专用测试运行器
func LRUCacheTestRunner() *TestRunner {
	methodMap := map[string]func(interface{}, ...interface{}) interface{}{
		"put": func(instance interface{}, args ...interface{}) interface{} {
			if len(args) != 2 {
				panic("put方法需要2个参数")
			}
			key, ok1 := args[0].(int)
			value, ok2 := args[1].(int)
			if !ok1 || !ok2 {
				panic("put方法参数类型错误")
			}
			cache := instance.(LRUCache)
			// 获取指针以调用方法
			cachePtr := &cache
			cachePtr.Put(key, value)
			return nil
		},
		"get": func(instance interface{}, args ...interface{}) interface{} {
			if len(args) != 1 {
				panic("get方法需要1个参数")
			}
			key, ok := args[0].(int)
			if !ok {
				panic("get方法参数类型错误")
			}
			cache := instance.(LRUCache)
			// 获取指针以调用方法
			cachePtr := &cache
			return cachePtr.Get(key)
		},
	}

	return NewTestRunner(func(capacity int) interface{} {
		return Constructor(capacity)
	}, methodMap)
}

// TestLRUCacheWithFramework 使用框架测试LRU缓存
func TestLRUCacheWithFramework(t *testing.T) {
	runner := LRUCacheTestRunner()

	testCases := []TestCase{
		{
			Name:    "示例测试",
			Methods: []string{"Constructor", "put", "put", "get", "put", "get", "put", "get", "get", "get"},
			Inputs: [][]interface{}{
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
			},
			Expected: []interface{}{nil, nil, nil, 1, nil, -1, nil, -1, 3, 4},
		},
		{
			Name:    "容量为1的缓存",
			Methods: []string{"Constructor", "put", "get", "put", "get", "get"},
			Inputs: [][]interface{}{
				{1},     // Constructor
				{1, 10}, // put
				{1},     // get
				{2, 20}, // put
				{1},     // get
				{2},     // get
			},
			Expected: []interface{}{nil, nil, 10, nil, -1, 20},
		},
		{
			Name:    "更新已存在的键值",
			Methods: []string{"Constructor", "put", "put", "put", "get", "get", "put", "get", "get", "get"},
			Inputs: [][]interface{}{
				{2},      // Constructor
				{1, 10},  // put
				{2, 20},  // put
				{1, 100}, // put (更新)
				{1},      // get
				{2},      // get
				{3, 30},  // put
				{2},      // get
				{1},      // get
				{3},      // get
			},
			Expected: []interface{}{nil, nil, nil, nil, 100, 20, nil, -1, 100, 30},
		},
	}

	runner.RunTestCases(t, testCases)
}

// CreateTestCaseFromLeetCode 从LeetCode格式创建测试用例
func CreateTestCaseFromLeetCode(name string, methods []string, inputs [][]int, expected []interface{}) TestCase {
	tc := TestCase{
		Name:     name,
		Methods:  methods,
		Inputs:   make([][]interface{}, len(inputs)),
		Expected: expected,
	}

	for i, input := range inputs {
		tc.Inputs[i] = make([]interface{}, len(input))
		for j, val := range input {
			tc.Inputs[i][j] = val
		}
	}

	return tc
}

// PrintTestCase 打印测试用例信息（用于调试）
func PrintTestCase(tc TestCase) {
	fmt.Printf("测试用例: %s\n", tc.Name)
	fmt.Printf("方法序列: %v\n", tc.Methods)
	fmt.Printf("输入参数: %v\n", tc.Inputs)
	fmt.Printf("期望输出: %v\n", tc.Expected)
	fmt.Println("---")
}
