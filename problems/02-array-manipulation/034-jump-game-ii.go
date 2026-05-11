//go:build ignore

// Problem 034: Jump Game II (LeetCode #45)
// Difficulty: Medium
// Categories: Array, DP, Greedy, BFS
//
// Description:
//   Given a 0-indexed array nums where nums[i] is the maximum jump length
//   from index i, return the minimum number of jumps to reach the last index.
//   You can assume you can always reach.
//
// Examples:
//   [2,3,1,1,4]   -> 2  (0->1->4)
//   [2,3,0,1,4]   -> 2
//
// Approach: BFS Layer-by-Layer (Greedy Range)
//   Track currentEnd (max idx reachable in this layer) and farthest (the max
//   reachable so far across the layer). When i hits currentEnd, increment
//   jump count and set currentEnd = farthest.
//
// Complexity:
//   Time:  O(n)
//   Space: O(1)

package main

import "fmt"

func jump(nums []int) int {
	jumps, end, far := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > far {
			far = i + nums[i]
		}
		if i == end {
			jumps++
			end = far
		}
	}
	return jumps
}

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4})) // 2
	fmt.Println(jump([]int{2, 3, 0, 1, 4})) // 2
	fmt.Println(jump([]int{1}))             // 0
}
