//go:build ignore

// Problem 067: Maximum Score from Performing Multiplication Operations (LC #1770)
// Difficulty: Hard
// Categories: DP
//
// Description:
//   You have nums (length n) and multipliers (length m, m <= n). On each of m
//   operations, choose to take nums[0] (front) or nums[n-1] (back), multiply
//   by multipliers[op], add to score, and remove that end. Return the max
//   total score.
//
// Examples:
//   nums=[1,2,3], multipliers=[3,2,1]      -> 14
//   nums=[-5,-3,-3,-2,7,1], multipliers=[-10,-5,3,4,6] -> 102
//
// Approach: 2D DP Indexed by (op, leftTaken)
//   After op operations and L taken from the left, R = op - L taken from
//   right. Right pointer index = n - 1 - R.
//     dp[op][L] = max(
//        dp[op-1][L-1] + multipliers[op-1] * nums[L-1],            // took left
//        dp[op-1][L]   + multipliers[op-1] * nums[n-1-(op-1-L)]    // took right
//     )
//   Iterate op from 1..m, L from 0..op. Bottom-up with an O(m) rolling array.
//
// Complexity:
//   Time:  O(m^2)
//   Space: O(m)

package main

import "fmt"

func maximumScore(nums, multipliers []int) int {
	n, m := len(nums), len(multipliers)
	const NEG = -1 << 60
	dp := make([]int, m+1)
	prev := make([]int, m+1)
	for i := range prev {
		prev[i] = NEG
	}
	prev[0] = 0
	for op := 1; op <= m; op++ {
		dp[op] = NEG
		for L := 0; L <= op; L++ {
			best := NEG
			// took left
			if L > 0 && prev[L-1] != NEG {
				v := prev[L-1] + multipliers[op-1]*nums[L-1]
				if v > best {
					best = v
				}
			}
			// took right
			if L <= op-1 && prev[L] != NEG {
				rightIdx := n - 1 - (op - 1 - L)
				v := prev[L] + multipliers[op-1]*nums[rightIdx]
				if v > best {
					best = v
				}
			}
			dp[L] = best
		}
		// shift
		for L := op + 1; L <= m; L++ {
			dp[L] = NEG
		}
		prev, dp = dp, prev
	}
	best := prev[0]
	for L := 1; L <= m; L++ {
		if prev[L] > best {
			best = prev[L]
		}
	}
	return best
}

func main() {
	fmt.Println(maximumScore([]int{1, 2, 3}, []int{3, 2, 1}))                       // 14
	fmt.Println(maximumScore([]int{-5, -3, -3, -2, 7, 1}, []int{-10, -5, 3, 4, 6})) // 102
}
