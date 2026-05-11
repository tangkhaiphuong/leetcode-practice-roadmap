//go:build ignore

// Problem 004: Longest Substring Without Repeating Characters (LeetCode #3)
// Difficulty: Medium
// Categories: String, Hash Table, Sliding Window
//
// Description:
//   Given a string s, find the length of the longest substring without
//   repeating characters.
//
// Examples:
//   "abcabcbb" -> 3 ("abc")
//   "bbbbb"    -> 1 ("b")
//   "pwwkew"   -> 3 ("wke")
//
// Constraints:
//   0 <= s.length <= 5*10^4
//   s consists of English letters, digits, symbols, and spaces.
//
// Approach: Sliding Window with last-seen map
//   Keep a window [l, r]. As we extend r, if s[r] was seen at position p >= l,
//   jump l to p+1 (no need to retreat). Update best = max(best, r-l+1).
//   Storing the *position* (not just presence) lets us advance l in O(1)
//   instead of shrinking one char at a time.
//
// Complexity:
//   Time:  O(n)
//   Space: O(min(n, charset))

package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	last := make(map[byte]int)
	best, l := 0, 0
	for r := 0; r < len(s); r++ {
		if p, ok := last[s[r]]; ok && p >= l {
			l = p + 1
		}
		last[s[r]] = r
		if r-l+1 > best {
			best = r - l + 1
		}
	}
	return best
}

func main() {
	cases := []struct {
		in   string
		want int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{"dvdf", 3},
		{" ", 1},
	}
	for _, c := range cases {
		got := lengthOfLongestSubstring(c.in)
		fmt.Printf("len(%q) = %d (want %d) pass=%v\n", c.in, got, c.want, got == c.want)
	}
}
