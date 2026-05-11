//go:build ignore

// Problem 021: One Edit Distance (LeetCode #161)
// Difficulty: Medium
// Categories: String, Two Pointers
//
// Description:
//   Given strings s and t, return true iff they are exactly one edit apart.
//   An edit is: insert one char, delete one char, or replace one char.
//
// Examples:
//   ("ab", "acb") -> true   (insert 'c')
//   ("",   "")    -> false  (zero edits, not one)
//   ("a", "A")    -> true   (replace)
//   ("abc","abc") -> false  (zero edits)
//
// Approach: Length-Difference Branch + Single Pass
//   - If |len(s) - len(t)| > 1, false.
//   - Make s the shorter one. Walk both with i, j.
//   - On first mismatch:
//       - if equal lengths: return s[i+1:] == t[j+1:] (replacement)
//       - else (t is longer): return s[i:] == t[j+1:] (insertion in t)
//   - If no mismatch found, true iff lengths differ by exactly 1.
//
// Complexity:
//   Time:  O(min(n,m))
//   Space: O(1)

package main

import "fmt"

func isOneEditDistance(s, t string) bool {
	if len(s) > len(t) {
		s, t = t, s
	}
	if len(t)-len(s) > 1 {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			if len(s) == len(t) {
				return s[i+1:] == t[i+1:]
			}
			return s[i:] == t[i+1:]
		}
	}
	// All chars matched up to len(s); valid iff t has exactly one more char.
	return len(t)-len(s) == 1
}

func main() {
	cases := []struct {
		a, b string
		want bool
	}{
		{"ab", "acb", true}, {"", "", false}, {"a", "A", true},
		{"abc", "abc", false}, {"", "a", true}, {"abc", "abxbc", false},
	}
	for _, c := range cases {
		fmt.Printf("oneEdit(%q,%q) = %v (want %v)\n", c.a, c.b, isOneEditDistance(c.a, c.b), c.want)
	}
}
