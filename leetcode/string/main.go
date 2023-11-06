package main

import (
	"fmt"
	"strings"
)

/*
1、byte和uint8等价。int32和rune等价。后者表征字符。
2、rune的目的是国际化，使非英文字符也能表示，比如中文就是占3个byte，但可以用1个rune来表示。
3、实际应用中，字符串的迭代是按照rune的个数，也就是unicode编码方式去迭代的。
4、rune占内存会多一点，但是不会出现因为编码方式的导致的错误。实际开发中适合
5、rune可以取负值！

*/
// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
func main() {
	s := "23avddfabnmsddfasddd"
	res := lengthOfLongestSubstring1(s)
	fmt.Printf("res:%d", res)

	s = "好家伙"
	sbb := []rune(s)
	fmt.Println(len(sbb))

	//
	s1 := "abcd"
	a := []uint8(s1)
	fmt.Println(a)
	fmt.Println(s1[0])
	fmt.Println(s1[1])

	b := []rune(s1)
	fmt.Println(b)
	fmt.Println(b[0])
	fmt.Println(b[1])

	s2 := "大家好"
	c := []int32(s2)
	fmt.Println(c)
	fmt.Println(c[0])
	fmt.Println(c[1])

	e := []uint8(s2)
	fmt.Println(e)
	fmt.Println(e[0])
	fmt.Println(e[1])

	tmp := e[:3]
	fmt.Println(string(tmp))
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

//利用一个map的特性是否含有相同key 去做判断，然后利用双指针做...
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

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	//使用 charIndex 哈希表来存储每个字符的最后一个出现位置
	charIndex := make(map[byte]int)
	//
	maxLength := 0
	// left 和 right 指针来表示当前的滑动窗口，left 指向窗口的左边界，right 指向窗口的右边界
	left := 0
	//逐步扩展窗口，如果发现重复字符，就将左边界 left 移动到重复字符的下一个位置，以确保窗口内没有重复字符。
	//在每一步中，我们都会更新 maxLength，以记录最长子串的长度。
	for right := 0; right < n; right++ {
		if index, found := charIndex[s[right]]; found {
			left = max(left, index+1)
		}
		charIndex[s[right]] = right
		maxLength = max(maxLength, right-left+1)
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
