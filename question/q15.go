package question

import (
	"sort"
)

//
//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
//
//注意：答案中不可以包含重复的三元组。
//
//示例：
//
//给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//
//满足要求的三元组集合为：
//[
//[-1, 0, 1],
//[-1, -1, 2]
//]

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	if n < 3 {
		return [][]int{}
	}
	if nums[0] > 0 {
		return [][]int{}
	}

	ret := make([][]int, 0)
	for i := n - 1; i >= 0 && nums[i] >= 0; i-- {
		if i < n-1 && nums[i] == nums[i+1] {
			continue
		}
		left, right := 0, i-1
		for left < right {
			target := nums[i] * -1
			sum := nums[left] + nums[right]
			if sum > target || (right < i-1 && nums[right] == nums[right+1]) {
				right--
				continue
			}
			if sum < target || (left > 0 && nums[left] == nums[left-1]) {
				left++
				continue
			}
			ret = append(ret, []int{nums[i], nums[left], nums[right]})
			left++
		}
	}
	return ret
}
