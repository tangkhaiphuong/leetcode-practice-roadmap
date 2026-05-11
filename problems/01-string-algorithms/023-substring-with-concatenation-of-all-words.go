//go:build ignore

// Problem 023: Substring with Concatenation of All Words (LeetCode #30)
// Difficulty: Hard
// Categories: String, Hash Table, Sliding Window
//
// Description:
//   Given string s and a list of words (all the same length L), return all
//   start indices of substrings of length len(words)*L that are formed by
//   concatenating each word in any order, exactly once each.
//
// Examples:
//   s="barfoothefoobarman", words=["foo","bar"] -> [0,9]
//   s="wordgoodgoodgoodbestword", words=["word","good","best","word"] -> []
//
// Approach: Sliding Window of words
//   Each word has length L. For each starting offset in [0, L), slide a
//   window of L-step indices. Maintain a multiset (count map) of words in
//   the window vs the target multiset. Slide one word at a time:
//   add right word; if it overflows or isn't a target, advance left until
//   valid. When window holds exactly len(words) target words, record left.
//
// Complexity:
//   Time:  O(n*L) where n = len(s), L = word length
//   Space: O(K*L), K = len(words)

package main

import "fmt"

func findSubstring(s string, words []string) []int {
	if len(words) == 0 || len(s) == 0 {
		return nil
	}
	L := len(words[0])
	K := len(words)
	target := map[string]int{}
	for _, w := range words {
		target[w]++
	}
	var res []int
	for off := 0; off < L; off++ {
		l, count := off, 0
		window := map[string]int{}
		for r := off; r+L <= len(s); r += L {
			w := s[r : r+L]
			if _, ok := target[w]; !ok {
				window = map[string]int{}
				count = 0
				l = r + L
				continue
			}
			window[w]++
			count++
			for window[w] > target[w] {
				window[s[l:l+L]]--
				count--
				l += L
			}
			if count == K {
				res = append(res, l)
				window[s[l:l+L]]--
				count--
				l += L
			}
		}
	}
	return res
}

func main() {
	fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
	fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}))
	fmt.Println(findSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"}))
}
