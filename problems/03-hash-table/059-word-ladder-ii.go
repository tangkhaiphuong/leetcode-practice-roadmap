//go:build ignore

// Problem 059: Word Ladder II (LeetCode #126)
// Difficulty: Hard
// Categories: Hash Table, Tree/Graph (BFS), Backtracking
//
// Description:
//   Same setup as Word Ladder; return ALL shortest transformation sequences.
//
// Approach: BFS to Build Parent Graph + DFS to Reconstruct
//   Run BFS layer by layer; track parents[word] = list of predecessors found
//   at the previous shortest level. Once endWord first appears, stop adding
//   new words but continue the current level so all parents are recorded.
//   Then DFS from endWord up via parents, prepending words.
//
// Complexity:
//   Time:  O(N*L^2 + paths) — last term can be exponential.
//   Space: O(N*L^2 + paths)

package main

import "fmt"

func findLadders(begin, end string, wordList []string) [][]string {
	dict := map[string]bool{}
	for _, w := range wordList {
		dict[w] = true
	}
	if !dict[end] {
		return nil
	}
	L := len(begin)
	parents := map[string][]string{}
	level := map[string]bool{begin: true}
	visited := map[string]bool{begin: true}
	found := false
	for len(level) > 0 && !found {
		next := map[string]bool{}
		for w := range level {
			bytes := []byte(w)
			for i := 0; i < L; i++ {
				orig := bytes[i]
				for c := byte('a'); c <= 'z'; c++ {
					if c == orig {
						continue
					}
					bytes[i] = c
					nw := string(bytes)
					if dict[nw] && !visited[nw] {
						if nw == end {
							found = true
						}
						next[nw] = true
						parents[nw] = append(parents[nw], w)
					}
				}
				bytes[i] = orig
			}
		}
		for w := range next {
			visited[w] = true
		}
		level = next
	}
	if !found {
		return nil
	}
	var res [][]string
	var dfs func(w string, path []string)
	dfs = func(w string, path []string) {
		if w == begin {
			full := append([]string{begin}, path...)
			res = append(res, full)
			return
		}
		for _, p := range parents[w] {
			dfs(p, append([]string{w}, path...))
		}
	}
	dfs(end, nil)
	return res
}

func main() {
	out := findLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})
	for _, p := range out {
		fmt.Println(p)
	}
}
