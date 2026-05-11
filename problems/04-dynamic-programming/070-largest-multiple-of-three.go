//go:build ignore

// Problem 070: Largest Multiple of Three (LeetCode #1363)
// Difficulty: Hard
// Categories: DP, Greedy, Math
//
// Description:
//   Given a non-empty array of digits, return the largest multiple of three
//   you can form by concatenating some of the digits in any order. If
//   impossible, return "". If the result is "0", return "0" (no leading 0s).
//
// Examples:
//   [8,1,9]       -> "981"
//   [8,6,7,1,0]   -> "8760"
//   [1]           -> ""
//   [0,0,0,0,0,0] -> "0"
//
// Approach: Greedy on Sum mod 3
//   Sum all digits. If sum % 3 == 0: use all (sorted desc).
//   If sum % 3 == 1: drop the smallest digit with d%3==1; if none, drop two
//     smallest with d%3==2.
//   If sum % 3 == 2: drop the smallest with d%3==2; else drop two smallest
//     with d%3==1.
//   Then sort remaining digits descending; if first is '0', return "0".
//
// Complexity:
//   Time:  O(n log n) (or O(n) with bucket sort by digit)
//   Space: O(n)

package main

import (
	"fmt"
	"sort"
	"strings"
)

func largestMultipleOfThree(digits []int) string {
	sum := 0
	rem := [3][]int{}
	for _, d := range digits {
		sum += d
		rem[d%3] = append(rem[d%3], d)
	}
	for i := 0; i < 3; i++ {
		sort.Ints(rem[i]) // ascending
	}
	r := sum % 3
	drop := map[int]int{} // digit -> count to drop
	switch r {
	case 1:
		if len(rem[1]) > 0 {
			drop[rem[1][0]]++
		} else if len(rem[2]) >= 2 {
			drop[rem[2][0]]++
			drop[rem[2][1]]++
		} else {
			return ""
		}
	case 2:
		if len(rem[2]) > 0 {
			drop[rem[2][0]]++
		} else if len(rem[1]) >= 2 {
			drop[rem[1][0]]++
			drop[rem[1][1]]++
		} else {
			return ""
		}
	}
	cnt := [10]int{}
	for _, d := range digits {
		cnt[d]++
	}
	for d, c := range drop {
		cnt[d] -= c
	}
	var sb strings.Builder
	for d := 9; d >= 0; d-- {
		for i := 0; i < cnt[d]; i++ {
			sb.WriteByte(byte('0' + d))
		}
	}
	res := sb.String()
	if len(res) == 0 {
		return ""
	}
	if res[0] == '0' {
		return "0"
	}
	return res
}

func main() {
	fmt.Println(largestMultipleOfThree([]int{8, 1, 9}))             // "981"
	fmt.Println(largestMultipleOfThree([]int{8, 6, 7, 1, 0}))       // "8760"
	fmt.Println(largestMultipleOfThree([]int{1}))                   // ""
	fmt.Println(largestMultipleOfThree([]int{0, 0, 0, 0, 0, 0}))    // "0"
}
