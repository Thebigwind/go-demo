package leetcode

func reverse(x int) int {
	var res int
	for x != 0 {
		if temp := int32(res); (temp*10)/10 != temp {
			return 0
		}
		res = res*10 + x%10
		x = x / 10
	}
	return res
}

//链接：https://leetcode.cn/problems/reverse-integer/solution/shuang-100ji-bai-yong-hu-gan-shou-goyu-yan-zhi-you/
