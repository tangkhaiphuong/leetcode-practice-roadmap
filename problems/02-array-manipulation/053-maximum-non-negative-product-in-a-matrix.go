//go:build ignore

// Problem 053: Maximum Non Negative Product in a Matrix (LeetCode #1594)
// Difficulty: Medium
// Categories: Array (Matrix), DP
//
// Description:
//   m x n grid of integers (may be negative). Start at (0,0), end at (m-1,n-1),
//   moving only right or down. Return the maximum non-negative product of all
//   numbers along a path, modulo 1e9+7. If no non-negative product is
//   achievable, return -1.
//
// Examples:
//   [[-1,-2,-3],[-2,-3,-3],[-3,-3,-2]] -> -1
//   [[1,-2,1],[1,-2,1],[3,-4,1]]       -> 8
//   [[1,3],[0,-4]]                     -> 0
//
// Approach: Track Min AND Max Products at Each Cell
//   Negative * negative could become a large positive. dp[i][j] holds (min,max).
//     Coming from (i-1,j) or (i,j-1):
//       options = {prevMin*v, prevMax*v}; for each predecessor.
//     newMin = min over all options; newMax = max.
//   Final answer = dp[m-1][n-1].max if >= 0 else -1, modulo 1e9+7 after.
//
// Complexity:
//   Time:  O(m*n)
//   Space: O(m*n)

package main

import "fmt"

const mod53 = 1_000_000_007

func maxProductPath(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	type pair struct{ lo, hi int64 }
	dp := make([][]pair, m)
	for i := range dp {
		dp[i] = make([]pair, n)
	}
	dp[0][0] = pair{int64(grid[0][0]), int64(grid[0][0])}
	for j := 1; j < n; j++ {
		v := int64(grid[0][j])
		dp[0][j] = pair{dp[0][j-1].lo * v, dp[0][j-1].hi * v}
		if v < 0 {
			dp[0][j].lo, dp[0][j].hi = dp[0][j].hi, dp[0][j].lo
		}
	}
	for i := 1; i < m; i++ {
		v := int64(grid[i][0])
		dp[i][0] = pair{dp[i-1][0].lo * v, dp[i-1][0].hi * v}
		if v < 0 {
			dp[i][0].lo, dp[i][0].hi = dp[i][0].hi, dp[i][0].lo
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			v := int64(grid[i][j])
			cands := []int64{
				dp[i-1][j].lo * v, dp[i-1][j].hi * v,
				dp[i][j-1].lo * v, dp[i][j-1].hi * v,
			}
			lo, hi := cands[0], cands[0]
			for _, c := range cands[1:] {
				if c < lo {
					lo = c
				}
				if c > hi {
					hi = c
				}
			}
			dp[i][j] = pair{lo, hi}
		}
	}
	res := dp[m-1][n-1].hi
	if res < 0 {
		return -1
	}
	return int(res % mod53)
}

func main() {
	fmt.Println(maxProductPath([][]int{{-1, -2, -3}, {-2, -3, -3}, {-3, -3, -2}})) // -1
	fmt.Println(maxProductPath([][]int{{1, -2, 1}, {1, -2, 1}, {3, -4, 1}}))       // 8
	fmt.Println(maxProductPath([][]int{{1, 3}, {0, -4}}))                          // 0
}
