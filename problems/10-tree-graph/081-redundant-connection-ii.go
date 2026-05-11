//go:build ignore

// Problem 081: Redundant Connection II (LeetCode #685)
// Difficulty: Hard
// Categories: Tree, Graph, Union Find
//
// Description:
//   A directed graph that started as a rooted tree (every node has exactly
//   one parent except the root) had ONE extra edge added. The result has n
//   nodes (1..n) and n directed edges. Return the redundant edge to remove
//   so the graph becomes a rooted tree. If multiple, return the last in input.
//
// Examples:
//   [[1,2],[1,3],[2,3]]            -> [2,3]   (node 3 has two parents)
//   [[1,2],[2,3],[3,4],[4,1],[1,5]]-> [4,1]
//
// Approach: Three Cases
//   Two failure modes can occur, separately or together:
//     A) Some node has two parents (in-degree 2). Detect by tracking parent[].
//     B) The directed graph contains a cycle.
//   Algorithm:
//     - First pass: find a node with two parents — remember the two edges
//       (cand1: earlier, cand2: later).
//     - Skip cand2 and try Union-Find building tree:
//         - if no cycle: cand2 is the redundant edge.
//         - if cycle: cand1 is the redundant edge (case B alone or in
//           combination with A).
//     - If no two-parents conflict, find the cycle edge with normal Union-Find.
//
// Complexity:
//   Time:  O(n α(n))
//   Space: O(n)

package main

import "fmt"

func findRedundantDirectedConnection(edges [][]int) []int {
	n := len(edges)
	parent := make([]int, n+1)
	var cand1, cand2 []int
	for _, e := range edges {
		u, v := e[0], e[1]
		if parent[v] != 0 {
			cand1 = []int{parent[v], v}
			cand2 = []int{u, v}
			e[1] = 0 // mark to skip later
			break
		}
		parent[v] = u
	}

	uf := make([]int, n+1)
	for i := range uf {
		uf[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if uf[x] != x {
			uf[x] = find(uf[x])
		}
		return uf[x]
	}

	for _, e := range edges {
		if e[1] == 0 {
			continue
		}
		u, v := e[0], e[1]
		if find(u) == find(v) {
			if cand1 == nil {
				return e
			}
			return cand1
		}
		uf[find(u)] = find(v)
	}
	return cand2
}

func main() {
	fmt.Println(findRedundantDirectedConnection([][]int{{1, 2}, {1, 3}, {2, 3}}))           // [2 3]
	fmt.Println(findRedundantDirectedConnection([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 5}})) // [4 1]
}
