//go:build ignore

// Problem 063: Count Good Meals (LeetCode #1711)
// Difficulty: Medium
// Categories: Hash Table, Math, Two-Sum-like
//
// Description:
//   Given deliciousness[], count pairs (i,j), i<j, whose sum is a power of 2.
//   Return modulo 1e9+7.
//
// Examples:
//   [1,3,5,7,9]      -> 4   (pairs summing to 4 or 8 or 16: many)
//   [1,1,1,3,3,3,7]  -> 15
//
// Approach: Two-Sum Hashing for Each Power-of-2 Target
//   Walk values left-to-right. For each x, for each power-of-2 P in
//   {1, 2, 4, ..., 2^21} (since max value is 2^20, max sum 2^21), look up
//   count of (P - x) seen so far and add to result. Then count[x]++.
//   Constraint: deliciousness[i] in [0, 2^20].
//
// Complexity:
//   Time:  O(n * 22)
//   Space: O(n)

package main

import "fmt"

const mod63 = 1_000_000_007

func countPairs(deliciousness []int) int {
	count := map[int]int{}
	res := 0
	for _, x := range deliciousness {
		for p := 1; p <= 1<<21; p <<= 1 {
			res = (res + count[p-x]) % mod63
		}
		count[x]++
	}
	return res
}

func main() {
	fmt.Println(countPairs([]int{1, 3, 5, 7, 9}))         // 4
	fmt.Println(countPairs([]int{1, 1, 1, 3, 3, 3, 7}))   // 15
}
