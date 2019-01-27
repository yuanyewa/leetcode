package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	var c *ListNode
	for c = head; c != nil && c.Next != nil; {
		if c.Val == c.Next.Val {
			c.Next = c.Next.Next
		} else {
			c = c.Next
		}
	}
	return head
}

func main() {
	var l0, l1, l2, l3, l4 ListNode
	l0.Val = 1
	l0.Next = &l1
	l1.Val = 1
	l1.Next = &l2
	l2.Val = 1
	l2.Next = &l3
	l3.Val = 4
	l3.Next = &l4
	l4.Val = 4
	l4.Next = nil
	c := deleteDuplicates(&l0)
	for {
		fmt.Println(c.Val)
		c = c.Next
		if c == nil {
			break
		}
	}
}
