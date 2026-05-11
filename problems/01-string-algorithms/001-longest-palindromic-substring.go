//go:build ignore

// Problem 001: Longest Palindromic Substring (LeetCode #5)
// Difficulty: Medium
// Categories: String, Dynamic Programming
//
// Description:
//   Given a string s, return the longest palindromic substring in s.
//
// Examples:
//   Input:  s = "babad"     Output: "bab" (or "aba")
//   Input:  s = "cbbd"      Output: "bb"
//
// Constraints:
//   1 <= s.length <= 1000
//   s consists of only digits and English letters.
//
// Approach: Expand Around Center
//   Every palindrome has a center. For odd-length palindromes the center is a
//   single char; for even-length, it's between two chars. Try all 2n-1 centers
//   and expand outward while characters match. Track the best window.
//
//   Alternative: Manacher's algorithm gives O(n) but is rarely required in
//   interviews — expand-around-center hits the sweet spot of clarity vs perf.
//
// Complexity:
//   Time:  O(n^2) — n centers, each expands up to n.
//   Space: O(1)

package main

import "fmt"

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	start, maxLen := 0, 1
	expand := func(l, r int) {
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > maxLen {
				start, maxLen = l, r-l+1
			}
			l--
			r++
		}
	}
	for i := 0; i < len(s); i++ {
		expand(i, i)   // odd-length center
		expand(i, i+1) // even-length center
	}
	return s[start : start+maxLen]
}

func main() {
	cases := []struct {
		in  string
		hit []string // any of these accepted
	}{
		{"babad", []string{"bab", "aba"}},
		{"cbbd", []string{"bb"}},
		{"a", []string{"a"}},
		{"ac", []string{"a", "c"}},
		{"forgeeksskeegfor", []string{"geeksskeeg"}},
	}
	for _, c := range cases {
		got := longestPalindrome(c.in)
		ok := false
		for _, h := range c.hit {
			if got == h {
				ok = true
				break
			}
		}
		fmt.Printf("longestPalindrome(%q) = %q  pass=%v\n", c.in, got, ok)
	}
}
