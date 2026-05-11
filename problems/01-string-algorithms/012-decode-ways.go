//go:build ignore

// Problem 012: Decode Ways (LeetCode #91)
// Difficulty: Medium
// Categories: String, DP
//
// Description:
//   Letters A-Z map to "1".."26". Given a digit string s, count the number
//   of distinct ways to decode it. Leading zeros invalid (e.g. "06" can't
//   decode).
//
// Examples:
//   "12"   -> 2  ("AB" or "L")
//   "226"  -> 3  ("BZ", "VF", "BBF")
//   "06"   -> 0
//   "10"   -> 1
//
// Approach: Linear DP
//   Let dp[i] = ways to decode s[:i]. Then:
//     - if s[i-1] != '0': dp[i] += dp[i-1]
//     - if "10".."26" formed by s[i-2..i-1]: dp[i] += dp[i-2]
//   Use two rolling vars to save space.
//
// Complexity:
//   Time:  O(n)
//   Space: O(1)

package main

import "fmt"

func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}
	prev2, prev1 := 1, 1 // dp[0]=1, dp[1]=1 (since s[0]!='0' here)
	for i := 1; i < len(s); i++ {
		cur := 0
		if s[i] != '0' {
			cur += prev1
		}
		if two := (s[i-1]-'0')*10 + (s[i] - '0'); two >= 10 && two <= 26 {
			cur += prev2
		}
		prev2, prev1 = prev1, cur
	}
	return prev1
}

func main() {
	cases := []struct {
		s    string
		want int
	}{
		{"12", 2}, {"226", 3}, {"06", 0}, {"10", 1}, {"0", 0}, {"27", 1}, {"100", 0},
	}
	for _, c := range cases {
		got := numDecodings(c.s)
		fmt.Printf("numDecodings(%q) = %d (want %d)\n", c.s, got, c.want)
	}
}
