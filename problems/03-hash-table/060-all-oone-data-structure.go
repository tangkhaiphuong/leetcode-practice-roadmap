//go:build ignore

// Problem 060: All O`one Data Structure (LeetCode #432)
// Difficulty: Hard
// Categories: Hash Table, Linked List, Design
//
// Description:
//   Design a data structure with all four operations in O(1):
//     Inc(key)      - increase key's count by 1; insert at 1 if absent.
//     Dec(key)      - decrease key's count by 1; remove if count drops to 0.
//                     If absent, no-op.
//     GetMaxKey()   - return any key with the max count, "" if empty.
//     GetMinKey()   - return any key with the min count, "" if empty.
//
// Approach: Doubly Linked List of Buckets, Each Bucket = Set of Keys with Same Count
//   Map key -> bucket. Buckets ordered by count ascending. Inc moves key to
//   bucket[count+1] (creating one if necessary). Dec moves to bucket[count-1].
//   Empty buckets get unlinked. Min = head.next; Max = tail.prev.
//
// Complexity:
//   All ops O(1).

package main

import "fmt"

type bucket struct {
	count      int
	keys       map[string]struct{}
	prev, next *bucket
}

type AllOne struct {
	head, tail *bucket
	keyToBkt   map[string]*bucket
}

func ConstructorAllOne() AllOne {
	h := &bucket{count: 0}
	t := &bucket{count: 0}
	h.next = t
	t.prev = h
	return AllOne{head: h, tail: t, keyToBkt: map[string]*bucket{}}
}

func (a *AllOne) insertAfter(prev *bucket, count int) *bucket {
	b := &bucket{count: count, keys: map[string]struct{}{}}
	b.prev = prev
	b.next = prev.next
	prev.next.prev = b
	prev.next = b
	return b
}

func (a *AllOne) removeBucket(b *bucket) {
	b.prev.next = b.next
	b.next.prev = b.prev
}

func (a *AllOne) Inc(key string) {
	bkt, ok := a.keyToBkt[key]
	if !ok {
		// New key, count 1
		first := a.head.next
		var target *bucket
		if first != a.tail && first.count == 1 {
			target = first
		} else {
			target = a.insertAfter(a.head, 1)
		}
		target.keys[key] = struct{}{}
		a.keyToBkt[key] = target
		return
	}
	next := bkt.next
	var target *bucket
	if next != a.tail && next.count == bkt.count+1 {
		target = next
	} else {
		target = a.insertAfter(bkt, bkt.count+1)
	}
	target.keys[key] = struct{}{}
	a.keyToBkt[key] = target
	delete(bkt.keys, key)
	if len(bkt.keys) == 0 {
		a.removeBucket(bkt)
	}
}

func (a *AllOne) Dec(key string) {
	bkt, ok := a.keyToBkt[key]
	if !ok {
		return
	}
	delete(bkt.keys, key)
	if bkt.count == 1 {
		delete(a.keyToBkt, key)
	} else {
		prev := bkt.prev
		var target *bucket
		if prev != a.head && prev.count == bkt.count-1 {
			target = prev
		} else {
			target = a.insertAfter(prev, bkt.count-1)
		}
		target.keys[key] = struct{}{}
		a.keyToBkt[key] = target
	}
	if len(bkt.keys) == 0 {
		a.removeBucket(bkt)
	}
}

func (a *AllOne) GetMaxKey() string {
	if a.tail.prev == a.head {
		return ""
	}
	for k := range a.tail.prev.keys {
		return k
	}
	return ""
}
func (a *AllOne) GetMinKey() string {
	if a.head.next == a.tail {
		return ""
	}
	for k := range a.head.next.keys {
		return k
	}
	return ""
}

func main() {
	a := ConstructorAllOne()
	a.Inc("hello")
	a.Inc("hello")
	fmt.Println(a.GetMaxKey(), a.GetMinKey()) // hello hello
	a.Inc("leet")
	fmt.Println(a.GetMaxKey(), a.GetMinKey()) // hello leet
	a.Dec("hello")
	fmt.Println(a.GetMaxKey(), a.GetMinKey()) // hello leet (or leet hello, both count 1)
}
