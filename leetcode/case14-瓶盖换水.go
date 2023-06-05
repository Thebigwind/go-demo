package leetcode

import "fmt"

func main() {
	// 矿泉水单价
	price := 2.0
	// 瓶盖换水比例
	capsRatio := 4
	// 空瓶换水比例
	bottlesRatio := 2

	// 初始瓶数
	bottles := int(100 / price)

	// 初始瓶盖数和空瓶数
	caps := bottles
	emptyBottles := bottles

	// 统计换水次数
	var count int
	for caps >= capsRatio || emptyBottles >= bottlesRatio {
		// 计算可换的瓶数
		capsExchange := caps / capsRatio
		emptyBottlesExchange := emptyBottles / bottlesRatio

		// 计算换到的新瓶数
		newBottles := capsExchange + emptyBottlesExchange

		// 计算剩余瓶盖数和空瓶数
		caps = caps - capsExchange*capsRatio + newBottles
		emptyBottles = emptyBottles - emptyBottlesExchange*bottlesRatio + newBottles

		// 统计换水次数和新瓶数
		count += capsExchange + emptyBottlesExchange
		bottles += newBottles
	}

	fmt.Printf("100元钱可以喝 %d 瓶矿泉水\n", bottles)
}
