//go:build ignore

// Problem 007: String to Integer (atoi) (LeetCode #8)
// Difficulty: Medium
// Categories: String
//
// Description:
//   Implement atoi which converts a string to a 32-bit signed integer.
//   Algorithm:
//     1. Skip leading whitespace.
//     2. Optional '+' or '-'.
//     3. Read digits until non-digit or end.
//     4. If empty -> 0. Apply sign.
//     5. Clamp to [-2^31, 2^31 - 1].
//
// Examples:
//   "42"            -> 42
//   "   -42"        -> -42
//   "4193 with words" -> 4193
//   "words 987"     -> 0
//   "-91283472332"  -> -2147483648 (clamped)
//
// Approach: Linear scan, single pass.
//   Detect overflow before assignment: if result > maxInt32/10, OR
//   (result == maxInt32/10 AND digit > maxInt32%10) -> clamp.
//   Treat negative side similarly.
//
// Complexity:
//   Time:  O(n)
//   Space: O(1)

package main

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {
	i, n := 0, len(s)
	for i < n && s[i] == ' ' {
		i++
	}
	sign := 1
	if i < n && (s[i] == '+' || s[i] == '-') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}
	res := 0
	for i < n && s[i] >= '0' && s[i] <= '9' {
		d := int(s[i] - '0')
		// Overflow guard before res becomes invalid.
		if res > math.MaxInt32/10 || (res == math.MaxInt32/10 && d > math.MaxInt32%10) {
			if sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}
		res = res*10 + d
		i++
	}
	return sign * res
}

func main() {
	for _, s := range []string{"42", "   -42", "4193 with words", "words 987",
		"-91283472332", "+1", "  +0 123", ""} {
		fmt.Printf("atoi(%q) = %d\n", s, myAtoi(s))
	}
}
