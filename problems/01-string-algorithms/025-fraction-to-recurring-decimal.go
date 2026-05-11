//go:build ignore

// Problem 025: Fraction to Recurring Decimal (LeetCode #166)
// Difficulty: Medium
// Categories: String, Hash Table, Math
//
// Description:
//   Given numerator and denominator, return the fraction in string form. If
//   the fractional part is repeating, enclose the repeating part in parens.
//
// Examples:
//   1, 2   -> "0.5"
//   2, 1   -> "2"
//   4, 333 -> "0.(012)"
//   -50, 8 -> "-6.25"
//
// Constraints:
//   |num|, |den| <= 2^31 - 1, den != 0.
//
// Approach: Long Division with Remainder Map
//   Determine sign, work in absolute values (use int64 to avoid overflow on
//   INT_MIN). Compute integer part. Then perform long division: at each step,
//   record (remainder -> position in fractional buffer). When a remainder
//   recurs, we've found the cycle — wrap with parens. If remainder hits 0,
//   stop.
//
// Complexity:
//   Time:  O(d) where d = number of distinct remainders before cycle/0
//   Space: O(d)

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func fractionToDecimal(num, den int) string {
	if num == 0 {
		return "0"
	}
	var sb strings.Builder
	if (num < 0) != (den < 0) {
		sb.WriteByte('-')
	}
	n := absI64(int64(num))
	d := absI64(int64(den))
	sb.WriteString(strconv.FormatInt(n/d, 10))
	rem := n % d
	if rem == 0 {
		return sb.String()
	}
	sb.WriteByte('.')
	pos := map[int64]int{} // remainder -> position in builder
	for rem != 0 {
		if p, ok := pos[rem]; ok {
			out := sb.String()
			return out[:p] + "(" + out[p:] + ")"
		}
		pos[rem] = sb.Len()
		rem *= 10
		sb.WriteByte(byte('0' + rem/d))
		rem %= d
	}
	return sb.String()
}

func absI64(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	cases := [][2]int{{1, 2}, {2, 1}, {4, 333}, {-50, 8}, {1, 6}, {0, 5}, {-1, -2147483648}}
	for _, c := range cases {
		fmt.Printf("%d/%d = %q\n", c[0], c[1], fractionToDecimal(c[0], c[1]))
	}
}
