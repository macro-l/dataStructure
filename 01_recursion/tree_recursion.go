package main

import "fmt"

// 递归树
// 注：更试适用于分治问题
// 输出斐波那契数列
func treeRecur(n int) int {
	if 2 >= n {
		return n - 1
	}

	// return treeRecur(n-1) + treeRecur(n-2)
	// 分解过程记录
	last := treeRecur(n - 1)
	llast := treeRecur(n - 2)
	fmt.Println()

	res := last + llast
	return res

}

func main() {
	res := treeRecur(10)
	fmt.Println(res)

	// for i := 1; i <= 10; i++ {
	// 	res := treeRecur(i)
	// 	fmt.Println(res)
	// }
}
