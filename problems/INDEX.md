# LeetCode Practice Roadmap — Problem Index

91 unique problems (184 file entries across 17 category folders — many problems are tagged with multiple categories, so they appear in every folder where they belong per the original README).

## How to run a solution

Each `.go` file is self-contained with a `main()` and example test cases. Files use the `//go:build ignore` build tag so multiple files in the same folder don't conflict.

```bash
go run problems/01-string-algorithms/001-longest-palindromic-substring.go
```

## Practice strategy

1. **Read** the problem header (description, examples, constraints).
2. **Attempt first** without looking at the solution — set a timer (Easy: 15m, Medium: 25m, Hard: 45m).
3. **Compare** your approach to the one in the file; understand the complexity tradeoffs.
4. **Re-implement** from scratch the next day (spaced repetition).
5. **Tag** problems you struggled with for a 1-week / 1-month review.

Each category folder has its own README with practice tips, common pitfalls, and a suggested daily order.

---

## Primary location of each problem

Every problem has one "primary" folder (where it canonically lives). Cross-tagged copies in other folders are identical in content. Use this table to find the primary file path.

| # | Problem | Difficulty | Primary folder | Also appears in |
|---|---------|------------|----------------|------------------|
| 001 | Longest Palindromic Substring | Medium | 01-string-algorithms | 04-dp |
| 002 | Integer to English Words | Hard | 01-string-algorithms | 12-math |
| 003 | Compare Version Numbers | Medium | 01-string-algorithms | 05-two-pointers |
| 004 | Longest Substring Without Repeating Characters | Medium | 01-string-algorithms | 03-hash, 09-sliding-window |
| 005 | Find the Closest Palindrome | Hard | 01-string-algorithms | 12-math |
| 006 | Longest Duplicate Substring | Hard | 01-string-algorithms | 06-binary-search, 09-sliding-window |
| 007 | String to Integer (atoi) | Medium | 01-string-algorithms | — |
| 008 | Regular Expression Matching | Hard | 01-string-algorithms | 04-dp |
| 009 | Largest Number | Medium | 01-string-algorithms | 11-greedy |
| 010 | Longest Valid Parentheses | Hard | 01-string-algorithms | 04-dp, 08-stack |
| 011 | Interleaving String | Medium | 01-string-algorithms | 04-dp |
| 012 | Decode Ways | Medium | 01-string-algorithms | 04-dp |
| 013 | Reverse Words in a String | Medium | 01-string-algorithms | 05-two-pointers |
| 014 | Wildcard Matching | Hard | 01-string-algorithms | 04-dp |
| 015 | Valid Parenthesis String | Medium | 01-string-algorithms | 04-dp, 08-stack, 11-greedy |
| 016 | Remove K Digits | Medium | 01-string-algorithms | 07-sorting, 08-stack, 11-greedy |
| 017 | Palindrome Partitioning II | Hard | 01-string-algorithms | 04-dp |
| 018 | Decode Ways II | Hard | 01-string-algorithms | 04-dp |
| 019 | Repeated String Match | Medium | 01-string-algorithms | — |
| 020 | Validate IP Address | Medium | 01-string-algorithms | — |
| 021 | One Edit Distance | Medium | 01-string-algorithms | 05-two-pointers |
| 022 | Valid Number | Hard | 01-string-algorithms | — |
| 023 | Substring with Concatenation of All Words | Hard | 01-string-algorithms | 03-hash, 09-sliding-window |
| 024 | Strong Password Checker | Hard | 01-string-algorithms | 11-greedy, 13-heap |
| 025 | Fraction to Recurring Decimal | Medium | 01-string-algorithms | 03-hash, 12-math |
| 026 | Text Justification | Hard | 01-string-algorithms | 02-array |
| 027 | Shortest Palindrome | Hard | 01-string-algorithms | — |
| 028 | Making File Names Unique | Medium | 01-string-algorithms | 02-array, 03-hash |
| 029 | Decoded String at Index | Medium | 01-string-algorithms | 08-stack |
| 030 | Split Array With Same Average | Hard | 01-string-algorithms | 04-dp, 11-greedy, 15-bit |
| 031 | Count The Repetitions | Hard | 01-string-algorithms | 04-dp |
| 032 | Median of Two Sorted Arrays | Hard | 02-array-manipulation | 06-binary-search |
| 033 | 3Sum | Medium | 02-array-manipulation | 05-two-pointers, 07-sorting |
| 034 | Jump Game II | Medium | 02-array-manipulation | 04-dp |
| 035 | Maximum Product Subarray | Medium | 02-array-manipulation | 04-dp, 11-greedy |
| 036 | Next Permutation | Medium | 02-array-manipulation | 05-two-pointers |
| 037 | 132 Pattern | Medium | 02-array-manipulation | 06-binary-search, 08-stack |
| 038 | Maximum Number of Events | Medium | 02-array-manipulation | 07-sorting, 11-greedy, 13-heap |
| 039 | Surrounded Regions | Medium | 02-array-manipulation | 10-tree-graph |
| 040 | Max Points on a Line | Hard | 02-array-manipulation | 03-hash |
| 041 | Next Greater Element III | Medium | 02-array-manipulation | 12-math |
| 042 | Shortest Unsorted Continuous Subarray | Medium | 02-array-manipulation | 05-two-pointers, 08-stack |
| 043 | Reverse Pairs | Hard | 02-array-manipulation | 06-binary-search, 07-sorting |
| 044 | Shortest Subarray with Sum at Least K | Hard | 02-array-manipulation | 06-binary-search |
| 045 | Shortest Subarray to be Removed | Medium | 02-array-manipulation | — |
| 046 | Water and Jug Problem | Medium | 02-array-manipulation | 10-tree-graph |
| 047 | Self Crossing | Hard | 02-array-manipulation | 12-math |
| 048 | Non-decreasing Array | Medium | 02-array-manipulation | — |
| 049 | Continuous Subarray Sum | Medium | 02-array-manipulation | — |
| 050 | Dungeon Game | Hard | 02-array-manipulation | 04-dp |
| 051 | Wiggle Sort II | Medium | 02-array-manipulation | 07-sorting |
| 052 | Ways to Split Array Into Three Subarrays | Medium | 02-array-manipulation | 05-two-pointers |
| 053 | Maximum Non Negative Product in a Matrix | Medium | 02-array-manipulation | 04-dp |
| 054 | Frequency of the Most Frequent Element | Medium | 02-array-manipulation | 06-binary-search, 07-sorting, 09-sliding-window, 11-greedy |
| 055 | Maximum Number of Visible Points | Hard | 02-array-manipulation | 09-sliding-window, 12-math |
| 056 | Closest Room | Hard | 02-array-manipulation | 06-binary-search, 07-sorting |
| 057 | Minimize Difference Target | Hard | 02-array-manipulation | 04-dp |
| 058 | Word Ladder | Hard | 03-hash-table | — (treat as graph BFS in practice) |
| 059 | Word Ladder II | Hard | 03-hash-table | — (treat as graph BFS in practice) |
| 060 | All O`one Data Structure | Hard | 03-hash-table | 14-linked-list, 16-design |
| 061 | Contains Duplicate III | Hard | 03-hash-table | 09-sliding-window |
| 062 | Invalid Transactions | Medium | 03-hash-table | 07-sorting |
| 063 | Count Good Meals | Medium | 03-hash-table | — |
| 064 | Make Sum Divisible by P | Medium | 03-hash-table | 12-math |
| 065 | Line Reflection | Medium | 03-hash-table | — |
| 066 | Best Time to Buy and Sell Stock IV | Hard | 04-dynamic-programming | — |
| 067 | Maximum Score from Multiplication | Hard | 04-dynamic-programming | — |
| 068 | Super Egg Drop | Hard | 04-dynamic-programming | 06-binary-search |
| 069 | Integer Replacement | Medium | 04-dynamic-programming | 15-bit |
| 070 | Largest Multiple of Three | Hard | 04-dynamic-programming | — |
| 071 | Number of Digit One | Hard | 04-dynamic-programming | 12-math |
| 072 | Sum of Square Numbers | Medium | 05-two-pointers | 12-math |
| 073 | Search in Rotated Sorted Array II | Medium | 06-binary-search | — |
| 074 | Nth Digit | Medium | 06-binary-search | 12-math |
| 075 | Minimum Space Wasted From Packaging | Hard | 06-binary-search | — |
| 076 | Longest Common Subpath | Hard | 06-binary-search | — |
| 077 | Pow(x, n) | Medium | 06-binary-search | 12-math |
| 078 | Create Maximum Number | Hard | 08-stack-monotonic | 11-greedy |
| 079 | Alien Dictionary | Hard | 10-tree-graph | — |
| 080 | Validate Binary Search Tree | Medium | 10-tree-graph | — |
| 081 | Redundant Connection II | Hard | 10-tree-graph | — |
| 082 | Kth Ancestor of a Tree Node | Hard | 10-tree-graph | 16-design |
| 083 | Sell Diminishing-Valued Colored Balls | Medium | 11-greedy | 12-math, 13-heap |
| 084 | Prime Palindrome | Medium | 12-math-number-theory | — |
| 085 | Design Twitter | Medium | 13-heap-priority-queue | 14-linked-list, 16-design |
| 086 | Rotate List | Medium | 14-linked-list | — |
| 087 | Insert into a Sorted Circular Linked List | Medium | 14-linked-list | — |
| 088 | Design Linked List | Medium | 14-linked-list | 16-design |
| 089 | Divide Two Integers | Medium | 15-bit-manipulation | — |
| 090 | Nth Highest Salary (SQL) | Medium | 17-database-sql | 06-binary-search |
| 091 | Word Frequency (Bash) | Medium | 17-database-sql | — |

---

## File counts per folder

| Folder | Files | Notes |
|---|---|---|
| 01-string-algorithms | 31 | All originals |
| 02-array-manipulation | 28 | + 2 cross-tagged (026, 028) |
| 03-hash-table | 13 | + 5 cross-tagged (004, 023, 025, 028, 040) |
| 04-dynamic-programming | 22 | + 16 cross-tagged |
| 05-two-pointers | 8 | + 7 cross-tagged |
| 06-binary-search | 14 | + 8 cross-tagged + 090 SQL |
| 07-sorting | 8 | All cross-tagged |
| 08-stack-monotonic | 7 | + 6 cross-tagged |
| 09-sliding-window | 6 | All cross-tagged |
| 10-tree-graph | 6 | + 2 cross-tagged (039, 046) |
| 11-greedy | 10 | + 9 cross-tagged |
| 12-math-number-theory | 13 | + 12 cross-tagged |
| 13-heap-priority-queue | 4 | + 3 cross-tagged |
| 14-linked-list | 5 | + 2 cross-tagged (060, 085) |
| 15-bit-manipulation | 3 | + 2 cross-tagged (030, 069) |
| 16-design | 4 | All cross-tagged |
| 17-database-sql | 2 | All originals |
| **Total** | **184** | **91 unique problems** |
