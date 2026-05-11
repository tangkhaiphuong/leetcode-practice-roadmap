//go:build ignore

// Problem 042: Shortest Unsorted Continuous Subarray (LeetCode #581)
// Difficulty: Medium
// Categories: Array, Two Pointers, Stack
//
// Description:
//   Given nums, find the length of the shortest contiguous subarray that, if
//   sorted ascending, makes the whole array sorted. Return 0 if already
//   sorted.
//
// Examples:
//   [2,6,4,8,10,9,15] -> 5  (sort [6,4,8,10,9])
//   [1,2,3,4]         -> 0
//   [1]               -> 0
//
// Approach: Single-Pass Min/Max Window
//   Walk left-to-right tracking the running max; if nums[i] < runningMax,
//   it's "out of place" — record rightmost such i as 'right'.
//   Walk right-to-left tracking running min; if nums[i] > runningMin, record
//   leftmost such i as 'left'.
//   Answer = right - left + 1 (or 0 if right < left).
//
// Complexity:
//   Time:  O(n)
//   Space: O(1)

package main

import (
	"fmt"
	"math"
)

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	left, right := -1, -2 // so right-left+1 = 0 when not assigned
	maxL, minR := math.MinInt32, math.MaxInt32
	for i := 0; i < n; i++ {
		if nums[i] < maxL {
			right = i
		} else {
			maxL = nums[i]
		}
		j := n - 1 - i
		if nums[j] > minR {
			left = j
		} else {
			minR = nums[j]
		}
	}
	return right - left + 1
}

func main() {
	fmt.Println(findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15})) // 5
	fmt.Println(findUnsortedSubarray([]int{1, 2, 3, 4}))            // 0
	fmt.Println(findUnsortedSubarray([]int{1}))                     // 0
}
