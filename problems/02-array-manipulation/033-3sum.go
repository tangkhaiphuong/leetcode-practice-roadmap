//go:build ignore

// Problem 033: 3Sum (LeetCode #15)
// Difficulty: Medium
// Categories: Array, Two Pointers, Sorting
//
// Description:
//   Given integer array nums, return all unique triplets [a,b,c] such that
//   a+b+c == 0.
//
// Examples:
//   [-1,0,1,2,-1,-4] -> [[-1,-1,2],[-1,0,1]]
//   [0,0,0]          -> [[0,0,0]]
//   [0,1,1]          -> []
//
// Approach: Sort + Two Pointers
//   Sort. For each i, use two pointers l=i+1, r=n-1 to find pairs summing to
//   -nums[i]. Skip duplicates: at index i, if nums[i] == nums[i-1], skip; on
//   match, advance l/r past duplicates.
//   Early break when nums[i] > 0 (sum can't be 0 anymore).
//
// Complexity:
//   Time:  O(n^2)
//   Space: O(1) (excluding output)

package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	n := len(nums)
	for i := 0; i < n-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, n-1
		for l < r {
			s := nums[i] + nums[l] + nums[r]
			switch {
			case s == 0:
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			case s < 0:
				l++
			default:
				r--
			}
		}
	}
	return res
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4})) // [[-1 -1 2] [-1 0 1]]
	fmt.Println(threeSum([]int{0, 0, 0}))             // [[0 0 0]]
	fmt.Println(threeSum([]int{0, 1, 1}))             // []
}
