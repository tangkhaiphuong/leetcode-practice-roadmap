//go:build ignore

// Problem 024: Strong Password Checker (LeetCode #420)
// Difficulty: Hard
// Categories: String, Greedy, Heap
//
// Description:
//   A password is "strong" if:
//     1. length in [6, 20]
//     2. contains at least one lowercase, one uppercase, one digit
//     3. has no three repeating chars in a row (e.g. "aaa" forbidden)
//   Return the minimum number of edits (insert, delete, replace one char)
//   needed to make password strong.
//
// Examples:
//   "a"          -> 5 (need length 6 + classes)
//   "aA1"        -> 3
//   "1337C0d3"   -> 0
//   "aaaB2c"     -> 1 (replace one a)
//
// Approach: Case Analysis on Length
//   Let missing = 3 - (hasLower? 1:0) - (hasUpper? 1:0) - (hasDigit? 1:0).
//   Count "runs" of >=3 repeated chars; for each run of length r, replacements
//   needed to break it = r/3.
//   Three regimes:
//     - len < 6: max(missing, 6 - len) — inserts cover both length and missing
//     - 6 <= len <= 20: max(missing, totalReplacements)
//     - len > 20: must delete excess. Reduce replacements by greedily deleting
//       in runs of length r%3==0 first, then r%3==1, then any. Total =
//       deletes + max(missing, residualReplacements).
//
// Complexity:
//   Time:  O(n)
//   Space: O(n) for run lengths

package main

import "fmt"

func strongPasswordChecker(password string) int {
	n := len(password)
	hasLower, hasUpper, hasDigit := 0, 0, 0
	for i := 0; i < n; i++ {
		c := password[i]
		switch {
		case c >= 'a' && c <= 'z':
			hasLower = 1
		case c >= 'A' && c <= 'Z':
			hasUpper = 1
		case c >= '0' && c <= '9':
			hasDigit = 1
		}
	}
	missing := 3 - hasLower - hasUpper - hasDigit

	// Collect run lengths of repeated chars.
	var runs []int
	for i := 0; i < n; {
		j := i
		for j < n && password[j] == password[i] {
			j++
		}
		if j-i >= 3 {
			runs = append(runs, j-i)
		}
		i = j
	}

	if n < 6 {
		need := 6 - n
		if need > missing {
			return need
		}
		return missing
	}

	totalReplaces := 0
	for _, r := range runs {
		totalReplaces += r / 3
	}
	if n <= 20 {
		if missing > totalReplaces {
			return missing
		}
		return totalReplaces
	}

	// n > 20: must delete (n-20). Use deletions to reduce replacements first.
	delete := n - 20
	leftover := delete
	// Pass 1: r%3==0 — saves 1 replacement per 1 deletion
	for i := range runs {
		if leftover == 0 {
			break
		}
		if runs[i] >= 3 && runs[i]%3 == 0 {
			runs[i]--
			leftover--
		}
	}
	// Pass 2: r%3==1 — saves 1 per 2 deletions
	for i := range runs {
		if leftover < 2 {
			break
		}
		if runs[i] >= 3 && runs[i]%3 == 1 {
			runs[i] -= 2
			leftover -= 2
		}
	}
	// Pass 3: any — saves 1 per 3 deletions
	for i := range runs {
		for leftover >= 3 && runs[i] >= 3 {
			runs[i] -= 3
			leftover -= 3
		}
	}
	residual := 0
	for _, r := range runs {
		residual += r / 3
	}
	if missing > residual {
		return delete + missing
	}
	return delete + residual
}

func main() {
	for _, p := range []string{"a", "aA1", "1337C0d3", "aaaB2c", "aaaaaaaaaaaaaaaaaaaaaaaaA0", "ABABABABABABABABABAB1"} {
		fmt.Printf("strong(%q) = %d\n", p, strongPasswordChecker(p))
	}
}
