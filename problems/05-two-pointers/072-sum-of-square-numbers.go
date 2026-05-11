//go:build ignore

// Problem 072: Sum of Square Numbers (LeetCode #633)
// Difficulty: Medium
// Categories: Two Pointers, Math
//
// Description:
//   Given a non-negative integer c, return true iff there exist non-negative
//   integers a, b such that a^2 + b^2 == c.
//
// Examples:
//   c=5  -> true (1^2 + 2^2)
//   c=3  -> false
//   c=0  -> true (0^2 + 0^2)
//
// Approach: Two Pointers on [0, sqrt(c)]
//   l = 0, r = floor(sqrt(c)). Compute s = l*l + r*r.
//     - if s == c: true
//     - if s <  c: l++
//     - if s >  c: r--
//
//   Use uint64 to avoid overflow (c up to 2^31-1).
//
// Complexity:
//   Time:  O(sqrt(c))
//   Space: O(1)

package main

import (
	"fmt"
	"math"
)

func judgeSquareSum(c int) bool {
	l := uint64(0)
	r := uint64(math.Sqrt(float64(c)))
	for l <= r {
		s := l*l + r*r
		switch {
		case s == uint64(c):
			return true
		case s < uint64(c):
			l++
		default:
			r--
		}
	}
	return false
}

func main() {
	fmt.Println(judgeSquareSum(5)) // true
	fmt.Println(judgeSquareSum(3)) // false
	fmt.Println(judgeSquareSum(0)) // true
}
