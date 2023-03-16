package leetcode

//最长回文子串
func longestPalindrome(s string) string {
	LEN := len(s)
	if LEN <= 1 {
		return s
	}

	maxLen := 1
	begin := 0

	dp := make([][]bool, LEN)
	for i := range dp {
		dp[i] = make([]bool, LEN)
		dp[i][i] = true
	}

	for j := 1; j < LEN; j++ { // 从最左列, 到最右列, 的方向
		for i := 0; i < j; i++ { // 对这一列：从"第一行"到"对角线那一行"
			if s[i] != s[j] { // 首尾不相等，则不是回文
				dp[i][j] = false
			} else {
				if j-i < 3 { // 去除首尾后，只剩0或1个元素，则是回文
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1] // 根据去除首尾后剩下的部分，是否为回文
				}
			}

			if dp[i][j] && j-i+1 > maxLen { // 满足题意，的最长，记录
				maxLen = j - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+maxLen]
}

//链接：https://leetcode.cn/problems/longest-palindromic-substring/solution/by-wwsw-fwi1/

func longestPalindrome2(s string) string {
	max := ""
	for o := len(s) - 1; o >= 0; o-- {
		i := 0
		j := o
		for k := j; k <= len(s)-1; k++ {
			a := findaa(s, i+k-j, k)
			if a && len(max) < j-i+1 {
				max = s[i+k-j : k+1]
				return max
			}
		}
	}
	return max
}
func findaa(s string, i int, j int) bool {
	for i <= j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

//链接：https://leetcode.cn/problems/longest-palindromic-substring/solution/cong-da-ju-ru-shou-kan-zui-chang-hui-wen-3z0h/
