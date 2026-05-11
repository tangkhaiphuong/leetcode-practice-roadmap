//go:build ignore

// Problem 056: Closest Room (LeetCode #1847)
// Difficulty: Hard
// Categories: Array, Binary Search, Sorting, Offline Queries
//
// Description:
//   Rooms[i] = [roomId, size]. Queries[j] = [preferredId, minSize]. For each
//   query, return the room id with size >= minSize whose id is closest to
//   preferredId (smaller id wins on tie). Return -1 if no such room.
//
// Examples:
//   rooms=[[2,2],[1,2],[3,2]], queries=[[3,1],[3,3],[5,2]]  -> [3,-1,3]
//
// Approach: Offline Sweep, Sort Both by Size Descending
//   Sort rooms by size desc. Sort queries by minSize desc, preserving original
//   index. Maintain a sorted set of qualifying room ids (those with size >=
//   current minSize). For each query in order, insert all rooms with size >=
//   minSize (they qualify forever). Then binary-search the set for the
//   neighbor of preferredId.
//
// Complexity:
//   Time:  O((n+q) log n)
//   Space: O(n+q)

package main

import (
	"fmt"
	"sort"
)

func closestRoom(rooms [][]int, queries [][]int) []int {
	sort.Slice(rooms, func(i, j int) bool { return rooms[i][1] > rooms[j][1] })
	q := len(queries)
	order := make([]int, q)
	for i := range order {
		order[i] = i
	}
	sort.Slice(order, func(i, j int) bool {
		return queries[order[i]][1] > queries[order[j]][1]
	})

	ids := []int{}
	insert := func(x int) {
		i := sort.SearchInts(ids, x)
		ids = append(ids, 0)
		copy(ids[i+1:], ids[i:])
		ids[i] = x
	}

	res := make([]int, q)
	rIdx := 0
	for _, qi := range order {
		pref, minSize := queries[qi][0], queries[qi][1]
		for rIdx < len(rooms) && rooms[rIdx][1] >= minSize {
			insert(rooms[rIdx][0])
			rIdx++
		}
		if len(ids) == 0 {
			res[qi] = -1
			continue
		}
		k := sort.SearchInts(ids, pref)
		best := -1
		bestDiff := -1
		consider := func(idx int) {
			if idx < 0 || idx >= len(ids) {
				return
			}
			d := abs(ids[idx] - pref)
			if best == -1 || d < bestDiff || (d == bestDiff && ids[idx] < best) {
				best = ids[idx]
				bestDiff = d
			}
		}
		consider(k)
		consider(k - 1)
		res[qi] = best
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	rooms := [][]int{{2, 2}, {1, 2}, {3, 2}}
	queries := [][]int{{3, 1}, {3, 3}, {5, 2}}
	fmt.Println(closestRoom(rooms, queries)) // [3 -1 3]
}
