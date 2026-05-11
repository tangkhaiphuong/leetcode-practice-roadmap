//go:build ignore

// Problem 044: Shortest Subarray with Sum at Least K (LeetCode #862)
// Difficulty: Hard
// Categories: Array, Deque (Monotonic), Prefix Sum, Binary Search
//
// Description:
//   Given an integer array nums (may contain negatives) and integer k, return
//   the length of the shortest non-empty contiguous subarray with sum >= k,
//   or -1 if none.
//
// Examples:
//   nums=[1], k=1            -> 1
//   nums=[1,2], k=4          -> -1
//   nums=[2,-1,2], k=3       -> 3
//
// Approach: Monotonic Deque on Prefix Sums
//   Let P[i] = sum of nums[:i]. Subarray sum nums[l..r-1] = P[r] - P[l].
//   We want the smallest r-l with P[r] - P[l] >= k.
//   Slide r over P. Maintain deque of candidate l indices with INCREASING P
//   values:
//     - while P[r] - P[deque.front()] >= k: record r - front, pop front
//       (any later r yields longer answer for the same front).
//     - while deque non-empty and P[r] <= P[deque.back()]: pop back (worse
//       candidate — smaller prefix later means current r dominates).
//     - push r.
//
// Complexity:
//   Time:  O(n)
//   Space: O(n)

package main

import "fmt"

func shortestSubarray(nums []int, k int) int {
	n := len(nums)
	P := make([]int, n+1)
	for i, v := range nums {
		P[i+1] = P[i] + v
	}
	dq := make([]int, 0, n+1)
	best := n + 1
	for r := 0; r <= n; r++ {
		for len(dq) > 0 && P[r]-P[dq[0]] >= k {
			if r-dq[0] < best {
				best = r - dq[0]
			}
			dq = dq[1:]
		}
		for len(dq) > 0 && P[r] <= P[dq[len(dq)-1]] {
			dq = dq[:len(dq)-1]
		}
		dq = append(dq, r)
	}
	if best == n+1 {
		return -1
	}
	return best
}

func main() {
	fmt.Println(shortestSubarray([]int{1}, 1))           // 1
	fmt.Println(shortestSubarray([]int{1, 2}, 4))        // -1
	fmt.Println(shortestSubarray([]int{2, -1, 2}, 3))    // 3
	fmt.Println(shortestSubarray([]int{84, -37, 32}, 80))// 3
}
