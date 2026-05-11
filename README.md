# LeetCode Practice Roadmap: 90+ High-Impact Problems

> **Credit & Source:** The 90+ problem roadmap below is from **[Kartikey Kumar Srivastava](https://www.linkedin.com/in/kartikey-kumar-srivastava-3969b8aa/)** — shared as his personal interview-prep notes on LinkedIn. All problem selections, categorizations, and the curatorial framing originate from his post. This repository contains my Go solutions and per-category practice notes built on top of his roadmap.

Original author's intro (paraphrased from his LinkedIn post):

> In 2020, I cracked Amazon as an SDE.
> In 2022, I cracked Google as an SWE.
> In 2024, I cracked LinkedIn, Microsoft, as a Sr. SWE.
> In 2025, I am working as a Sr. Software Engineer at Meta.
>
> If I had to restart doing LeetCode practice and build my basics from zero in 2025, I wouldn’t solve the 3500+ problems available on LeetCode. Instead, I would 100% solve these 90+ problems from my own notes.

---

## 📂 Solutions in Go

Every problem is implemented in Go with a full template (problem statement, examples, approach, complexity, runnable solution + tests). Files live under [`problems/`](./problems/), organized by primary category.

**Start here:**
- 🗺 [`problems/INDEX.md`](./problems/INDEX.md) — full problem index with difficulty, cross-category tags, and number→folder lookup.

**Run any solution:**
```bash
go run problems/01-string-algorithms/001-longest-palindromic-substring.go
```

**Categories** (each folder has its own README with patterns, pitfalls, and a suggested practice order; problems with multiple tags appear in every relevant folder):

| # | Folder | Files | What you'll practice |
|---|---|---|---|
| 01 | [String & String Algorithms](./problems/01-string-algorithms/) | 31 | Palindromes, parsing, parens, regex, KMP, rolling hash |
| 02 | [Array Manipulation](./problems/02-array-manipulation/) | 28 | Two pointers, prefix sums, monotonic deques, matrix DP |
| 03 | [Hash Table](./problems/03-hash-table/) | 13 | Last-seen maps, BFS on implicit graphs, bucketing |
| 04 | [Dynamic Programming](./problems/04-dynamic-programming/) | 22 | State machines, matching DP, dimension flips, digit DP |
| 05 | [Two Pointers](./problems/05-two-pointers/) | 8 | Opposing & same-direction pointers, pivot+reverse |
| 06 | [Binary Search](./problems/06-binary-search/) | 14 | Search the answer, fast exponentiation, partition search |
| 07 | [Sorting & Order](./problems/07-sorting/) | 8 | Sort+two-pointer, sort+window, merge-sort counting |
| 08 | [Stack & Monotonic Stack](./problems/08-stack-monotonic/) | 7 | Monotonic stacks, sentinels, decoded-string virtualization |
| 09 | [Sliding Window](./problems/09-sliding-window/) | 6 | Variable/fixed windows, word multisets, rolling hash |
| 10 | [Tree & Graph](./problems/10-tree-graph/) | 6 | BST validation, topo sort, Union Find, binary lifting |
| 11 | [Greedy](./problems/11-greedy/) | 10 | Comparator-based, monotonic, strata, range tracking |
| 12 | [Math & Number Theory](./problems/12-math-number-theory/) | 13 | GCD, modular arithmetic, place-value, digit DP, geometry |
| 13 | [Heap / Priority Queue](./problems/13-heap-priority-queue/) | 4 | Top-K, multi-stream merge, day-sweep + heap |
| 14 | [Linked List](./problems/14-linked-list/) | 5 | Sentinels, two-pointer tricks, circular lists, DLL buckets |
| 15 | [Bit Manipulation](./problems/15-bit-manipulation/) | 3 | Repeated doubling, bitmask DP, bit-pattern greedy |
| 16 | [Design / System Design](./problems/16-design/) | 4 | DLL buckets, sentinel-headed LL, binary lifting, heap merge |
| 17 | [Database / SQL](./problems/17-database-sql/) | 2 | Window functions, LIMIT/OFFSET, bash pipelines |

> **91 unique problems**, **184 file entries** (problems are duplicated into every category they belong to per the README's classification). Open the [INDEX](./problems/INDEX.md) to find a problem's primary folder and all cross-category locations.

## 📚 Suggested study plan

| Week | Focus | Folders |
|---|---|---|
| 1 | Strings + Arrays warmup | 01, 02 (easy/medium first) |
| 2 | Strings + Arrays hards | 01, 02 (hards) |
| 3 | DP fluency | 04 + cross-tagged DP in 01 and 02 |
| 4 | Hash, Two Pointers, Binary Search | 03, 05, 06 |
| 5 | Stack, Greedy, Heap | 08, 11, 13 (and their cross-tagged peers) |
| 6 | Trees & Graphs | 10 + cross-tagged BFS/DFS in 02, 03 |
| 7 | Math, Bit, Linked List, SQL | 12, 14, 15, 17 |
| 8 | Spaced-repetition revisit on weak areas | — |

## 🎯 How to practice each problem

1. **Read** the problem header in the file (description, examples, constraints).
2. **Attempt cold** — set a timer (Easy: 15m, Medium: 25m, Hard: 45m). No peeking.
3. **Compare** your approach to the one in the file. Understand the complexity tradeoffs.
4. **Re-implement from scratch** the next day (spaced repetition).
5. **Tag** problems you struggled with for a 1-week and 1-month review.

---

## 📋 Full problem list by category

> The list below is the original roadmap from **[Kartikey Kumar Srivastava's LinkedIn post](https://www.linkedin.com/in/kartikey-kumar-srivastava-3969b8aa/)**. Use [INDEX.md](./problems/INDEX.md) for clickable navigation to each solution.

---

## String & String Algorithms

- Longest Palindromic Substring
- Integer to English Words
- Compare Version Numbers
- Longest Substring Without Repeating Characters
- Find the Closest Palindrome
- Longest Duplicate Substring
- String to Integer (atoi)
- Regular Expression Matching
- Largest Number
- Longest Valid Parentheses
- Interleaving String
- Decode Ways
- Reverse Words in a String
- Wildcard Matching
- Valid Parenthesis String
- Remove K Digits
- Palindrome Partitioning II
- Decode Ways II
- Repeated String Match
- Validate IP Address
- One Edit Distance
- Valid Number
- Substring with Concatenation of All Words
- Strong Password Checker
- Fraction to Recurring Decimal
- Text Justification
- Shortest Palindrome
- Making File Names Unique
- Decoded String at Index
- Split Array With Same Average
- Count The Repetitions

---

## Array & Array Manipulation

- Median of Two Sorted Arrays
- 3Sum
- Jump Game II
- Maximum Product Subarray
- Next Permutation
- 132 Pattern
- Maximum Number of Events That Can Be Attended
- Surrounded Regions
- Max Points on a Line
- Next Greater Element III
- Text Justification
- Shortest Unsorted Continuous Subarray
- Reverse Pairs
- Shortest Subarray with Sum at Least K
- Shortest Subarray to be Removed to Make Array Sorted
- Water and Jug Problem
- Making File Names Unique
- Self Crossing
- Non-decreasing Array
- Continuous Subarray Sum
- Dungeon Game
- Wiggle Sort II
- Ways to Split Array Into Three Subarrays
- Maximum Non Negative Product in a Matrix
- Frequency of the Most Frequent Element
- Maximum Number of Visible Points
- Closest Room
- Minimize the Difference Between Target and Chosen Elements

---

## Hash Table & Hashing

- Word Ladder
- Word Ladder II
- Longest Substring Without Repeating Characters
- Fraction to Recurring Decimal
- Making File Names Unique
- Substring with Concatenation of All Words
- Max Points on a Line
- All O`one Data Structure
- Contains Duplicate III
- Invalid Transactions
- Count Good Meals
- Make Sum Divisible by P
- Line Reflection

---

## Dynamic Programming

- Longest Palindromic Substring
- Jump Game II
- Maximum Product Subarray
- Regular Expression Matching
- Best Time to Buy and Sell Stock IV
- Longest Valid Parentheses
- Interleaving String
- Decode Ways
- Wildcard Matching
- Valid Parenthesis String
- Dungeon Game
- Count The Repetitions
- Palindrome Partitioning II
- Decode Ways II
- Maximum Score from Performing Multiplication Operations
- Super Egg Drop
- Integer Replacement
- Split Array With Same Average
- Largest Multiple of Three
- Minimize the Difference Between Target and Chosen Elements
- Maximum Non Negative Product in a Matrix
- Number of Digit One

---

## Two Pointers Technique

- Compare Version Numbers
- 3Sum
- Next Permutation
- Reverse Words in a String
- One Edit Distance
- Shortest Unsorted Continuous Subarray
- Ways to Split Array Into Three Subarrays
- Sum of Square Numbers

---

## Binary Search

- Median of Two Sorted Arrays
- 132 Pattern
- Search in Rotated Sorted Array II
- Nth Digit
- Shortest Subarray with Sum at Least K
- Reverse Pairs
- Closest Room
- Nth Highest Salary
- Super Egg Drop
- Minimum Space Wasted From Packaging
- Frequency of the Most Frequent Element
- Longest Duplicate Substring
- Longest Common Subpath
- Pow(x, n)

---

## Sorting & Order

- 3Sum
- Wiggle Sort II
- Maximum Number of Events That Can Be Attended
- Reverse Pairs
- Remove K Digits
- Invalid Transactions
- Frequency of the Most Frequent Element
- Closest Room

---

## Stack & Monotonic Stack

- 132 Pattern
- Longest Valid Parentheses
- Valid Parenthesis String
- Remove K Digits
- Create Maximum Number
- Shortest Unsorted Continuous Subarray
- Decoded String at Index

---

## Sliding Window

- Longest Substring Without Repeating Characters
- Substring with Concatenation of All Words
- Contains Duplicate III
- Longest Duplicate Substring
- Frequency of the Most Frequent Element
- Maximum Number of Visible Points

---

## Tree & Graph

Including BFS, DFS, and Topological Sort.

- Alien Dictionary
- Validate Binary Search Tree
- Redundant Connection II
- Kth Ancestor of a Tree Node
- Surrounded Regions
- Water and Jug Problem

---

## Greedy Algorithms

- Maximum Product Subarray
- Largest Number
- Strong Password Checker
- Valid Parenthesis String
- Remove K Digits
- Maximum Number of Events That Can Be Attended
- Split Array With Same Average
- Create Maximum Number
- Frequency of the Most Frequent Element
- Sell Diminishing-Valued Colored Balls

---

## Math & Number Theory

- Integer to English Words
- Find the Closest Palindrome
- Prime Palindrome
- Pow(x, n)
- Fraction to Recurring Decimal
- Next Greater Element III
- Sum of Square Numbers
- Nth Digit
- Self Crossing
- Maximum Number of Visible Points
- Make Sum Divisible by P
- Number of Digit One
- Sell Diminishing-Valued Colored Balls

---

## Heap / Priority Queue

- Maximum Number of Events That Can Be Attended
- Sell Diminishing-Valued Colored Balls
- Design Twitter
- Strong Password Checker

---

## Linked List

- Rotate List
- All O`one Data Structure
- Insert into a Sorted Circular Linked List
- Design Linked List
- Design Twitter

---

## Bit Manipulation / Bitmask

- Divide Two Integers
- Split Array With Same Average
- Integer Replacement

---

## Design / System Design

- All O`one Data Structure
- Design Twitter
- Design Linked List
- Kth Ancestor of a Tree Node

---

## Database / SQL

- Nth Highest Salary
- Word Frequency

---

## Other Advanced Patterns

### Rolling Hash / Suffix Array / Hash Function

- Longest Duplicate Substring
- Shortest Palindrome
- Longest Common Subpath

### Divide and Conquer

- Median of Two Sorted Arrays
- Reverse Pairs
- Wiggle Sort II

### Union Find

- Surrounded Regions
- Redundant Connection II

### Geometry

- Max Points on a Line
- Self Crossing
- Maximum Number of Visible Points


