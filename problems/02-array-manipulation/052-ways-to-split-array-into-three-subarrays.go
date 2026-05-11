//go:build ignore

// Problem 052: Ways to Split Array Into Three Subarrays (LeetCode #1712)
// Difficulty: Medium
// Categories: Array, Two Pointers, Prefix Sum, Binary Search
//
// Description:
//   Given non-negative array nums, count the number of ways to split it into
//   three non-empty contiguous parts (left, mid, right) such that
//   sum(left) <= sum(mid) <= sum(right). Return modulo 1e9+7.
//
// Examples:
//   [1,1,1]       -> 1
//   [1,2,2,2,5,0] -> 3
//   [3,2,1]       -> 0
//
// Approach: Prefix Sums + Two Pointers Bounds
//   P[i] = sum of nums[:i]. For each i (end of left), the mid starts at i and
//   we need to find the range of j where:
//     P[j] - P[i] >= P[i]              -> j minimum bound
//     P[n] - P[j] >= P[j] - P[i]       -> j maximum bound (this means
//                                         2*P[j] <= P[n]+P[i])
//   Slide two pointers lo, hi monotonically as i grows. Add (hi - lo + 1)
//   if hi >= lo and lo <= n-1.
//
// Complexity:
//   Time:  O(n)
//   Space: O(n)

package main

import "fmt"

const mod52 = 1_000_000_007

func waysToSplit(nums []int) int {
	n := len(nums)
	P := make([]int, n+1)
	for i, v := range nums {
		P[i+1] = P[i] + v
	}
	res := 0
	lo, hi := 1, 1
	for i := 1; i < n-1; i++ {
		// lo: smallest j > i with P[j]-P[i] >= P[i] => P[j] >= 2*P[i].
		if lo < i+1 {
			lo = i + 1
		}
		for lo < n && P[lo]-P[i] < P[i] {
			lo++
		}
		// hi: largest j < n with 2*P[j] <= P[n]+P[i].
		if hi < lo {
			hi = lo
		}
		for hi < n && 2*P[hi] <= P[n]+P[i] {
			hi++
		}
		if hi-1 >= lo {
			res = (res + hi - lo) % mod52
		}
	}
	return res
}

func main() {
	fmt.Println(waysToSplit([]int{1, 1, 1}))          // 1
	fmt.Println(waysToSplit([]int{1, 2, 2, 2, 5, 0})) // 3
	fmt.Println(waysToSplit([]int{3, 2, 1}))          // 0
}
