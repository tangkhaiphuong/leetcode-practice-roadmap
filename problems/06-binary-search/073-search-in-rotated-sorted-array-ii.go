//go:build ignore

// Problem 073: Search in Rotated Sorted Array II (LeetCode #81)
// Difficulty: Medium
// Categories: Binary Search, Array
//
// Description:
//   nums is a rotated, possibly duplicate-containing sorted array. Return
//   true iff target exists in nums.
//
// Examples:
//   [2,5,6,0,0,1,2], target=0 -> true
//   [2,5,6,0,0,1,2], target=3 -> false
//   [1,0,1,1,1],     target=0 -> true
//
// Approach: Modified Binary Search with Duplicate Skip
//   At each step, look at nums[lo], nums[mid], nums[hi]:
//     - if nums[mid] == target: true.
//     - if nums[lo] == nums[mid] == nums[hi]: shrink both ends by one — we
//       can't tell which half is sorted.
//     - else if nums[lo] <= nums[mid]: left half sorted. If target in
//       [nums[lo], nums[mid]), go left; else right.
//     - else: right half sorted. If target in (nums[mid], nums[hi]], go
//       right; else left.
//
// Complexity:
//   Time:  O(log n) average, O(n) worst case (all duplicates).
//   Space: O(1)

package main

import "fmt"

func search(nums []int, target int) bool {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := (lo + hi) / 2
		if nums[mid] == target {
			return true
		}
		if nums[lo] == nums[mid] && nums[mid] == nums[hi] {
			lo++
			hi--
			continue
		}
		if nums[lo] <= nums[mid] {
			if target >= nums[lo] && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	return false
}

func main() {
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 0)) // true
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 3)) // false
	fmt.Println(search([]int{1, 0, 1, 1, 1}, 0))       // true
}
