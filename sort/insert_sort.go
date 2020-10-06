package sort

func insertSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	for i := 1; i < len(arr); i++ {
		cur := arr[i]
		for j := i - 1; j >= 0; j-- {
			if cur < arr[j] {
				arr[j+1], arr[j] = arr[j], cur
			} else {
				break
			}
		}
	}
	return arr
}
