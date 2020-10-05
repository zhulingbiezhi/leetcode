package question

//给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
//
//示例 1：
//
//输入: "babad"
//输出: "bab"
//注意: "aba" 也是一个有效答案。
//示例 2：
//
//输入: "cbbd"
//输出: "bb"
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-palindromic-substring
//

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	ret := ""
	maxLen := 0
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}
	for right := 1; right < len(s); right++ {
		for left := 0; left < len(s); left++ {
			//length == 2时，aba肯定是对称的
			//length == 1时，aa肯定是对称的
			//length > 2时，dp[i,j] = (dp[i+1,j−1] && S[i]==S[j])
			dp[left][right] = (right-left <= 2 || dp[left+1][right-1]) && s[left] == s[right]
			curLen := right - left + 1
			if dp[left][right] && curLen > maxLen {
				maxLen = curLen
				ret = s[left : right+1]
			}
		}
	}
	return ret
}
