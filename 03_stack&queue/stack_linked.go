package main

// 链表栈

import (
	"container/list"
	"fmt"
)

// 基于链表实现的栈
type linkedListStack struct {
	// 借用list的节点结构
	data *list.List
}

// 初始化栈
func newLinkedListStack() *linkedListStack {
	return &linkedListStack{
		data: list.New(),
	}
}

// 入栈
func (s *linkedListStack) push(value int) {
	s.data.PushBack(value)
}

// 出栈
func (s *linkedListStack) pop() any {
	if s.isEmpty() {
		return nil
	}
	e := s.data.Back()
	s.data.Remove(e)
	return e.Value
}

// 访问栈顶元素
func (s *linkedListStack) peek() any {
	if s.isEmpty() {
		return nil
	}
	e := s.data.Back()
	return e.Value
}

// 获取栈的长度
func (s *linkedListStack) size() int {
	return s.data.Len()
}

// 判断栈是否为空
func (s *linkedListStack) isEmpty() bool {
	return s.data.Len() == 0
}

// 获取list
func (s *linkedListStack) toList() *list.List {
	return s.data
}

// 输出栈
func (s *linkedListStack) echo() {
	d := s.data.Front()
	fmt.Printf("<bottom> ")
	for d != nil {
		fmt.Printf("%d ", d.Value)
		d = d.Next()
	}
	fmt.Printf(" <top>")
	fmt.Printf("\n")
}

func main() {
	s := newLinkedListStack()
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
