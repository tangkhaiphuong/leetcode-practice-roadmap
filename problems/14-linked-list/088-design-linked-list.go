//go:build ignore

// Problem 088: Design Linked List (LeetCode #707)
// Difficulty: Medium
// Categories: Linked List, Design
//
// Description:
//   Implement a singly linked list with these methods:
//     Get(index)            -> value at index, or -1 if invalid.
//     AddAtHead(val)
//     AddAtTail(val)
//     AddAtIndex(index, val) -> insert before existing index.
//     DeleteAtIndex(index)
//
// Approach: Sentinel-Headed Singly Linked List
//   A sentinel head simplifies AddAtIndex(0). Track size for bounds checking.
//
// Complexity:
//   Get/AddAtIndex/DeleteAtIndex: O(index)
//   AddAtHead: O(1)
//   AddAtTail: O(n) without tail pointer; O(1) if tracked.

package main

import "fmt"

type listNode struct {
	val  int
	next *listNode
}

type MyLinkedList struct {
	head *listNode
	size int
}

func ConstructorMLL() MyLinkedList { return MyLinkedList{head: &listNode{}} }

func (l *MyLinkedList) Get(index int) int {
	if index < 0 || index >= l.size {
		return -1
	}
	cur := l.head.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur.val
}

func (l *MyLinkedList) AddAtIndex(index, val int) {
	if index < 0 {
		index = 0
	}
	if index > l.size {
		return
	}
	cur := l.head
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	cur.next = &listNode{val: val, next: cur.next}
	l.size++
}

func (l *MyLinkedList) AddAtHead(val int) { l.AddAtIndex(0, val) }
func (l *MyLinkedList) AddAtTail(val int) { l.AddAtIndex(l.size, val) }

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= l.size {
		return
	}
	cur := l.head
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	cur.next = cur.next.next
	l.size--
}

func main() {
	l := ConstructorMLL()
	l.AddAtHead(1)
	l.AddAtTail(3)
	l.AddAtIndex(1, 2)
	fmt.Println(l.Get(1)) // 2
	l.DeleteAtIndex(1)
	fmt.Println(l.Get(1)) // 3
}
