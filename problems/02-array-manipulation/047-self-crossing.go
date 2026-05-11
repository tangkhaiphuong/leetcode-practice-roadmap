//go:build ignore

// Problem 047: Self Crossing (LeetCode #335)
// Difficulty: Hard
// Categories: Array, Math, Geometry
//
// Description:
//   You walk along the +Y, -X, -Y, +X axes (counterclockwise) by distances
//   distance[0..n-1]. Return true iff your path crosses itself.
//
// Examples:
//   [2,1,1,2]   -> true
//   [1,2,3,4]   -> false
//   [1,1,1,1]   -> true
//
// Approach: Three Geometric Cases
//   Crossings can only occur in three patterns when looking back at the most
//   recent moves:
//     Case 1: 4th line crosses 1st  -> d[i] >= d[i-2] AND d[i-1] <= d[i-3]
//     Case 2: 5th line meets 1st    -> d[i-1]==d[i-3] AND d[i]+d[i-4] >= d[i-2]
//     Case 3: 6th line crosses 1st  -> d[i-2] >= d[i-4] AND
//                                       d[i-3] >= d[i-1] >= d[i-3]-d[i-5] AND
//                                       d[i-2]-d[i-4] <= d[i] <= d[i-2] AND ...
//
// Complexity:
//   Time:  O(n)
//   Space: O(1)

package main

import "fmt"

func isSelfCrossing(d []int) bool {
	n := len(d)
	for i := 3; i < n; i++ {
		// Case 1
		if d[i] >= d[i-2] && d[i-1] <= d[i-3] {
			return true
		}
		// Case 2
		if i >= 4 && d[i-1] == d[i-3] && d[i]+d[i-4] >= d[i-2] {
			return true
		}
		// Case 3
		if i >= 5 &&
			d[i-2] >= d[i-4] &&
			d[i-3] >= d[i-1] && d[i-1] >= d[i-3]-d[i-5] &&
			d[i] >= d[i-2]-d[i-4] && d[i] <= d[i-2] {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(isSelfCrossing([]int{2, 1, 1, 2})) // true
	fmt.Println(isSelfCrossing([]int{1, 2, 3, 4})) // false
	fmt.Println(isSelfCrossing([]int{1, 1, 1, 1})) // true
}
