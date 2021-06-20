package array

import (
	"fmt"
	"strconv"
)

//给定一个无重复元素的有序整数数组 nums 。
//
//返回 恰好覆盖数组中所有数字 的 最小有序 区间范围列表。也就是说，nums 的每个元素都恰好被某个区间范围所覆盖，并且不存在属于某个范围但不属于 nums 的数字 x 。
//
//列表中的每个区间范围 [a,b] 应该按如下格式输出：
//
//"a->b" ，如果 a != b
//"a" ，如果 a == b
//示例 1：
//
//输入：nums = [0,1,2,4,5,7]
//输出：["0->2","4->5","7"]
//解释：区间范围是：
//[0,2] --> "0->2"
//[4,5] --> "4->5"
//[7,7] --> "7"
//示例 2：
//
//输入：nums = [0,2,3,4,6,8,9]
//输出：["0","2->4","6","8->9"]
//解释：区间范围是：
//[0,0] --> "0"
//[2,4] --> "2->4"
//[6,6] --> "6"
//[8,9] --> "8->9"
//示例 3：
//
//输入：nums = []
//输出：[]
//示例 4：
//
//输入：nums = [-1]
//输出：["-1"]
//示例 5：
//
//输入：nums = [0]
//输出：["0"]
func summaryRanges(nums []int) []string {
	if len(nums) <= 0 {
		return []string{}
	}

	res := make([]string, 0)
	startIdx, endIdx := 0, 0
	for i := 1; i < len(nums); i++ {
		// 如果连续，更新endIdx
		if nums[i] <= nums[endIdx]+1 {
			endIdx = i
			continue
		}
		// 如果不连续
		var rangeStr string
		if nums[endIdx] == nums[startIdx] {
			rangeStr = strconv.Itoa(nums[endIdx])
		} else {
			rangeStr = fmt.Sprintf("%d->%d", nums[startIdx], nums[endIdx])
		}
		res = append(res, rangeStr)
		//	更新idx
		startIdx, endIdx = i, i
	}
	if nums[endIdx] == nums[startIdx] {
		res = append(res, strconv.Itoa(nums[endIdx]))
	} else {
		res = append(res, fmt.Sprintf("%d->%d", nums[startIdx], nums[endIdx]))
	}

	return res
}
