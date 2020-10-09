package sort

func quickSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	if len(arr) == 2 {
		if arr[0] > arr[1] {
			arr[0], arr[1] = arr[1], arr[0]
		}
		return arr
	}
	left, right := 0, len(arr)-1
	pivotIndex := len(arr) / 2
	pivot := arr[pivotIndex]
	for left < right && left < len(arr) && right >= 0 {
		if arr[left] < pivot {
			left++
			continue
		}
		if arr[right] > pivot {
			right--
			continue
		}
		arr[left], arr[right] = arr[right], arr[left]
	}
	quickSort(arr[:left])
	quickSort(arr[left:])
	return arr
}
