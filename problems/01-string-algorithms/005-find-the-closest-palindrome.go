//go:build ignore

// Problem 005: Find the Closest Palindrome (LeetCode #564)
// Difficulty: Hard
// Categories: String, Math
//
// Description:
//   Given an integer n represented as a string, return the closest integer
//   (not equal to n) that is a palindrome. If there's a tie, return the
//   smaller one.
//
// Examples:
//   "123" -> "121"
//   "1"   -> "0"
//   "11"  -> "9"   (must not equal n; "22" is farther than "9")
//
// Constraints:
//   1 <= n.length <= 18, no leading zeros, n is a non-negative integer.
//
// Approach: Mirror Half + Edge Candidates
//   The closest palindrome is one of just FIVE candidates:
//     1. Mirror the left half onto the right (the obvious palindrome at this length).
//     2. Mirror after incrementing the left half (covers next palindrome up).
//     3. Mirror after decrementing the left half (next palindrome down).
//     4. 10^(L-1) - 1, e.g. 999...9 (one digit shorter, max palindrome below).
//     5. 10^L + 1,     e.g. 1000...01 (one digit longer, min palindrome above).
//   Among candidates not equal to n, pick the one minimizing |cand - n|; on
//   tie, the smaller. Edge candidates 4 & 5 cover length-boundary cases like
//   "1000" -> "999".
//
// Complexity:
//   Time:  O(L) string ops, where L = len(n)
//   Space: O(L)

package main

import (
	"fmt"
	"strconv"
)

func nearestPalindromic(n string) string {
	L := len(n)
	num, _ := strconv.ParseInt(n, 10, 64)

	// Edge palindromes: e.g. 999...9 and 100..001.
	candidates := map[int64]bool{}
	pow := int64(1)
	for i := 0; i < L-1; i++ {
		pow *= 10
	}
	candidates[pow*10+1] = true // 10^L + 1
	candidates[pow-1] = true    // 10^(L-1) - 1

	// Mirror around the left-half (rounded up).
	half, _ := strconv.ParseInt(n[:(L+1)/2], 10, 64)
	for _, h := range []int64{half - 1, half, half + 1} {
		candidates[mirror(h, L%2 == 1)] = true
	}
	delete(candidates, num)

	var best int64 = -1
	for c := range candidates {
		if c < 0 {
			continue
		}
		if best == -1 || abs(c-num) < abs(best-num) ||
			(abs(c-num) == abs(best-num) && c < best) {
			best = c
		}
	}
	return strconv.FormatInt(best, 10)
}

// mirror builds a palindrome from its left half. If oddLen, the last char of
// the half is the center (not duplicated). e.g. half=12, oddLen=true -> 121.
func mirror(half int64, oddLen bool) int64 {
	s := strconv.FormatInt(half, 10)
	end := len(s)
	if oddLen {
		end-- // skip the center digit when reversing
	}
	for i := end - 1; i >= 0; i-- {
		s += string(s[i])
	}
	v, _ := strconv.ParseInt(s, 10, 64)
	return v
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	for _, n := range []string{"123", "1", "11", "10", "1000", "12932"} {
		fmt.Printf("nearest(%q) = %q\n", n, nearestPalindromic(n))
	}
}
