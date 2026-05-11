//go:build ignore

// Problem 051: Wiggle Sort II (LeetCode #324)
// Difficulty: Medium
// Categories: Array, Sorting, Quickselect, Divide & Conquer
//
// Description:
//   Given an integer array, reorder so nums[0] < nums[1] > nums[2] < nums[3]...
//   You may assume the input has a valid answer.
//
// Examples:
//   [1,5,1,1,6,4]   -> [1,4,1,5,1,6] (one valid)
//   [1,3,2,2,3,1]   -> [2,3,1,3,1,2] (one valid)
//
// Approach: Sort + Index Rewiring (O(n log n) clean version)
//   Sort. Split into low half L and high half H (with the larger group going
//   to L if odd length).
//   Place L (reversed) at even indices, H (reversed) at odd indices.
//   Reversing prevents equal medians from sitting adjacent.
//
//   The optimal O(n) version uses quickselect for the median + Dutch National
//   Flag with index mapping; the sort-based version is what most candidates
//   should produce in interviews.
//
// Complexity:
//   Time:  O(n log n)
//   Space: O(n)

package main

import (
	"fmt"
	"sort"
)

func wiggleSort(nums []int) {
	n := len(nums)
	tmp := make([]int, n)
	copy(tmp, nums)
	sort.Ints(tmp)
	mid := (n + 1) / 2
	// Place reversed lower half at even indices, reversed upper half at odd.
	j := mid - 1
	for i := 0; i < n; i += 2 {
		nums[i] = tmp[j]
		j--
	}
	j = n - 1
	for i := 1; i < n; i += 2 {
		nums[i] = tmp[j]
		j--
	}
}

func main() {
	a := []int{1, 5, 1, 1, 6, 4}
	wiggleSort(a)
	fmt.Println(a)
	b := []int{1, 3, 2, 2, 3, 1}
	wiggleSort(b)
	fmt.Println(b)
}
