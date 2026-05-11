//go:build ignore

// Problem 019: Repeated String Match (LeetCode #686)
// Difficulty: Medium
// Categories: String
//
// Description:
//   Given strings a and b, return the minimum number of times a must be
//   repeated so that b is a substring of the result. Return -1 if impossible.
//
// Examples:
//   ("abcd", "cdabcdab") -> 3   ("abcd"*3 = "abcdabcdabcd" contains b)
//   ("a", "aa")          -> 2
//   ("a", "a")           -> 1
//   ("abc", "wxyz")      -> -1
//
// Approach: Bound + Substring Search
//   Repeat a until len >= len(b). Check if it contains b. If not, repeat once
//   more (to cover the wraparound where b straddles boundaries). If still
//   missing, return -1.
//
//   Why it suffices: any occurrence of b in a^k starts at some offset in
//   [0, len(a)). Once the prefix of length len(a) is repeated enough to cover
//   len(b)+len(a), we've checked every possible starting position.
//
// Complexity:
//   Time:  O(n*m) using strings.Contains (or O(n+m) with KMP).
//   Space: O(n+m)

package main

import (
	"fmt"
	"strings"
)

func repeatedStringMatch(a, b string) int {
	count := 1
	sb := strings.Builder{}
	sb.WriteString(a)
	for sb.Len() < len(b) {
		sb.WriteString(a)
		count++
	}
	if strings.Contains(sb.String(), b) {
		return count
	}
	sb.WriteString(a)
	if strings.Contains(sb.String(), b) {
		return count + 1
	}
	return -1
}

func main() {
	cases := []struct {
		a, b string
		want int
	}{
		{"abcd", "cdabcdab", 3},
		{"a", "aa", 2},
		{"a", "a", 1},
		{"abc", "wxyz", -1},
	}
	for _, c := range cases {
		fmt.Printf("repeated(%q,%q) = %d (want %d)\n", c.a, c.b, repeatedStringMatch(c.a, c.b), c.want)
	}
}
