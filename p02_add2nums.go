package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	add1 := 0
	l3 := &ListNode{}
	head := l3
	var v1, v2 int
	for {
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		} else {
			v1 = 0
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		} else {
			v2 = 0
		}
		s := v1 + v2 + add1
		if s == 0 && l1 == nil && l2 == nil {
			l3 = nil
			break
		}
		l3.Val = s % 10
		add1 = s / 10
		if l1 == nil && l2 == nil && add1 == 0 {
			break
		}
		l3.Next = &ListNode{}
		l3 = l3.Next
	}
	return head
}
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	curr := l3
	carry := 0
	var v1, v2 int
	for l1 != nil || l2 != nil {
		if l1 != nil {
			v1 = l1.Val
		} else {
			v1 = 0
		}
		if l2 != nil {
			v2 = l2.Val
		} else {
			v2 = 0
		}
		sum := v1 + v2 + carry
		carry /= 10
		

	}

}

// public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
//     ListNode dummyHead = new ListNode(0);
//     ListNode p = l1, q = l2, curr = dummyHead;
//     int carry = 0;
//     while (p != null || q != null) {
//         int x = (p != null) ? p.val : 0;
//         int y = (q != null) ? q.val : 0;
//         int sum = carry + x + y;
//         carry = sum / 10;
//         curr.next = new ListNode(sum % 10);
//         curr = curr.next;
//         if (p != null) p = p.next;
//         if (q != null) q = q.next;
//     }
//     if (carry > 0) {
//         curr.next = new ListNode(carry);
//     }
//     return dummyHead.next;
// }

func main() {
	// l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}}
	// l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: &ListNode{Val: 9}}}}
	l1 := &ListNode{Val: 0}
	l2 := &ListNode{Val: 0, Next: &ListNode{Val: 6}}
	l3 := addTwoNumbers(l1, l2)
	for {
		fmt.Println(l3.Val)
		if l3.Next == nil {
			break
		} else {
			l3 = l3.Next
		}
	}
}
