//go:build ignore

// Problem 039: Surrounded Regions (LeetCode #130)
// Difficulty: Medium
// Categories: Array (Matrix), Tree/Graph (DFS/BFS), Union Find
//
// Description:
//   Given an m x n board of 'X' and 'O', flip all 'O' regions completely
//   surrounded by 'X' to 'X'. A region is captured if no 'O' in it touches
//   the border.
//
// Examples:
//   X X X X       X X X X
//   X O O X  ->   X X X X
//   X X O X       X X X X
//   X O X X       X O X X
//
// Approach: Border-Connected DFS
//   Any 'O' connected (4-directionally) to a border 'O' is safe. Mark all
//   such 'O's by DFS/BFS from every border 'O'; flip them to a placeholder
//   like '#'. Then in a single pass: '#' -> 'O', everything else -> 'X'.
//
// Complexity:
//   Time:  O(m*n)
//   Space: O(m*n) recursion (BFS to avoid stack overflow on huge grids).

package main

import "fmt"

func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}
	m, n := len(board), len(board[0])
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= m || c < 0 || c >= n || board[r][c] != 'O' {
			return
		}
		board[r][c] = '#'
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}
	for i := 0; i < m; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}
	for j := 0; j < n; j++ {
		dfs(0, j)
		dfs(m-1, j)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch board[i][j] {
			case '#':
				board[i][j] = 'O'
			case 'O':
				board[i][j] = 'X'
			}
		}
	}
}

func main() {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	solve(board)
	for _, row := range board {
		fmt.Println(string(row))
	}
}
