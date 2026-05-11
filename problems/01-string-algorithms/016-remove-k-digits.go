//go:build ignore

// Problem 016: Remove K Digits (LeetCode #402)
// Difficulty: Medium
// Categories: String, Stack, Greedy
//
// Description:
//   Given a non-negative integer string num and integer k, remove exactly k
//   digits so that the new number is the smallest possible. Result must not
//   have leading zeros (except "0" itself).
//
// Examples:
//   ("1432219", 3) -> "1219"
//   ("10200", 1)   -> "200"
//   ("10", 2)      -> "0"
//
// Approach: Monotonic Increasing Stack
//   Walk digits left to right. While k>0 and the last kept digit is greater
//   than current, pop it (we'd rather keep the smaller digit higher up).
//   Push current. After scan, if k>0 trim from the right.
//   Finally strip leading zeros.
//
// Complexity:
//   Time:  O(n)
//   Space: O(n)

package main

import (
	"fmt"
	"strings"
)

func removeKdigits(num string, k int) string {
	stack := make([]byte, 0, len(num))
	for i := 0; i < len(num); i++ {
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > num[i] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, num[i])
	}
	stack = stack[:len(stack)-k] // remove remaining from tail (already sorted up)
	res := strings.TrimLeft(string(stack), "0")
	if res == "" {
		return "0"
	}
	return res
}

func main() {
	cases := []struct {
		num  string
		k    int
		want string
	}{
		{"1432219", 3, "1219"},
		{"10200", 1, "200"},
		{"10", 2, "0"},
		{"112", 1, "11"},
	}
	for _, c := range cases {
		fmt.Printf("removeKdigits(%q,%d) = %q (want %q)\n", c.num, c.k, removeKdigits(c.num, c.k), c.want)
	}
}
