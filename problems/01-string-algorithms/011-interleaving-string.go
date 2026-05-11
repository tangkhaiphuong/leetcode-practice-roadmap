//go:build ignore

// Problem 011: Interleaving String (LeetCode #97)
// Difficulty: Medium
// Categories: String, Dynamic Programming
//
// Description:
//   Given strings s1, s2, s3, return true iff s3 is formed by interleaving
//   s1 and s2 (preserving order within each).
//
// Examples:
//   ("aabcc", "dbbca", "aadbbcbcac") -> true
//   ("aabcc", "dbbca", "aadbbbaccc") -> false
//   ("", "", "")                     -> true
//
// Constraints:
//   0 <= s1.len, s2.len <= 100; 0 <= s3.len <= 200.
//
// Approach: 2D DP
//   dp[i][j] = true iff s1[:i] + s2[:j] interleaves to form s3[:i+j].
//   Transition: dp[i][j] true if
//     (dp[i-1][j] AND s1[i-1] == s3[i+j-1])
//   OR
//     (dp[i][j-1] AND s2[j-1] == s3[i+j-1]).
//
// Complexity:
//   Time:  O(n*m)
//   Space: O(m), rolling 1D array.

package main

import "fmt"

func isInterleave(s1, s2, s3 string) bool {
	n, m := len(s1), len(s2)
	if n+m != len(s3) {
		return false
	}
	dp := make([]bool, m+1)
	dp[0] = true
	for j := 1; j <= m; j++ {
		dp[j] = dp[j-1] && s2[j-1] == s3[j-1]
	}
	for i := 1; i <= n; i++ {
		dp[0] = dp[0] && s1[i-1] == s3[i-1]
		for j := 1; j <= m; j++ {
			dp[j] = (dp[j] && s1[i-1] == s3[i+j-1]) ||
				(dp[j-1] && s2[j-1] == s3[i+j-1])
		}
	}
	return dp[m]
}

func main() {
	cases := []struct {
		s1, s2, s3 string
		want       bool
	}{
		{"aabcc", "dbbca", "aadbbcbcac", true},
		{"aabcc", "dbbca", "aadbbbaccc", false},
		{"", "", "", true},
		{"a", "", "a", true},
		{"a", "b", "ab", true},
	}
	for _, c := range cases {
		got := isInterleave(c.s1, c.s2, c.s3)
		fmt.Printf("isInterleave(%q,%q,%q) = %v (want %v)\n", c.s1, c.s2, c.s3, got, c.want)
	}
}
