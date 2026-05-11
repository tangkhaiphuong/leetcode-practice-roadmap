//go:build ignore

// Problem 018: Decode Ways II (LeetCode #639)
// Difficulty: Hard
// Categories: String, DP
//
// Description:
//   Same mapping as Decode Ways (#91), but s may contain '*' which represents
//   any digit '1'..'9'. Return the number of ways to decode s, modulo 1e9+7.
//
// Examples:
//   "*"   -> 9   (1..9)
//   "1*"  -> 18  ("1A".."1I" if * is 1..9, plus "K".."S" -> 11..19, but 10 not in 1..9)
//                Actually: dp gives 9 (single-digit) + 9 (two-digit 11..19) = 18.
//   "2*"  -> 15  (single 9 + two-digit 21..26 = 6) -> 15
//
// Approach: DP with case-by-case enumerations.
//   Let f(c1) = #digits c1 can represent as a single decode (1..9 -> 1, 0->0,
//   '*' -> 9). Let g(c1, c2) = #two-digit decodings 10..26.
//
//   dp[i] = f(s[i-1]) * dp[i-1] + g(s[i-2], s[i-1]) * dp[i-2]
//
//   For g(a, b):
//     '*','*' -> 15  (11..19, 21..26)
//     '*', d  -> if d<=6: 2 (1d, 2d); else: 1 (1d)
//      a, '*' -> a==1: 9; a==2: 6; else 0
//      a,  d  -> 1 if 10*a+d in [10..26], else 0
//
// Complexity:
//   Time:  O(n)
//   Space: O(1)

package main

import "fmt"

const mod = 1_000_000_007

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	one := func(c byte) int {
		switch {
		case c == '*':
			return 9
		case c == '0':
			return 0
		default:
			return 1
		}
	}
	two := func(a, b byte) int {
		switch {
		case a == '*' && b == '*':
			return 15
		case a == '*':
			if b <= '6' {
				return 2
			}
			return 1
		case b == '*':
			if a == '1' {
				return 9
			}
			if a == '2' {
				return 6
			}
			return 0
		default:
			n := int(a-'0')*10 + int(b-'0')
			if n >= 10 && n <= 26 {
				return 1
			}
			return 0
		}
	}
	prev2, prev1 := 1, one(s[0])
	for i := 1; i < len(s); i++ {
		cur := (one(s[i])*prev1 + two(s[i-1], s[i])*prev2) % mod
		prev2, prev1 = prev1, cur
	}
	return prev1
}

func main() {
	cases := []struct {
		in   string
		want int
	}{
		{"*", 9}, {"1*", 18}, {"2*", 15}, {"**", 96}, {"0", 0},
	}
	for _, c := range cases {
		fmt.Printf("numDecodings(%q) = %d (want %d)\n", c.in, numDecodings(c.in), c.want)
	}
}
