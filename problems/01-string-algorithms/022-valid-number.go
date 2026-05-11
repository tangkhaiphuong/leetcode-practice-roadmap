//go:build ignore

// Problem 022: Valid Number (LeetCode #65)
// Difficulty: Hard
// Categories: String, Finite State Machine
//
// Description:
//   Return true iff s is a valid number. A valid number can be:
//     - decimal: optional sign, digits with optional dot, OR dot then digits,
//       OR digits then dot.
//     - integer: optional sign, digits.
//     - either of the above followed by 'e'/'E' and an integer.
//
// Examples (true): "0", "0.1", "abc"->false, "1 a"->false, "2e10", "-90E3", "-.9", "+6e-1"
// Examples (false): ".", "e3", "+-2", "1e2.5"
//
// Approach: Boolean Flags Single Pass
//   Track seenDigit, seenDot, seenE, plus digit-after-E flag. Sign valid only
//   at start or right after E. Dot valid only if not seenDot and not seenE.
//   At end, require seenDigit AND (no E OR digit-after-E).
//
// Complexity:
//   Time:  O(n)
//   Space: O(1)

package main

import "fmt"

func isNumber(s string) bool {
	seenDigit, seenDot, seenE, digitAfterE := false, false, false, true
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case c >= '0' && c <= '9':
			seenDigit = true
			digitAfterE = true
		case c == '+' || c == '-':
			if i != 0 && s[i-1] != 'e' && s[i-1] != 'E' {
				return false
			}
		case c == '.':
			if seenDot || seenE {
				return false
			}
			seenDot = true
		case c == 'e' || c == 'E':
			if seenE || !seenDigit {
				return false
			}
			seenE = true
			digitAfterE = false
		default:
			return false
		}
	}
	return seenDigit && digitAfterE
}

func main() {
	cases := []struct {
		s    string
		want bool
	}{
		{"0", true}, {"0.1", true}, {"abc", false}, {"1 a", false},
		{"2e10", true}, {"-90E3", true}, {"-.9", true}, {"+6e-1", true},
		{".", false}, {"e3", false}, {"+-2", false}, {"1e2.5", false},
	}
	for _, c := range cases {
		fmt.Printf("isNumber(%q) = %v (want %v)\n", c.s, isNumber(c.s), c.want)
	}
}
