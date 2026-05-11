//go:build ignore

// Problem 035: Maximum Product Subarray (LeetCode #152)
// Difficulty: Medium
// Categories: Array, DP, Greedy
//
// Description:
//   Find the contiguous non-empty subarray within nums having the largest
//   product, and return that product.
//
// Examples:
//   [2,3,-2,4]   -> 6
//   [-2,0,-1]    -> 0
//   [-2,3,-4]    -> 24
//
// Approach: Track Min and Max Ending Here
//   With negatives, a small (very negative) min can become the max if
//   multiplied by another negative. Maintain curMax, curMin at each step:
//     newMax = max(nums[i], curMax*nums[i], curMin*nums[i])
//     newMin = min(nums[i], curMax*nums[i], curMin*nums[i])
//   Track best across positions.
//
// Complexity:
//   Time:  O(n)
//   Space: O(1)

package main

import "fmt"

func maxProduct(nums []int) int {
	curMax, curMin, best := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		x := nums[i]
		a, b, c := x, curMax*x, curMin*x
		newMax := max3(a, b, c)
		newMin := min3(a, b, c)
		curMax, curMin = newMax, newMin
		if curMax > best {
			best = curMax
		}
	}
	return best
}

func max3(a, b, c int) int {
	m := a
	if b > m { m = b }
	if c > m { m = c }
	return m
}
func min3(a, b, c int) int {
	m := a
	if b < m { m = b }
	if c < m { m = c }
	return m
}

func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4})) // 6
	fmt.Println(maxProduct([]int{-2, 0, -1}))   // 0
	fmt.Println(maxProduct([]int{-2, 3, -4}))   // 24
}
