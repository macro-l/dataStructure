package main

import "fmt"

// 链表节点结构体
type ListNode struct {
	Val  int       // 节点值
	Next *ListNode // 指向下一节点的指针
}

// 链表构造函数
func NewListNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

// 插入节点
// n0后插入P
func insertNode(n0 *ListNode, P *ListNode) {
	// 获取no下个节点n1
	n1 := n0.Next
	// P下个节点指向n1
	P.Next = n1
	//n0个节点指向P
	n0.Next = P
}

// 删除节点
func removeNode(n0 *ListNode) {
	if n0.Next == nil {
		return
	}

	// 获取下个节点
	P := n0.Next
	// 后个节点
	n1 := P.Next
	// 指向后个节点
	n0.Next = n1

}

// 访问节点
func access(head *ListNode, index int) *ListNode {
	for i := 0; i < index; i++ {
		if head == nil {
			return nil
		}
		head = head.Next
	}
	return head
}

// 查找节点
func findNode(head *ListNode, target int) int {
	index := 0
	for head != nil {
		if head.Val == target {
			return index
		}
		head = head.Next
		index++
	}
	return -1
}

func main() {
	n0 := NewListNode(1)
	n1 := NewListNode(2)
	n2 := NewListNode(3)
	n3 := NewListNode(4)
	n4 := NewListNode(5)
	n5 := NewListNode(6)
	insertNode(n0, n1)
	insertNode(n1, n2)
	insertNode(n2, n3)
	insertNode(n3, n4)
	insertNode(n4, n5)

	x := access(n0, 5)
	y := findNode(n0, 5)

	fmt.Println(x)
	fmt.Println(y)

}
