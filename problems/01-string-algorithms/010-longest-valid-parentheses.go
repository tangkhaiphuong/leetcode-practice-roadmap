//go:build ignore

// Problem 010: Longest Valid Parentheses (LeetCode #32)
// Difficulty: Hard
// Categories: String, DP, Stack
//
// Description:
//   Given a string containing only '(' and ')', return the length of the
//   longest valid (well-formed) parentheses substring.
//
// Examples:
//   "(()"     -> 2     ("()")
//   ")()())"  -> 4     ("()()")
//   ""        -> 0
//
// Constraints:
//   0 <= s.length <= 3*10^4. s in { '(', ')' }.
//
// Approach: Stack of Indices
//   Push -1 as a sentinel "last unmatched index". For '(': push index. For
//   ')': pop. If stack now empty, push i as new sentinel. Otherwise length =
//   i - stack.top(). Track max.
//
//   Alt: O(1)-space scan twice with counters l/r — covered in many solutions.
//
// Complexity:
//   Time:  O(n)
//   Space: O(n)

package main

import "fmt"

func longestValidParentheses(s string) int {
	stk := []int{-1}
	best := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stk = append(stk, i)
			continue
		}
		stk = stk[:len(stk)-1]
		if len(stk) == 0 {
			stk = append(stk, i)
		} else if cur := i - stk[len(stk)-1]; cur > best {
			best = cur
		}
	}
	return best
}

func main() {
	cases := []struct {
		in   string
		want int
	}{
		{"(()", 2},
		{")()())", 4},
		{"", 0},
		{"()(()", 2},
		{"()(())", 6},
	}
	for _, c := range cases {
		got := longestValidParentheses(c.in)
		fmt.Printf("longest(%q) = %d (want %d)\n", c.in, got, c.want)
	}
}
