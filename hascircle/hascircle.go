package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func isCircle(head *Node) bool {
	if head == nil {
		return false
	}
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

func main() {

	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	fmt.Println(isCircle(node1))

	node1 = &Node{Val: 1}
	node2 = &Node{Val: 2}
	node3 = &Node{Val: 3}
	node4 = &Node{Val: 4}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node3
	fmt.Println(isCircle(node1))
}
