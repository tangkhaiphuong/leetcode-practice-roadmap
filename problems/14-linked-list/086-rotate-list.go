//go:build ignore

// Problem 086: Rotate List (LeetCode #61)
// Difficulty: Medium
// Categories: Linked List, Two Pointers
//
// Description:
//   Rotate a singly linked list to the right by k places.
//
// Examples:
//   1->2->3->4->5, k=2  ->  4->5->1->2->3
//   0->1->2,        k=4 ->  2->0->1
//
// Approach: Make Circular, Find New Tail
//   Compute length n. Connect tail to head (circular). New tail is the
//   (n - k%n - 1)-th node from start. New head = newTail.next. Break the
//   cycle by setting newTail.next = nil.
//
// Complexity:
//   Time:  O(n)
//   Space: O(1)

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	n := 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		n++
	}
	k %= n
	if k == 0 {
		return head
	}
	tail.Next = head // circular
	steps := n - k - 1
	newTail := head
	for i := 0; i < steps; i++ {
		newTail = newTail.Next
	}
	newHead := newTail.Next
	newTail.Next = nil
	return newHead
}

func build(vals ...int) *ListNode {
	d := &ListNode{}
	cur := d
	for _, v := range vals {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return d.Next
}
func print(h *ListNode) {
	for h != nil {
		fmt.Printf("%d ", h.Val)
		h = h.Next
	}
	fmt.Println()
}

func main() {
	print(rotateRight(build(1, 2, 3, 4, 5), 2))
	print(rotateRight(build(0, 1, 2), 4))
}
