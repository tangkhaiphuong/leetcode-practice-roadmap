//go:build ignore

// Problem 017: Palindrome Partitioning II (LeetCode #132)
// Difficulty: Hard
// Categories: String, DP
//
// Description:
//   Given s, partition s such that every substring of the partition is a
//   palindrome. Return the minimum number of cuts needed.
//
// Examples:
//   "aab"   -> 1   ("aa" | "b")
//   "a"     -> 0
//   "ab"    -> 1
//
// Approach: O(n^2) DP with palindrome table
//   1. Build pal[i][j] = is s[i..j] palindrome (expand-around-center O(n^2)).
//   2. cuts[i] = min cuts for s[:i+1].
//      For each i, if pal[0][i] then cuts[i] = 0.
//      Else cuts[i] = min over j<=i of (cuts[j-1] + 1) where pal[j][i].
//
// Complexity:
//   Time:  O(n^2)
//   Space: O(n^2). Can be O(n) using on-the-fly palindrome expansion + cuts.

package main

import "fmt"

func minCut(s string) int {
	n := len(s)
	pal := make([][]bool, n)
	for i := range pal {
		pal[i] = make([]bool, n)
	}
	// Expand around centers.
	for c := 0; c < n; c++ {
		for l, r := c, c; l >= 0 && r < n && s[l] == s[r]; l, r = l-1, r+1 {
			pal[l][r] = true
		}
		for l, r := c, c+1; l >= 0 && r < n && s[l] == s[r]; l, r = l-1, r+1 {
			pal[l][r] = true
		}
	}
	cuts := make([]int, n)
	for i := 0; i < n; i++ {
		if pal[0][i] {
			cuts[i] = 0
			continue
		}
		cuts[i] = i // upper bound
		for j := 1; j <= i; j++ {
			if pal[j][i] && cuts[j-1]+1 < cuts[i] {
				cuts[i] = cuts[j-1] + 1
			}
		}
	}
	return cuts[n-1]
}

func main() {
	cases := []struct {
		in   string
		want int
	}{
		{"aab", 1}, {"a", 0}, {"ab", 1}, {"aabaa", 0}, {"abcbm", 2},
	}
	for _, c := range cases {
		fmt.Printf("minCut(%q) = %d (want %d)\n", c.in, minCut(c.in), c.want)
	}
}
