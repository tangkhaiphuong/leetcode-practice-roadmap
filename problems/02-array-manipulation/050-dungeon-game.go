//go:build ignore

// Problem 050: Dungeon Game (LeetCode #174)
// Difficulty: Hard
// Categories: Array, DP
//
// Description:
//   m x n grid; each cell has an integer (positive = potion, negative = damage).
//   The knight starts top-left, must reach bottom-right, only moving right or
//   down. HP must remain >= 1 at all times. Return the minimum starting HP
//   that lets him reach the princess.
//
// Examples:
//   [[-2,-3,3],[-5,-10,1],[10,30,-5]] -> 7
//
// Approach: Bottom-Up DP from Destination
//   dp[i][j] = min HP needed when entering cell (i,j) so the knight survives
//   to the end. Recurrence:
//     need = min(dp[i+1][j], dp[i][j+1]) - dungeon[i][j]
//     dp[i][j] = max(1, need)
//   Sentinels at right/bottom borders are +inf (or use 1 for the cell after
//   the goal).
//
// Complexity:
//   Time:  O(m*n)
//   Space: O(n)

package main

import (
	"fmt"
	"math"
)

func calculateMinimumHP(d [][]int) int {
	m, n := len(d), len(d[0])
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[n-1] = 1 // sentinel beyond bottom-right cell, so transitioning into the goal works
	// Actually, treat the cell-after-goal as needing 1 HP.
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			var nextMin int
			if i == m-1 && j == n-1 {
				nextMin = 1
			} else {
				nextMin = dp[j]
				if j+1 < n && dp[j+1] < nextMin {
					nextMin = dp[j+1]
				}
			}
			need := nextMin - d[i][j]
			if need < 1 {
				need = 1
			}
			dp[j] = need
		}
	}
	return dp[0]
}

func main() {
	d := [][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}
	fmt.Println(calculateMinimumHP(d)) // 7
}
