package interview

//给定二维数组，横竖都是递增的整数，判断是否存在某个整数

func IsInArray(arr [][]int, target int) bool {
	for row := range arr {
		if arr[row][0] > target {
			return false
		}
		l := len(arr[row])
		if arr[row][l-1] < target {
			continue
		}
		for _, val := range arr[row] {
			if val > target {
				break
			}
			if val == target {
				return true
			}
		}
	}
	return false
}
