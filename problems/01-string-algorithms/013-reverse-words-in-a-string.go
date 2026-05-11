//go:build ignore

// Problem 013: Reverse Words in a String (LeetCode #151)
// Difficulty: Medium
// Categories: String, Two Pointers
//
// Description:
//   Reverse the order of words in s. Words are non-space sequences. Trim
//   leading/trailing spaces. Collapse multiple internal spaces to one.
//
// Examples:
//   "the sky is blue"        -> "blue is sky the"
//   "  hello world  "        -> "world hello"
//   "a good   example"       -> "example good a"
//
// Approach: In-place via byte slice (O(1) extra)
//   1. Reverse the entire byte slice.
//   2. Reverse each word, collapsing spaces by writing back to a tail pointer.
//
//   For a clearer Go solution we also include the slice/strings.Fields version.
//
// Complexity:
//   Time:  O(n)
//   Space: O(n) (immutable strings in Go).

package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	fields := strings.Fields(s) // splits on whitespace runs and trims
	for i, j := 0, len(fields)-1; i < j; i, j = i+1, j-1 {
		fields[i], fields[j] = fields[j], fields[i]
	}
	return strings.Join(fields, " ")
}

func main() {
	for _, s := range []string{"the sky is blue", "  hello world  ", "a good   example", "  "} {
		fmt.Printf("reverseWords(%q) = %q\n", s, reverseWords(s))
	}
}
