package sort

func mergeSort(arr []int) []int {
	l := len(arr)
	if l == 2 {
		if arr[0] > arr[1] {
			arr[0], arr[1] = arr[1], arr[0]
		}
		return arr
	}
	if len(arr) == 1 {
		return arr
	}
	m := len(arr) / 2
	return merge(mergeSort(arr[:m]), mergeSort(arr[m:]))
}

func merge(left, right []int) []int {
	m := len(left)
	n := len(right)
	i, j := 0, 0
	ret := make([]int, 0, m+n)
	for i < m && j < n {
		if left[i] < right[j] {
			ret = append(ret, left[i])
			i++
		} else {
			ret = append(ret, right[j])
			j++
		}
	}
	for ; j < n; j++ {
		ret = append(ret, right[j])
	}
	for ; i < m; i++ {
		ret = append(ret, left[i])
	}
	return ret
}
