package leetcode

import (
	"fmt"
	"strings"
)

// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
func main() {
	s := "23addfasddfasddd"
	res := lengthOfLongestSubstring1(s)
	fmt.Printf("res;%d", res)
}
func lengthOfLongestSubstring1(s string) int {
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		index := strings.Index(s[start:i], string(s[i]))
		if index == -1 {
			if i+1 > end {
				end = i + 1
			}
		} else {
			start += index + 1
			end += index + 1
		}
	}
	return end - start
}

//链接：https://leetcode.cn/problems/longest-substring-without-repeating-characters/solution/zhi-zeng-da-bu-jian-xiao-de-hua-dong-chuang-kou-10/

func lengthOfLongestSubstring2(s string) int {
	block := make(map[uint8]int, len(s))
	max := 0
	j := -1
	for i := 0; i < len(s); i++ {
		if _, ok := block[s[i]]; ok {
			if block[s[i]] >= j {
				j = block[s[i]]
			}
			block[s[i]] = i
		} else {
			block[s[i]] = i
		}
		if max < i-j {
			max = i - j
		}
	}
	return max
}

//链接：https://leetcode.cn/problems/longest-substring-without-repeating-characters/solution/zhe-shi-ge-te-shu-de-hua-dong-chuang-kou-5huk/
