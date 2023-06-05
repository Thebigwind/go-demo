package leetcode

//是将数字本身反转，然后将反转后的数字与原始数字进行比较，如果它们是相同的，那么这个数字就是回文
func isPalindrome(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) { // 小于0或者为只有一个数
		return false
	}
	num, res := x, 0
	for num > 0 {
		res = res*10 + num%10
		num = num / 10
	}
	return x == res
}
func isPalindrome2(x int) bool {
	// 倒序后  判断是不是和原来的数字相等
	if x < 0 {
		return false
	}
	origin := x
	redirect := 0
	for x != 0 {
		redirect = redirect*10 + x%10
		x /= 10
	}
	return origin == redirect
}
