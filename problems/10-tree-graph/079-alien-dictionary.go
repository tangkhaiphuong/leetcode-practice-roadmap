//go:build ignore

// Problem 079: Alien Dictionary (LeetCode #269)
// Difficulty: Hard
// Categories: Tree/Graph, Topological Sort, BFS
//
// Description:
//   Given a list of words sorted lexicographically by the rules of an alien
//   alphabet, return any valid ordering of the alphabet's letters. If there
//   are multiple, return any. If invalid, return "".
//
// Examples:
//   ["wrt","wrf","er","ett","rftt"] -> "wertf"
//   ["z","x"]                       -> "zx"
//   ["z","x","z"]                   -> ""  (cycle)
//
// Approach: Build Graph + Kahn's Topological Sort
//   For each adjacent pair of words, find the first differing char (a,b);
//   add edge a -> b. Edge case: if word1 is a prefix of word2 in the wrong
//   way (longer is prefix of shorter), invalid.
//   Compute indegree, BFS topo. If result has fewer than #unique chars,
//   cycle exists.
//
// Complexity:
//   Time:  O(C) where C = total characters across words
//   Space: O(unique letters)

package main

import "fmt"

func alienOrder(words []string) string {
	indeg := map[byte]int{}
	graph := map[byte]map[byte]bool{}
	for _, w := range words {
		for i := 0; i < len(w); i++ {
			indeg[w[i]] = 0
			if graph[w[i]] == nil {
				graph[w[i]] = map[byte]bool{}
			}
		}
	}
	for i := 0; i+1 < len(words); i++ {
		a, b := words[i], words[i+1]
		k := 0
		for k < len(a) && k < len(b) && a[k] == b[k] {
			k++
		}
		if k == len(b) && len(a) > len(b) {
			return ""
		}
		if k < len(a) && k < len(b) {
			x, y := a[k], b[k]
			if !graph[x][y] {
				graph[x][y] = true
				indeg[y]++
			}
		}
	}
	queue := []byte{}
	for c, d := range indeg {
		if d == 0 {
			queue = append(queue, c)
		}
	}
	out := []byte{}
	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]
		out = append(out, c)
		for nb := range graph[c] {
			indeg[nb]--
			if indeg[nb] == 0 {
				queue = append(queue, nb)
			}
		}
	}
	if len(out) != len(indeg) {
		return ""
	}
	return string(out)
}

func main() {
	fmt.Println(alienOrder([]string{"wrt", "wrf", "er", "ett", "rftt"})) // wertf (one valid)
	fmt.Println(alienOrder([]string{"z", "x"}))                          // zx
	fmt.Println(alienOrder([]string{"z", "x", "z"}))                     // ""
}
