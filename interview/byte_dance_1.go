package interview

//给一个正整数n，输出用3进制表示后，从右往左数第1、2、3、5、8、... 位的数字
//
//比如输入n=5，3进制表示为12，从右往左数第一位是2，第二位是1，对应的输出为21

func NumberToThreeJinZhi(n int) []int {
	ret := make([]int, 0)
	result := make([]int, 0)
	m := n
	for m > 0 {
		r := m % 3
		m = m / 3
		result = append(result, r)
	}
	arr := []int{0, 1}
	index := 0
	for i := len(result) - 1; i >= 0; i-- {
		index++
		if MatchNumber(arr, index) {
			ret = append([]int{result[i]}, ret...)
		}
	}

	return ret
}

func MatchNumber(arr []int, n int) bool {
	if len(arr) < 2 {
		return false
	}
	if n == arr[0]+arr[1] {
		arr[0], arr[1] = arr[1], n
		return true
	}
	return false
}
