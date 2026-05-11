//go:build ignore

// Problem 068: Super Egg Drop (LeetCode #887)
// Difficulty: Hard
// Categories: DP, Binary Search, Math
//
// Description:
//   You have k eggs and a building with n floors. There exists a critical
//   floor f (0 <= f <= n) such that any egg dropped from > f breaks; from
//   <= f survives. With optimal strategy, return the MIN number of moves to
//   determine f, in the worst case.
//
// Examples:
//   k=1, n=2  -> 2
//   k=2, n=6  -> 3
//   k=3, n=14 -> 4
//
// Approach: DP "Max Floors with m moves and k eggs"
//   dp[m][k] = max floors that can be tested with m moves, k eggs.
//   Recurrence:
//     dp[m][k] = dp[m-1][k-1] + dp[m-1][k] + 1
//     (egg breaks: cover floors below; egg survives: floors above; +1 for
//      current.)
//   Answer = smallest m such that dp[m][k] >= n. m grows from 1 upward.
//
// Complexity:
//   Time:  O(k * answer)  with answer = O(log n)-ish for fixed k>=2
//   Space: O(k)

package main

import "fmt"

func superEggDrop(k, n int) int {
	dp := make([]int, k+1)
	m := 0
	for dp[k] < n {
		m++
		for j := k; j >= 1; j-- {
			dp[j] = dp[j-1] + dp[j] + 1
		}
	}
	return m
}

func main() {
	fmt.Println(superEggDrop(1, 2))  // 2
	fmt.Println(superEggDrop(2, 6))  // 3
	fmt.Println(superEggDrop(3, 14)) // 4
}
