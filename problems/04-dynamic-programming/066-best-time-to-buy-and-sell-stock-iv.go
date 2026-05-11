//go:build ignore

// Problem 066: Best Time to Buy and Sell Stock IV (LeetCode #188)
// Difficulty: Hard
// Categories: DP
//
// Description:
//   Given prices[] and integer k, return the max profit with at most k
//   transactions (must sell before buying again).
//
// Examples:
//   k=2, prices=[2,4,1]      -> 2
//   k=2, prices=[3,2,6,5,0,3]-> 7
//
// Approach: O(n*k) DP with Two States Per Transaction
//   buy[j]  = max profit after j-th buy
//   sell[j] = max profit after j-th sell
//   Recurrence:
//     buy[j]  = max(buy[j], sell[j-1] - p)
//     sell[j] = max(sell[j], buy[j] + p)
//   Optimization: when k >= n/2, it's unlimited transactions — sum positive
//   diffs.
//
// Complexity:
//   Time:  O(n*k)
//   Space: O(k)

package main

import "fmt"

func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n == 0 || k == 0 {
		return 0
	}
	if k >= n/2 {
		profit := 0
		for i := 1; i < n; i++ {
			if d := prices[i] - prices[i-1]; d > 0 {
				profit += d
			}
		}
		return profit
	}
	buy := make([]int, k+1)
	sell := make([]int, k+1)
	for j := range buy {
		buy[j] = -1 << 30
	}
	for _, p := range prices {
		for j := 1; j <= k; j++ {
			if sell[j-1]-p > buy[j] {
				buy[j] = sell[j-1] - p
			}
			if buy[j]+p > sell[j] {
				sell[j] = buy[j] + p
			}
		}
	}
	return sell[k]
}

func main() {
	fmt.Println(maxProfit(2, []int{2, 4, 1}))          // 2
	fmt.Println(maxProfit(2, []int{3, 2, 6, 5, 0, 3})) // 7
}
