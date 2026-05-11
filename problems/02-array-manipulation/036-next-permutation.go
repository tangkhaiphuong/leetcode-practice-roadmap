//go:build ignore

// Problem 036: Next Permutation (LeetCode #31)
// Difficulty: Medium
// Categories: Array, Two Pointers
//
// Description:
//   Rearrange numbers into the lexicographically next greater permutation.
//   If not possible (already largest), rearrange to the lowest (sorted asc).
//   Modify in place, O(1) extra space.
//
// Examples:
//   [1,2,3] -> [1,3,2]
//   [3,2,1] -> [1,2,3]
//   [1,1,5] -> [1,5,1]
//
// Approach: Find Pivot, Swap, Reverse
//   1. Walk from right; find the first i with nums[i] < nums[i+1] (pivot).
//      If none, reverse whole array.
//   2. Find the smallest j > i with nums[j] > nums[i] (rightmost such j).
//   3. Swap nums[i], nums[j].
//   4. Reverse nums[i+1:] (it was descending, now ascending = smallest tail).
//
// Complexity:
//   Time:  O(n)
//   Space: O(1)

package main

import "fmt"

func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := n - 1
		for nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	for l, r := i+1, n-1; l < r; l, r = l+1, r-1 {
		nums[l], nums[r] = nums[r], nums[l]
	}
}

func main() {
	for _, c := range [][]int{{1, 2, 3}, {3, 2, 1}, {1, 1, 5}, {1, 3, 2}} {
		nextPermutation(c)
		fmt.Println(c)
	}
}
