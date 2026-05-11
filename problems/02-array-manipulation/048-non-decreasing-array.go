//go:build ignore

// Problem 048: Non-decreasing Array (LeetCode #665)
// Difficulty: Medium
// Categories: Array, Greedy
//
// Description:
//   Given nums, return true iff it can be made non-decreasing by modifying
//   at most ONE element.
//
// Examples:
//   [4,2,3]   -> true   (->[2,2,3])
//   [4,2,1]   -> false
//   [3,4,2,3] -> false
//
// Approach: Single Pass with One Repair
//   Scan; on the first violation nums[i] > nums[i+1], increment a counter.
//   To "repair", lower nums[i+1] to nums[i] when possible:
//     - if i==0 OR nums[i-1] <= nums[i+1], set nums[i] = nums[i+1] (drop it).
//     - else, raise nums[i+1] to nums[i] (set nums[i+1] = nums[i]).
//   If a second violation is seen, return false.
//
// Complexity:
//   Time:  O(n)
//   Space: O(1)

package main

import "fmt"

func checkPossibility(nums []int) bool {
	mods := 0
	for i := 0; i+1 < len(nums); i++ {
		if nums[i] <= nums[i+1] {
			continue
		}
		mods++
		if mods > 1 {
			return false
		}
		if i == 0 || nums[i-1] <= nums[i+1] {
			nums[i] = nums[i+1]
		} else {
			nums[i+1] = nums[i]
		}
	}
	return true
}

func main() {
	fmt.Println(checkPossibility([]int{4, 2, 3}))    // true
	fmt.Println(checkPossibility([]int{4, 2, 1}))    // false
	fmt.Println(checkPossibility([]int{3, 4, 2, 3})) // false
}
