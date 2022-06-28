package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(shuffle(arr1))
}

// 每次随机挑选一个值，放在数组末尾。然后在n-1个元素的数组中再随机挑选一个值，放在数组末尾，以此类推
func shuffle(nums []int) []int {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := len(nums); i > 0; i-- {
		last := i - 1
		idx := rand.Intn(i)
		nums[last], nums[idx] = nums[idx], nums[last]
	}
	return nums
}
