//go:build ignore

// Problem 078: Create Maximum Number (LeetCode #321)
// Difficulty: Hard
// Categories: Stack (Monotonic), Greedy
//
// Description:
//   Given two int arrays of digits (length m and n) and integer k = m+n... no:
//   Given two arrays nums1, nums2 (each non-negative digits 0..9), and k <=
//   m+n, create the maximum number of length k preserving the relative order
//   of digits within each input. Return the resulting array.
//
// Examples:
//   nums1=[3,4,6,5], nums2=[9,1,2,5,8,3], k=5 -> [9,8,6,5,3]
//   nums1=[6,7], nums2=[6,0,4], k=5            -> [6,7,6,0,4]
//
// Approach: Try Each Split + Maxify Each Subarray + Best Merge
//   For each i in [max(0, k-n), min(k, m)]:
//     - Pick the largest length-i subsequence from nums1 (monotonic-decreasing
//       stack with allowed pops up to m-i).
//     - Pick the largest length-(k-i) subsequence from nums2 similarly.
//     - Merge them choosing the lexicographically larger remaining each step.
//   Track best across all splits.
//
// Complexity:
//   Time:  O(k * (m+n)^2) in the worst case
//   Space: O(k)

package main

import "fmt"

func maxNumber(nums1, nums2 []int, k int) []int {
	m, n := len(nums1), len(nums2)
	var best []int
	for i := max78(0, k-n); i <= min78(k, m); i++ {
		a := maxSub(nums1, i)
		b := maxSub(nums2, k-i)
		c := merge(a, b)
		if greater(c, 0, best, 0) {
			best = c
		}
	}
	return best
}

func maxSub(nums []int, k int) []int {
	stk := make([]int, 0, k)
	drop := len(nums) - k
	for _, x := range nums {
		for len(stk) > 0 && drop > 0 && stk[len(stk)-1] < x {
			stk = stk[:len(stk)-1]
			drop--
		}
		stk = append(stk, x)
	}
	return stk[:k]
}

func merge(a, b []int) []int {
	res := make([]int, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) || j < len(b) {
		if greater(a, i, b, j) {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}
	return res
}

func greater(a []int, i int, b []int, j int) bool {
	for i < len(a) && j < len(b) && a[i] == b[j] {
		i++
		j++
	}
	return j == len(b) || (i < len(a) && a[i] > b[j])
}

func max78(a, b int) int { if a > b { return a }; return b }
func min78(a, b int) int { if a < b { return a }; return b }

func main() {
	fmt.Println(maxNumber([]int{3, 4, 6, 5}, []int{9, 1, 2, 5, 8, 3}, 5)) // [9 8 6 5 3]
	fmt.Println(maxNumber([]int{6, 7}, []int{6, 0, 4}, 5))                // [6 7 6 0 4]
}
