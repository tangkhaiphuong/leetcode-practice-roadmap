//go:build ignore

// Problem 002: Integer to English Words (LeetCode #273)
// Difficulty: Hard
// Categories: String, Math
//
// Description:
//   Convert a non-negative integer num to its English words representation.
//
// Examples:
//   123                  -> "One Hundred Twenty Three"
//   12345                -> "Twelve Thousand Three Hundred Forty Five"
//   1234567              -> "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"
//   1234567891           -> "One Billion Two Hundred Thirty Four Million Five Hundred Sixty Seven Thousand Eight Hundred Ninety One"
//
// Constraints:
//   0 <= num <= 2^31 - 1
//
// Approach: Group-by-Thousand Recursion
//   English organizes large numbers in groups of three digits, named by
//   thousands powers: "", Thousand, Million, Billion. So:
//     - Split num into 3-digit chunks from least to most significant.
//     - Convert each chunk (0..999) to words.
//     - Append the appropriate scale word.
//   For each <1000 chunk, handle hundreds, then 20-99, then 1-19 specially
//   (English has irregular names for 11..19).
//
// Complexity:
//   Time:  O(1) — input is 32-bit (at most 4 chunks).
//   Space: O(1)

package main

import (
	"fmt"
	"strings"
)

var (
	below20 = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven",
		"Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen",
		"Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	tens   = []string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
	scales = []string{"", "Thousand", "Million", "Billion"}
)

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	var parts []string
	for i := 0; num > 0; i++ {
		if chunk := num % 1000; chunk != 0 {
			w := chunkToWords(chunk)
			if scales[i] != "" {
				w += " " + scales[i]
			}
			parts = append([]string{w}, parts...)
		}
		num /= 1000
	}
	return strings.Join(parts, " ")
}

func chunkToWords(n int) string {
	var b []string
	if n >= 100 {
		b = append(b, below20[n/100], "Hundred")
		n %= 100
	}
	if n >= 20 {
		b = append(b, tens[n/10])
		n %= 10
		if n > 0 {
			b = append(b, below20[n])
		}
	} else if n > 0 {
		b = append(b, below20[n])
	}
	return strings.Join(b, " ")
}

func main() {
	tests := []int{0, 5, 13, 20, 100, 123, 12345, 1234567, 1234567891}
	for _, n := range tests {
		fmt.Printf("%-12d -> %q\n", n, numberToWords(n))
	}
}
