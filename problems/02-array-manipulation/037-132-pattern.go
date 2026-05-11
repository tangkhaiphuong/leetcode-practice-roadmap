//go:build ignore

// Problem 037: 132 Pattern (LeetCode #456)
// Difficulty: Medium
// Categories: Array, Stack, Monotonic Stack
//
// Description:
//   Given nums, find any subsequence of indices i<j<k such that
//   nums[i] < nums[k] < nums[j]. Return true iff such triple exists.
//
// Examples:
//   [1,2,3,4]    -> false
//   [3,1,4,2]    -> true   (1 < 2 < 4)
//   [-1,3,2,0]   -> true   (-1 < 0 < 3)
//
// Approach: Right-to-Left Monotonic Stack
//   Walk from right keeping "potential nums[k]" as the largest popped value.
//   Stack holds candidates for nums[j] in decreasing order. For each x:
//     - while stack non-empty and x > top: pop, set k = popped (x is "j",
//       popped is "k"; popped < x).
//     - if x < k: x can serve as nums[i], we have a 132 — return true.
//     - else push x.
//
// Complexity:
//   Time:  O(n)
//   Space: O(n)

package main

import (
	"fmt"
	"math"
)

func find132pattern(nums []int) bool {
	stack := []int{}
	k := math.MinInt32
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < k {
			return true
		}
		for len(stack) > 0 && stack[len(stack)-1] < nums[i] {
			k = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return false
}

func main() {
	fmt.Println(find132pattern([]int{1, 2, 3, 4}))    // false
	fmt.Println(find132pattern([]int{3, 1, 4, 2}))    // true
	fmt.Println(find132pattern([]int{-1, 3, 2, 0}))   // true
}
