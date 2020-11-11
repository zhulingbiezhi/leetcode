package interview

//有面额为1、4、5的三种硬币，输出组成金额n需要的最少硬币数
func MinCoin(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		min := 0
		for _, val := range []int{1, 4, 5} {
			if val > i {
				break
			}
			dp[val] = 1
			cnt := dp[val] + dp[i-val]
			if cnt < min || min == 0 {
				min = cnt
			}
		}
		dp[i] = min
	}
	return dp[n]
}
