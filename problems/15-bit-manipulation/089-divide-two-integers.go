//go:build ignore

// Problem 089: Divide Two Integers (LeetCode #29)
// Difficulty: Medium
// Categories: Bit Manipulation, Math
//
// Description:
//   Divide two integers without using multiplication, division, or mod.
//   Truncate toward zero. Clamp result to int32 range [-2^31, 2^31 - 1].
//
// Examples:
//   10 / 3   -> 3
//   7 / -3   -> -2
//   INT_MIN / -1 -> INT_MAX (clamped)
//
// Approach: Repeated Doubling Subtraction
//   Determine sign. Work with absolute values via uint64 to avoid INT_MIN
//   pitfall. While dividend >= divisor, double the divisor and corresponding
//   quotient until next double would exceed dividend; subtract; accumulate.
//   Apply sign and clamp.
//
// Complexity:
//   Time:  O(log^2 n)
//   Space: O(1)

package main

import (
	"fmt"
	"math"
)

func divide(dividend, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	neg := (dividend < 0) != (divisor < 0)
	a := absI64(int64(dividend))
	b := absI64(int64(divisor))
	var q int64
	for a >= b {
		t := b
		mul := int64(1)
		for a >= (t << 1) && (t<<1) > 0 {
			t <<= 1
			mul <<= 1
		}
		a -= t
		q += mul
	}
	if neg {
		q = -q
	}
	if q > math.MaxInt32 {
		return math.MaxInt32
	}
	if q < math.MinInt32 {
		return math.MinInt32
	}
	return int(q)
}

func absI64(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(divide(10, 3))                  // 3
	fmt.Println(divide(7, -3))                  // -2
	fmt.Println(divide(math.MinInt32, -1))      // 2147483647
	fmt.Println(divide(math.MinInt32, 1))       // -2147483648
}
