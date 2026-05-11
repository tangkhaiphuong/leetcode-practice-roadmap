//go:build ignore

// Problem 031: Count The Repetitions (LeetCode #466)
// Difficulty: Hard
// Categories: String, DP, Cycle Detection
//
// Description:
//   Given strings s1 (length up to 100) and s2, define S1 = [s1, n1] meaning
//   s1 repeated n1 times concatenated. We say "S can be obtained from T" if
//   S can be made by deleting some chars from T. Return the maximum integer M
//   such that [S2, M] (= [s2, n2*M]) can be obtained from S1.
//
//   Equivalently: count how many times s2 repeated n2 times fits as a
//   subsequence inside s1 repeated n1 times.
//
// Examples:
//   s1="acb", n1=4, s2="ab", n2=2 -> 2
//
// Approach: Cycle Detection on s2-Pointer State
//   For each new copy of s1, advance an index j into s2 (and a counter of
//   completed s2's) by walking matching chars. Record state j after each s1
//   copy. After at most |s2| copies of s1, we revisit some j -> a cycle.
//   Use the cycle length and gain to fast-forward through n1 copies, then
//   handle the tail. Final answer = (total s2 completions) / n2.
//
// Complexity:
//   Time:  O(|s1| * min(|s2|, n1))
//   Space: O(|s2|)

package main

import "fmt"

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	if n1 == 0 {
		return 0
	}
	// For state j in [0, len(s2)), record (block_index_when_first_seen, total_s2_count).
	type state struct{ block, count int }
	first := make(map[int]state)
	j := 0
	count := 0
	for i := 1; i <= n1; i++ {
		for k := 0; k < len(s1); k++ {
			if s1[k] == s2[j] {
				j++
				if j == len(s2) {
					j = 0
					count++
				}
			}
		}
		if prev, ok := first[j]; ok {
			cycleBlocks := i - prev.block
			cycleGain := count - prev.count
			remaining := n1 - i
			full := remaining / cycleBlocks
			count += full * cycleGain
			i += full * cycleBlocks
			// Continue with no more cycle handling.
			first = map[int]state{} // clear so we don't loop again
			continue
		}
		first[j] = state{i, count}
	}
	return count / n2
}

func main() {
	fmt.Println(getMaxRepetitions("acb", 4, "ab", 2)) // 2
	fmt.Println(getMaxRepetitions("aaa", 3, "aa", 1)) // 4
}
