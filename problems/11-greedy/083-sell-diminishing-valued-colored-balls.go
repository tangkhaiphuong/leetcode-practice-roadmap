//go:build ignore

// Problem 083: Sell Diminishing-Valued Colored Balls (LeetCode #1648)
// Difficulty: Medium
// Categories: Greedy, Heap, Math, Binary Search, Sorting
//
// Description:
//   inventory[i] = number of balls of color i. Each ball of color i sold for
//   the CURRENT count of color-i balls (then count decreases). Sell exactly
//   orders balls. Return the max profit modulo 1e9+7.
//
// Examples:
//   inventory=[2,5], orders=4    -> 14   (5+4+3+2)
//   inventory=[3,5], orders=6    -> 19
//   inventory=[2,8,4,10,6], orders=20 -> 110
//
// Approach: Sort Desc + Bulk Levels
//   Process colors in decreasing inventory order. Conceptually we're filling
//   layers: "everyone tied at the top" gets a level, then drops by 1.
//   With n_top colors tied at value v and next color value u (u < v), we can
//   sell n_top balls per level for levels v, v-1, ..., u+1. That contributes
//   n_top * sum(u+1..v).
//   When orders is exhausted in mid-band, distribute remaining: full levels
//   + remainder orders at the next level value.
//
//   Alt: binary-search the floor price T such that selling everything above T
//   uses <= orders.
//
// Complexity:
//   Time:  O(n log n)
//   Space: O(1)

package main

import (
	"fmt"
	"sort"
)

const mod83 = 1_000_000_007

func maxProfit(inventory []int, orders int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(inventory)))
	inventory = append(inventory, 0)
	res := int64(0)
	for i := 0; i < len(inventory)-1; i++ {
		v := inventory[i]
		u := inventory[i+1]
		width := i + 1
		levels := v - u
		levelsCap := orders / width
		if levels <= levelsCap {
			// take all 'levels' levels at full width
			res = (res + int64(width)*sumRange(int64(u+1), int64(v))) % mod83
			orders -= levels * width
		} else {
			// take levelsCap full levels, then remainder
			res = (res + int64(width)*sumRange(int64(v-levelsCap+1), int64(v))) % mod83
			rem := orders - levelsCap*width
			res = (res + int64(rem)*int64(v-levelsCap)) % mod83
			orders = 0
			break
		}
		if orders == 0 {
			break
		}
	}
	return int(res)
}

// sumRange returns (a + a+1 + ... + b) mod mod83.
func sumRange(a, b int64) int64 {
	if a > b {
		return 0
	}
	n := b - a + 1
	total := (a + b) * n / 2
	return total % mod83
}

func main() {
	fmt.Println(maxProfit([]int{2, 5}, 4))             // 14
	fmt.Println(maxProfit([]int{3, 5}, 6))             // 19
	fmt.Println(maxProfit([]int{2, 8, 4, 10, 6}, 20))  // 110
}
