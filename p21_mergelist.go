package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	head := l3
	for l1 != nil || l2 != nil {
		if l1 == nil {
			l3.Next = l2
			return head.Next
		}
		if l2 == nil {
			l3.Next = l1
			return head.Next
		}
		if l1.Val < l2.Val {
			l3.Next = l1
			l1 = l1.Next
			l3.Next.Next = l2
			l2 = l2.Next
		} else {
			l3.Next = l2
			l2 = l2.Next
			l3.Next.Next = l1
			l1 = l1.Next
		}
		l3 = l3.Next.Next
	}
	return head.Next
}

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	head := l3
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			l3.Next = l1
			l1 = l1.Next
		} else {
			l3.Next = l2
			l2 = l2.Next
		}
		l3 = l3.Next
	}
	if l1 != nil {
		l3.Next = l1
	}
	if l2 != nil {
		l3.Next = l2
	}
	return head.Next
}
func main() {
	l1 := &ListNode{1, &ListNode{2, &ListNode{6, nil}}}
	l2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	// l1 := &ListNode{5, nil}
	// l2 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	l3 := mergeTwoLists2(l1, l2)
	for l3 != nil {
		fmt.Println(l3.Val)
		l3 = l3.Next
	}
}
