//go:build ignore

// Problem 008: Regular Expression Matching (LeetCode #10)
// Difficulty: Hard
// Categories: String, Dynamic Programming
//
// Description:
//   Given a string s and a pattern p, implement regex matching with support
//   for '.' (any single char) and '*' (zero or more of preceding element).
//   The match must cover the *entire* string.
//
// Examples:
//   ("aa", "a")    -> false
//   ("aa", "a*")   -> true
//   ("ab", ".*")   -> true
//   ("aab", "c*a*b") -> true
//   ("mississippi", "mis*is*p*.") -> false
//
// Constraints:
//   1 <= s.length <= 20, 1 <= p.length <= 20.
//   '*' always preceded by a valid element. p contains only a-z, '.', '*'.
//
// Approach: 2D DP
//   dp[i][j] = true iff s[:i] matches p[:j].
//   Transitions:
//     - p[j-1] == '*': two cases
//         (a) zero of preceding: dp[i][j-2]
//         (b) one+ of preceding: dp[i-1][j] AND (s[i-1] matches p[j-2])
//     - else: dp[i-1][j-1] AND (s[i-1] matches p[j-1])
//   "matches" means equal char or pattern char is '.'.
//
// Complexity:
//   Time:  O(n*m)
//   Space: O(n*m), reducible to O(m).

package main

import "fmt"

func isMatch(s, p string) bool {
	n, m := len(s), len(p)
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, m+1)
	}
	dp[0][0] = true
	// Empty s can match patterns like a*b*c* -> true at j positions where '*'
	// allows the preceding element to vanish.
	for j := 2; j <= m; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}
	matches := func(si, pi int) bool {
		return p[pi] == '.' || p[pi] == s[si]
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-2] // zero occurrences
				if matches(i-1, j-2) {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			} else if matches(i-1, j-1) {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[n][m]
}

func main() {
	cases := []struct {
		s, p string
		want bool
	}{
		{"aa", "a", false},
		{"aa", "a*", true},
		{"ab", ".*", true},
		{"aab", "c*a*b", true},
		{"mississippi", "mis*is*p*.", false},
		{"mississippi", "mis*is*ip*.", true},
	}
	for _, c := range cases {
		got := isMatch(c.s, c.p)
		fmt.Printf("isMatch(%q, %q) = %v (want %v)\n", c.s, c.p, got, c.want)
	}
}
