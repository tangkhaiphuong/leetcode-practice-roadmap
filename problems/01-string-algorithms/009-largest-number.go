//go:build ignore

// Problem 009: Largest Number (LeetCode #179)
// Difficulty: Medium
// Categories: String, Greedy, Sorting
//
// Description:
//   Given a list of non-negative integers nums, arrange them so they form the
//   largest possible number. Return as a string.
//
// Examples:
//   [10,2]              -> "210"
//   [3,30,34,5,9]       -> "9534330"
//   [0,0]               -> "0"
//
// Constraints:
//   1 <= nums.length <= 100;  0 <= nums[i] <= 10^9.
//
// Approach: Custom Comparator on String Concat
//   For two strings a, b: a should come before b iff a+b > b+a (lexicographic
//   on equal-length combinations works because both are equal length). This
//   yields a total order. Sort, concat, then strip leading zeros (only "0"
//   case if all zeros).
//
// Complexity:
//   Time:  O(n log n * k), where k = average digit length
//   Space: O(n)

package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	strs := make([]string, len(nums))
	for i, v := range nums {
		strs[i] = strconv.Itoa(v)
	}
	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] > strs[j]+strs[i]
	})
	res := strings.Join(strs, "")
	if res[0] == '0' {
		return "0"
	}
	return res
}

func main() {
	cases := []struct {
		in   []int
		want string
	}{
		{[]int{10, 2}, "210"},
		{[]int{3, 30, 34, 5, 9}, "9534330"},
		{[]int{0, 0}, "0"},
		{[]int{1}, "1"},
	}
	for _, c := range cases {
		fmt.Printf("largestNumber(%v) = %q (want %q)\n", c.in, largestNumber(c.in), c.want)
	}
}
