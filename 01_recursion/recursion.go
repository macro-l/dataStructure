package main

import (
	"fmt"
)

// 普通递归
// 阶乘计算
func recur(n int) int {
	fmt.Printf("叠乘%d阶段开始\n", n)

	if 0 > n {
		return 0
	}
	if 1 == n {
		fmt.Printf("叠成1阶段无需计算")
		fmt.Printf("叠成1阶段直接返回1")
		return 1
	}

	// return n * recur(n - 1)
	// 分解过程记录
	sub := recur(n - 1)
	res := n * sub
	fmt.Printf("叠成%d阶段计算过程： %d * %d\n", n, n, sub)
	fmt.Printf("叠成%d阶段返回结果： %d \n", n, res)
	return res
}

func main() {
	recur(10)
}
