package main

// 数组栈

import (
	"fmt"
)

// 基于链表实现的栈
type arrayStack struct {
	// 借用切片
	data []int
}

// 初始化栈
func newArrayStack() *arrayStack {
	return &arrayStack{
		// 设置栈的长度为 0，容量为 16
		data: make([]int, 0, 16),
	}
}

// 栈的长度
func (s *arrayStack) size() int {
	return len(s.data)
}

// 栈是否为空
func (s *arrayStack) isEmpty() bool {
	return s.size() == 0
}

// 入栈
func (s *arrayStack) push(v int) {
	s.data = append(s.data, v)
}

// 出栈
func (s *arrayStack) pop() any {
	val := s.peek()
	s.data = s.data[:len(s.data)-1]
	return val
}

// 获取栈顶元素
func (s *arrayStack) peek() any {
	if s.isEmpty() {
		return nil
	}
	val := s.data[len(s.data)-1]
	return val
}

// 获取slick 用于打印
func (s *arrayStack) toSlice() []int {
	return s.data
}

// 打印
func (s arrayStack) echo() {
	fmt.Println(s.toSlice())
}

func main() {
	s := newArrayStack()
	s.push(3)
	s.push(1)
	s.push(2)
	s.push(4)
	s.push(5)
	s.push(6)
	s.echo()

	top := s.pop()
	fmt.Printf("top:%d\n", top)
	s.echo()

	peek := s.peek()
	fmt.Printf("peek:%d\n", peek)
	s.echo()

}
