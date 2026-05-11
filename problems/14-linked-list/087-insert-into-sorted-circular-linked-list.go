//go:build ignore

// Problem 087: Insert into a Sorted Circular Linked List (LeetCode #708)
// Difficulty: Medium
// Categories: Linked List
//
// Description:
//   Given a sorted (non-decreasing) circular singly linked list. Insert a
//   value preserving the order. Return any node of the resulting list. If
//   the list is empty (head==nil), create a single-node list pointing to
//   itself.
//
// Approach: Single Walk, Three Cases
//   Walk from head; track each (cur, cur.Next). Insert if any of:
//     - cur.Val <= x <= cur.Next.Val   (between two nodes)
//     - cur.Val > cur.Next.Val (the wraparound boundary) AND
//         (x >= cur.Val OR x <= cur.Next.Val)
//     - we returned to head (uniform list) — insert anywhere.
//
// Complexity:
//   Time:  O(n)
//   Space: O(1)

package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func insert(head *Node, val int) *Node {
	n := &Node{Val: val}
	if head == nil {
		n.Next = n
		return n
	}
	cur := head
	for {
		nxt := cur.Next
		if cur.Val <= val && val <= nxt.Val {
			break
		}
		if cur.Val > nxt.Val && (val >= cur.Val || val <= nxt.Val) {
			break
		}
		cur = nxt
		if cur == head {
			break // looped through uniform list
		}
	}
	n.Next = cur.Next
	cur.Next = n
	return head
}

func main() {
	// Build [3,4,1] circular
	a := &Node{Val: 3}
	b := &Node{Val: 4}
	c := &Node{Val: 1}
	a.Next = b
	b.Next = c
	c.Next = a
	insert(a, 2)
	cur := a
	for i := 0; i < 4; i++ {
		fmt.Printf("%d ", cur.Val)
		cur = cur.Next
	}
	fmt.Println()
}
