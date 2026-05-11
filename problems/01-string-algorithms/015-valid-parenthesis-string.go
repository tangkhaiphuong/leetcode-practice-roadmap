//go:build ignore

// Problem 015: Valid Parenthesis String (LeetCode #678)
// Difficulty: Medium
// Categories: String, DP, Stack, Greedy
//
// Description:
//   String contains '(', ')', and '*'. '*' may be treated as '(' or ')' or
//   empty. Return true iff the string can be made valid.
//
// Examples:
//   "()"      -> true
//   "(*)"     -> true
//   "(*))"    -> true
//   "(((*"    -> false
//
// Approach: Greedy with Lo/Hi range of open count
//   Track [lo, hi] = the min and max possible number of unmatched '('.
//   - '(' -> lo++, hi++
//   - ')' -> lo--, hi--
//   - '*' -> lo--, hi++  (could close, open, or empty)
//   Clamp lo at 0 (can't be negative). If hi < 0, more ')' than possible '('.
//   At end, return lo == 0.
//
// Complexity:
//   Time:  O(n)
//   Space: O(1)

package main

import "fmt"

func checkValidString(s string) bool {
	lo, hi := 0, 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			lo++
			hi++
		case ')':
			if lo > 0 {
				lo--
			}
			hi--
		case '*':
			if lo > 0 {
				lo--
			}
			hi++
		}
		if hi < 0 {
			return false
		}
	}
	return lo == 0
}

func main() {
	cases := []struct {
		in   string
		want bool
	}{
		{"()", true}, {"(*)", true}, {"(*))", true}, {"(((*", false}, {"", true}, {"((**", true},
	}
	for _, c := range cases {
		fmt.Printf("checkValidString(%q) = %v (want %v)\n", c.in, checkValidString(c.in), c.want)
	}
}
