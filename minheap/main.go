package main

import (
	"container/heap"
	"fmt"
)

// 定义一个最小堆，使用[]int作为底层数据结构
type MinHeap []int

// 实现heap.Interface中的Len方法
func (h MinHeap) Len() int { return len(h) }

// 实现heap.Interface中的Less方法，确定堆的顺序
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }

// 实现heap.Interface中的Swap方法，交换堆中元素的位置
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// 实现heap.Interface中的Push方法，将元素添加到堆中
func (h *MinHeap) Push(x interface{}) {
	// Push方法接收的是interface{}类型，需要转换为int类型
	*h = append(*h, x.(int))
}

// 实现heap.Interface中的Pop方法，移除堆顶元素
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 测试最小堆的用法
func main() {
	// 创建一个最小堆
	h := &MinHeap{}

	// 使用heap.Interface中定义的方法操作最小堆
	heap.Init(h)
	heap.Push(h, 5)
	heap.Push(h, 1)
	heap.Push(h, 3)
	heap.Push(h, 9)
	heap.Push(h, 4)
	heap.Push(h, 7)
	heap.Push(h, 9)
	fmt.Printf("最小元素: %d\n", (*h)[0])
	fmt.Printf("最小元素: %d\n", h)

	// 取出最小元素
	for h.Len() > 0 {
		fmt.Printf("弹出元素: %d\n", heap.Pop(h))
	}
}
