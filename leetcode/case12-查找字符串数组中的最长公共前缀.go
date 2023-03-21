package leetcode

//编写一个函数来查找字符串数组中的最长公共前缀。
//如果不存在公共前缀，返回空字符串 ""。
/*
思路:化繁为简，多个字符串的最长公共前缀不好求，两个字符串的最长公共前缀是很容易求得的，那么我们可以先求出字符串数组中前两个字符串的最长公共前缀prefix, 然后在遍历字符串数组strs时，迭代这个prefix就好了，也就是求prefix和下一个字符串strs[i]的最长公共前缀。特别的，如果循环中，prefix长度为0，说明strs[0:i]范围内的所有字符串最长公共前缀为空串，后续的遍历也就没有意义了，直接break退出循环。当然，还需要考虑特殊情况，如果字符串数组的长度为0，直接返回空串。

作者：ybzdqhl
链接：https://leetcode.cn/problems/longest-common-prefix/solution/zui-zhi-guan-de-jie-fa-by-ybzdqhl-70do/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func LongestCommonPrefix(strs []string) string {
	count := len(strs)
	if count == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < count; i++ {
		prefix = lcp(prefix, strs[i])
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}

// lcp 求两个字符串str1和str2的最长公共前缀
func lcp(str1, str2 string) string {
	length := Min(len(str1), len(str2))
	index := 0
	for index < length && str1[index] == str2[index] {
		index++
	}
	return str1[:index]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

///////////////

//工作中这种函数就怕来一个nil或[]string{}给你搞出一个死循环来。
//这里，我们按题设知道不用做“排错”，所以，第一行的strs[0]如果越界panic相当于一个自排错，好过死循环。

func longestCommonPrefix2(strs []string) string {
	base := strs[0] //1)panic自排错不会死循环
	for i := 0; i < len(base); i++ {
		for _, str := range strs[1:] {
			if i >= len(str) || str[i] != base[i] {
				return base[:i]
			}
		}
	}
	return base
}

//链接：https://leetcode.cn/problems/longest-common-prefix/solution/by-pandaoknight-wmix/ 权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

/*
解题思路
先找到最短的元素，因为最长公共前缀的长度肯定小于等于这个元素的长度
接着遍历最短的元素，和数组中的每个元素的相同位置比较看是否相同
如果遇到不同的了，那就返回不同的之前的所有元素
如果遍历完了这个最短元素，说明这个最短元素就是最长公共前缀，直接返回



作者：DCCooper
链接：https://leetcode.cn/problems/longest-common-prefix/solution/gojie-fa-xiang-xi-zhu-shi-kan-bu-dong-da-si-wo-by-/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func longestCommonPrefix(strs []string) string {
	// 最长的公共前缀的长度肯定小于等于数组中最短的元素
	// 所以从这个元素开始当基准
	short := findShortestInArray(strs)
	// 早发现早治疗
	if len(short) == 0 {
		return ""
	}
	// 遍历这个最短的每个位置元素，用来判断是不是相等
	for i, v := range short {
		// 要判断多少次，取决于数组strs中有多少个元素，所以用的len(strs)
		for j := 0; j < len(strs); j++ {
			// 数组的第j个元素的第i个位置不等于我们的short的第i个位置的元素
			// 写成strs[j][i] 是为了和short里面的每个元素一一对应比较
			if strs[j][i] != byte(v) {
				// 到了第[j][i]个没有匹配上，那么就说明之前的都匹配上了，所以直接返回[j][:i]
				return strs[j][:i]
			}
		}
	}
	// 遍历完short了，说明short就是最长的，直接返回
	return short
}

func findShortestInArray(s []string) string {
	// 空字符数组返回空
	if len(s) == 0 {
		return ""
	}
	// 临时定义最短为数组第一个
	shortest := s[0]
	// 遍历数组每个元素
	for _, v := range s {
		// 找到当前小于res
		if len(v) < len(shortest) {
			// 看看是否是空的，空的说明数组中有空字符，所以最长公共前缀肯定为空
			if len(v) == 0 {
				return ""
			}
			// 替换当前最小为当前遍历到的元素
			shortest = v
		}
	}
	return shortest
}
