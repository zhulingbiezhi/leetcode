package sort

//冒泡排序是一种简单的排序算法。
//它重复地走访过要排序的数列，一次比较两个元素，如果它们的顺序错误就把它们交换过来。
//走访数列的工作是重复地进行直到没有再需要交换，也就是说该数列已经排序完成。
//这个算法的名字由来是因为越小的元素会经由交换慢慢“浮”到数列的顶端。

func bubbleSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	l := len(arr)
	for i := 0; i < l-1; i++ {
		for j := 0; j < l-1-i; j++ {
			if arr[j] > arr[j+1] { // 相邻元素两两对比
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
	return arr
}
