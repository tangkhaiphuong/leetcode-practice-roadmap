//go:build ignore

// Problem 045: Shortest Subarray to be Removed to Make Array Sorted (LC #1574)
// Difficulty: Medium
// Categories: Array, Two Pointers, Binary Search
//
// Description:
//   Given a non-decreasing-or-not array arr, remove ONE contiguous subarray
//   so that the remaining is non-decreasing. Return the minimum length of
//   such removal. (Empty array counts as sorted.)
//
// Examples:
//   [1,2,3,10,4,2,3,5] -> 3   (remove [10,4,2])
//   [5,4,3,2,1]        -> 4
//   [1,2,3]            -> 0
//
// Approach: Two Pointers from Each End
//   Find longest non-decreasing prefix end l, longest non-decreasing suffix
//   start r. If l >= r, already sorted -> 0.
//   Initialize answer = min(n - l - 1, r). Then for each i in [0..l],
//   advance r so arr[i] <= arr[r]; update answer = min(answer, r - i - 1).
//
// Complexity:
//   Time:  O(n)
//   Space: O(1)

package main

import "fmt"

func findLengthOfShortestSubarray(arr []int) int {
	n := len(arr)
	l := 0
	for l+1 < n && arr[l] <= arr[l+1] {
		l++
	}
	if l == n-1 {
		return 0
	}
	r := n - 1
	for r > 0 && arr[r-1] <= arr[r] {
		r--
	}
	best := minInt(n-l-1, r)
	i, j := 0, r
	for i <= l && j < n {
		if arr[i] <= arr[j] {
			best = minInt(best, j-i-1)
			i++
		} else {
			j++
		}
	}
	return best
}

func minInt(a, b int) int { if a < b { return a }; return b }

func main() {
	fmt.Println(findLengthOfShortestSubarray([]int{1, 2, 3, 10, 4, 2, 3, 5})) // 3
	fmt.Println(findLengthOfShortestSubarray([]int{5, 4, 3, 2, 1}))           // 4
	fmt.Println(findLengthOfShortestSubarray([]int{1, 2, 3}))                 // 0
}
