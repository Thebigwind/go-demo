package common

import (
	"fmt"
	"strings"
)

// 求出交集
func Intersection(a []int, b []int) (inter []int) {
	m := make(map[int]int)
	nn := make([]int, 0)
	for _, v := range a {
		m[v] = v
	}
	for _, v := range b {
		times := m[v]
		if times > 0 {
			nn = append(nn, v)
		}
	}
	return nn
}

// 求差集
func SupplementarySet(slice1, slice2 []string) []string {
	m := make(map[string]string)
	for _, v := range slice1 {
		m[v] = v
	}
	for _, v := range slice2 {
		if m[v] != "" {
			delete(m, v)
		}
	}
	var str []string
	for _, s2 := range m {
		str = append(str, s2)
	}
	return str
}

// StringBuild 根据输入的参数输出一个拼接好的字符串
func StringBuild(args ...string) string {
	var builder strings.Builder
	for _, str := range args {
		builder.WriteString(str)
	}
	return builder.String()
}

// SliceIndexOf 返回 elem 在切片 haystack 中的位置下标，如果不存在则返回-1
func SliceIndexOf(haystack []string, elem string) int {
	for k, v := range haystack {
		if v == elem {
			return k
		}
	}
	return -1
}

// ContainString 返回 target 是否在 arr 中出现
func ContainString(arr []string, target string) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}

// RemoveRepByLoop 通过两重循环过滤重复元素
func RemoveRepByLoop(slc []int64) []int64 {
	result := make([]int64, 0) // 存放结果
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				flag = false // 存在重复元素，标识为false
				break
			}
		}
		if flag { // 标识为false，不添加进结果
			result = append(result, slc[i])
		}
	}
	return result
}

// RemoveRepByMap 通过map主键唯一的特性过滤重复元素
func RemoveRepByMap(slc []int64) []int64 {
	result := make([]int64, 0)
	tempMap := map[int64]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	return result
}

// RemoveRep 元素去重
func RemoveRep(slc []int64) []int64 {
	if len(slc) < 1024 {
		// 切片长度小于1024的时候，循环来过滤
		return RemoveRepByLoop(slc)
	} else {
		// 大于的时候，通过map来过滤
		return RemoveRepByMap(slc)
	}
}

func CompareSlice(slice1, slice2 []string) error {
	if len(slice1) != len(slice2) {
		return fmt.Errorf("slice num not equal")
	}
	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return fmt.Errorf("slice index %d not equal", i)
		}
	}
	return nil
}
