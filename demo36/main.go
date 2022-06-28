package main

import "fmt"

//n个人,每到m就出圈
func josephus(n int, m int) int {
	if n == 1 {
		return 0
	} else {
		return (josephus(n-1, m) + m) % n
	}
}

func main() {
	n := 3
	m := 2
	res := josephus(n, m)
	fmt.Println(res)

	test(3)
}

func test(n int) {
	var find, count int //find=1找到下一个猴子让它出去,count计数是否数到3
	//fmt.Scan(&n)
	num := 0        //当num到达N-1时，只剩一只猴子
	flag := []int{} //flag[i],该猴子已经出去了

	for i := 0; i < n; i++ {
		flag = append(flag, 0)
	}
	i := 0 //找下一只要出去的猴子从i开始找
	for num != n-1 {
		find = 0
		count = 0
		for ; find == 0; i++ {
			t := i % n
			if flag[t] == 0 {
				count += 1
			}
			if count == 3 {
				num++
				find = 1
				flag[t] = 1
			}
		}
	}
	fmt.Println(flag)
	for i = 0; i < n; i++ {
		if flag[i] == 0 {
			fmt.Println(i)
		}
	}
}
