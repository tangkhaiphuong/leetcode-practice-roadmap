//go:build ignore

// Problem 038: Maximum Number of Events That Can Be Attended (LeetCode #1353)
// Difficulty: Medium
// Categories: Array, Greedy, Heap, Sorting
//
// Description:
//   You have events[i] = [start, end]. You can attend at most one event per
//   day, on any single day in [start, end]. Return the max number of events
//   you can attend.
//
// Examples:
//   [[1,2],[2,3],[3,4]]      -> 3
//   [[1,2],[2,3],[3,4],[1,2]]-> 4
//
// Approach: Sort by Start + Min-Heap on End Day
//   Iterate day d from 1..maxEnd. At each d:
//     - push every event whose start == d into a min-heap keyed by end.
//     - pop expired events (end < d).
//     - if heap non-empty, attend the event with smallest end (greedy: it
//       has fewest remaining options), pop it, count++.
//
// Complexity:
//   Time:  O((n + maxDay) log n)
//   Space: O(n)

package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type intHeap []int

func (h intHeap) Len() int            { return len(h) }
func (h intHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *intHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *intHeap) Pop() interface{}   { old := *h; n := len(old); x := old[n-1]; *h = old[:n-1]; return x }

func maxEvents(events [][]int) int {
	sort.Slice(events, func(i, j int) bool { return events[i][0] < events[j][0] })
	h := &intHeap{}
	res, i := 0, 0
	maxDay := 0
	for _, e := range events {
		if e[1] > maxDay {
			maxDay = e[1]
		}
	}
	for d := 1; d <= maxDay; d++ {
		for i < len(events) && events[i][0] == d {
			heap.Push(h, events[i][1])
			i++
		}
		for h.Len() > 0 && (*h)[0] < d {
			heap.Pop(h)
		}
		if h.Len() > 0 {
			heap.Pop(h)
			res++
		}
	}
	return res
}

func main() {
	fmt.Println(maxEvents([][]int{{1, 2}, {2, 3}, {3, 4}}))            // 3
	fmt.Println(maxEvents([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 2}}))    // 4
	fmt.Println(maxEvents([][]int{{1, 4}, {4, 4}, {2, 2}, {3, 4}, {1, 1}})) // 4
}
