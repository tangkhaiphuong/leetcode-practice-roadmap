//go:build ignore

// Problem 069: Integer Replacement (LeetCode #397)
// Difficulty: Medium
// Categories: DP / Greedy / Bit Manipulation, Memoization
//
// Description:
//   Given n, in one step you can:
//     - if n is even: n /= 2
//     - if n is odd: n += 1 OR n -= 1
//   Return the min steps to reduce n to 1.
//
// Examples:
//   n=8   -> 3   (8->4->2->1)
//   n=7   -> 4   (7->8->4->2->1)
//   n=4   -> 2
//
// Approach: Greedy on Bit Patterns
//   When n is even: shift right.
//   When n is odd: look at the last two bits.
//     - n == 3 -> n-1 (special; 3->2->1 cheaper than 3->4->2->1).
//     - bits 11 -> n+1 (creates a trailing 0 to shift more later).
//     - bits 01 -> n-1.
//
//   This works because adding 1 to '...11' clears multiple trailing ones,
//   while subtracting only removes one — but for '...01' subtracting wins.
//
// Complexity:
//   Time:  O(log n)
//   Space: O(1)

package main

import "fmt"

func integerReplacement(n int) int {
	steps := 0
	x := uint64(n)
	for x != 1 {
		if x&1 == 0 {
			x >>= 1
		} else if x == 3 || (x&3) == 1 {
			x--
		} else {
			x++
		}
		steps++
	}
	return steps
}

func main() {
	fmt.Println(integerReplacement(8)) // 3
	fmt.Println(integerReplacement(7)) // 4
	fmt.Println(integerReplacement(4)) // 2
}
