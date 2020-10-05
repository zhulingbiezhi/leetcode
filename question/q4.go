package question

//
//给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的中位数。
//
//进阶：你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗？
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays

//暴力解法
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	odd := (m+n)%2 != 0

	total := (m + n + 1) / 2
	if !odd {
		total += 1
	}

	i, j := 0, 0
	ret := make([]int, 0, total)
	for len(ret) < total && i < m && j < n {
		if nums1[i] > nums2[j] {
			ret = append(ret, nums2[j])
			j++
		} else {
			ret = append(ret, nums1[i])
			i++
		}
	}
	for len(ret) < total && i < m {
		ret = append(ret, nums1[i])
		i++
	}
	for len(ret) < total && j < n {
		ret = append(ret, nums2[j])
		j++
	}
	if !odd {
		return float64(ret[total-1]+ret[total-2]) / 2
	}
	return float64(ret[total-1])
}
