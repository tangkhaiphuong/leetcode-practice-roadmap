//go:build ignore

// Problem 058: Word Ladder (LeetCode #127)
// Difficulty: Hard
// Categories: Hash Table, Tree/Graph (BFS), String
//
// Description:
//   Given two words beginWord, endWord, and a wordList, return the length of
//   the shortest transformation sequence from beginWord to endWord, where
//   each step changes exactly one letter and every intermediate word is in
//   wordList. Return 0 if no such sequence.
//
// Examples:
//   "hit", "cog", ["hot","dot","dog","lot","log","cog"] -> 5
//   "hit", "cog", ["hot","dot","dog","lot","log"]       -> 0
//
// Approach: BFS using Wildcard Buckets
//   Group words by patterns like "h*t", "*ot", "ho*" — neighbors are words
//   that share at least one such pattern.
//   BFS from beginWord. At each level, mark visited; return level when
//   endWord reached.
//
// Complexity:
//   Time:  O(N * L^2) where N = words, L = word length
//   Space: O(N * L^2)

package main

import "fmt"

func ladderLength(begin, end string, wordList []string) int {
	dict := map[string]bool{}
	for _, w := range wordList {
		dict[w] = true
	}
	if !dict[end] {
		return 0
	}
	// Build pattern adjacency
	L := len(begin)
	patterns := map[string][]string{}
	for w := range dict {
		for i := 0; i < L; i++ {
			p := w[:i] + "*" + w[i+1:]
			patterns[p] = append(patterns[p], w)
		}
	}
	visited := map[string]bool{begin: true}
	queue := []string{begin}
	level := 1
	for len(queue) > 0 {
		var next []string
		for _, w := range queue {
			if w == end {
				return level
			}
			for i := 0; i < L; i++ {
				p := w[:i] + "*" + w[i+1:]
				for _, nb := range patterns[p] {
					if !visited[nb] {
						visited[nb] = true
						next = append(next, nb)
					}
				}
				patterns[p] = nil // avoid re-traversal
			}
		}
		queue = next
		level++
	}
	return 0
}

func main() {
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})) // 5
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}))        // 0
}
