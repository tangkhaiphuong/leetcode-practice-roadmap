//go:build ignore

// Problem 032: Median of Two Sorted Arrays (LeetCode #4)
// Difficulty: Hard
// Categories: Array, Binary Search, Divide & Conquer
//
// Description:
//   Given two sorted arrays nums1 and nums2 of sizes m and n, return the
//   median of the combined sorted array. Required runtime: O(log(m+n)).
//
// Examples:
//   [1,3], [2]      -> 2.0
//   [1,2], [3,4]    -> 2.5
//   [], [1]         -> 1.0
//
// Approach: Binary Search Partition
//   We want a partition (i in nums1, j in nums2 with i+j = (m+n+1)/2) such
//   that everything left <= everything right. Binary-search i in [0..m] on
//   the SHORTER array. With i fixed, j = half - i. Check:
//     L1 = nums1[i-1] (or -inf), R1 = nums1[i] (or +inf), similarly L2, R2.
//   If L1 <= R2 AND L2 <= R1: found. Median = (max(L1,L2) + min(R1,R2)) / 2
//   if total even, else max(L1,L2).
//
// Complexity:
//   Time:  O(log(min(m,n)))
//   Space: O(1)

package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(a, b []int) float64 {
	if len(a) > len(b) {
		a, b = b, a
	}
	m, n := len(a), len(b)
	half := (m + n + 1) / 2
	lo, hi := 0, m
	for lo <= hi {
		i := (lo + hi) / 2
		j := half - i
		L1, R1 := math.MinInt32, math.MaxInt32
		L2, R2 := math.MinInt32, math.MaxInt32
		if i > 0 {
			L1 = a[i-1]
		}
		if i < m {
			R1 = a[i]
		}
		if j > 0 {
			L2 = b[j-1]
		}
		if j < n {
			R2 = b[j]
		}
		if L1 <= R2 && L2 <= R1 {
			if (m+n)%2 == 1 {
				return float64(max(L1, L2))
			}
			return float64(max(L1, L2)+min(R1, R2)) / 2.0
		}
		if L1 > R2 {
			hi = i - 1
		} else {
			lo = i + 1
		}
	}
	return 0
}

func max(a, b int) int { if a > b { return a }; return b }
func min(a, b int) int { if a < b { return a }; return b }

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))    // 2
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4})) // 2.5
	fmt.Println(findMedianSortedArrays([]int{}, []int{1}))        // 1
}
