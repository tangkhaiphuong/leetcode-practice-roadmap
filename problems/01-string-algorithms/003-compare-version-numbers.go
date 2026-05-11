//go:build ignore

// Problem 003: Compare Version Numbers (LeetCode #165)
// Difficulty: Medium
// Categories: String, Two Pointers
//
// Description:
//   Given two version strings version1 and version2, compare them. Versions
//   are revision-dot-separated (e.g. "1.2", "1.0.1"). Compare revision-by-
//   revision as integers. Missing revisions are treated as 0.
//   Return:
//     -1 if version1 < version2
//      1 if version1 > version2
//      0 if equal
//
// Examples:
//   "1.01"   vs "1.001"   -> 0  (leading zeros ignored)
//   "1.0"    vs "1.0.0"   -> 0  (missing revision = 0)
//   "0.1"    vs "1.1"     -> -1
//   "1.0.1"  vs "1"       -> 1
//
// Constraints:
//   1 <= version1.length, version2.length <= 500
//   Each component is a non-negative integer not exceeding 2^31 - 1.
//
// Approach: Two Pointers Without Splitting
//   Splitting via strings.Split costs extra allocation. Walk both strings
//   simultaneously: at each step parse the next integer (until '.' or end),
//   compare them, advance past the dots. When one string ends, treat its
//   remaining revisions as 0.
//
// Complexity:
//   Time:  O(n + m)
//   Space: O(1)

package main

import "fmt"

func compareVersion(v1, v2 string) int {
	i, j := 0, 0
	for i < len(v1) || j < len(v2) {
		var a, b int
		for i < len(v1) && v1[i] != '.' {
			a = a*10 + int(v1[i]-'0')
			i++
		}
		for j < len(v2) && v2[j] != '.' {
			b = b*10 + int(v2[j]-'0')
			j++
		}
		if a != b {
			if a < b {
				return -1
			}
			return 1
		}
		i++ // skip '.'
		j++
	}
	return 0
}

func main() {
	cases := [][3]string{
		{"1.01", "1.001", "0"},
		{"1.0", "1.0.0", "0"},
		{"0.1", "1.1", "-1"},
		{"1.0.1", "1", "1"},
		{"7.5.2.4", "7.5.3", "-1"},
	}
	for _, c := range cases {
		fmt.Printf("compare(%q, %q) = %d (expect %s)\n", c[0], c[1], compareVersion(c[0], c[1]), c[2])
	}
}
