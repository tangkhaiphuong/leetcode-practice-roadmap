//go:build ignore

// Problem 014: Wildcard Matching (LeetCode #44)
// Difficulty: Hard
// Categories: String, DP, Greedy
//
// Description:
//   Implement matching with '?' (any single char) and '*' (any sequence
//   including empty). The match must cover the entire string.
//
// Examples:
//   ("aa", "a")    -> false
//   ("aa", "*")    -> true
//   ("cb", "?a")   -> false
//   ("adceb", "*a*b") -> true
//
// Approach: 2D DP
//   dp[i][j] = true iff s[:i] matches p[:j].
//   - p[j-1] == '*': dp[i][j] = dp[i-1][j] (consume one s char) OR dp[i][j-1]
//     (treat * as empty so far).
//   - p[j-1] == '?' OR equals s[i-1]: dp[i][j] = dp[i-1][j-1].
//
//   An O(1)-extra-space greedy alternative also exists; DP is the safe choice
//   for interviews to demonstrate the recurrence.
//
// Complexity:
//   Time:  O(n*m)
//   Space: O(m), rolling 1D.

package main

import "fmt"

func isMatch(s, p string) bool {
	n, m := len(s), len(p)
	dp := make([]bool, m+1)
	dp[0] = true
	for j := 1; j <= m && p[j-1] == '*'; j++ {
		dp[j] = true
	}
	for i := 1; i <= n; i++ {
		prevDiag := dp[0]
		dp[0] = false
		for j := 1; j <= m; j++ {
			cur := dp[j]
			switch {
			case p[j-1] == '*':
				dp[j] = dp[j] || dp[j-1] // consume s char OR star empty
			case p[j-1] == '?' || p[j-1] == s[i-1]:
				dp[j] = prevDiag
			default:
				dp[j] = false
			}
			prevDiag = cur
		}
	}
	return dp[m]
}

func main() {
	cases := []struct {
		s, p string
		want bool
	}{
		{"aa", "a", false},
		{"aa", "*", true},
		{"cb", "?a", false},
		{"adceb", "*a*b", true},
		{"acdcb", "a*c?b", false},
		{"", "*", true},
	}
	for _, c := range cases {
		fmt.Printf("isMatch(%q, %q) = %v (want %v)\n", c.s, c.p, isMatch(c.s, c.p), c.want)
	}
}
