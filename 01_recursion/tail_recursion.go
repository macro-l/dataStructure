package main

import "fmt"

// 尾递归
// 反序输出start到end
func tailRecur(n int, res int) int {
	fmt.Printf("叠乘%d阶段开始\n", n)

	if 0 > n {
		return 0
	}
	if 1 == n {
		fmt.Printf("叠成1阶段无需计算")
		fmt.Printf("叠成1阶段直接返回1")
		return 1
	}

	fmt.Printf("叠成%d阶段计算过程： %d * %d\n", n, n, res)
	fmt.Printf("叠成%d阶段返回结果： %d \n", n, res*n)
	fmt.Printf("%d * %d = %d\n", n, res, res*n)
	return tailRecur(n-1, res*n)
}

func main() {
	tailRecur(10, 1)
}
