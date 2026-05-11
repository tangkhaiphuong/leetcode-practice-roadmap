//go:build ignore

// Problem 049: Continuous Subarray Sum (LeetCode #523)
// Difficulty: Medium
// Categories: Array, Hash Table, Prefix Sum, Math
//
// Description:
//   Given nums and integer k, return true iff nums has a continuous subarray
//   of size at least 2 whose sum is a multiple of k (positive, negative, or
//   zero, where 0 is a multiple of every k != 0).
//
// Examples:
//   nums=[23,2,4,6,7], k=6        -> true ([2,4])
//   nums=[23,2,6,4,7], k=6        -> true ([23,2,6,4,7])
//   nums=[23,2,6,4,7], k=13       -> false
//
// Approach: Prefix Sum Mod k
//   Two prefix sums with the same remainder mod k bracket a subarray with
//   sum divisible by k. Track first index where each remainder appears
//   (initialize {0: -1} for the empty prefix). When the same remainder repeats
//   at index j > first_index + 1, return true.
//
// Complexity:
//   Time:  O(n)
//   Space: O(min(n, k))

package main

import "fmt"

func checkSubarraySum(nums []int, k int) bool {
	first := map[int]int{0: -1}
	sum := 0
	for i, v := range nums {
		sum += v
		r := sum
		if k != 0 {
			r %= k
		}
		if p, ok := first[r]; ok {
			if i-p >= 2 {
				return true
			}
		} else {
			first[r] = i
		}
	}
	return false
}

func main() {
	fmt.Println(checkSubarraySum([]int{23, 2, 4, 6, 7}, 6))  // true
	fmt.Println(checkSubarraySum([]int{23, 2, 6, 4, 7}, 6))  // true
	fmt.Println(checkSubarraySum([]int{23, 2, 6, 4, 7}, 13)) // false
}
