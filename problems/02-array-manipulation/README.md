# Array Manipulation

28 problems on indexing tricks, prefix sums, partitioning, monotonic structures, matrix DP, and 2D geometry.

## Core patterns

| Pattern | When to use | Problems in this folder |
|---|---|---|
| **Two pointers** | Sorted arrays, partitioning, in-place ops | 033, 036, 042, 052 |
| **Binary search** | Sorted / monotonic predicate over array | 032, 037, 043, 044, 056 |
| **Prefix sum + hash** | Subarray with sum/mod property | 049, 052 |
| **Monotonic deque** | Min/max in window, shortest sum-≥-K | 044 |
| **Sweep / event sorting** | Intervals, attendance scheduling | 038 |
| **Kadane-style DP** | Best subarray ending here | 035, 053 |
| **Quickselect / partition** | Top-K, wiggle, Dutch flag | 051 |
| **Matrix DP** | Path constraints, min cost / max product | 050, 053 |
| **Geometry / slopes** | Collinearity, visibility, self-crossing | 040, 047, 055 |
| **Greedy line packing** | Text justification | 026 |

## Common interview pitfalls

- **Sort-then-two-pointer dedup** — after a hit, advance BOTH pointers past duplicates *and* the outer index past duplicates. Forgetting one creates duplicate triplets in 3Sum.
- **Off-by-one in binary search** — `lo <= hi` vs `lo < hi`; `hi = mid` vs `hi = mid - 1`. Pick a template and stick with it.
- **Median-of-two-arrays** — partition on the SHORTER array to keep `j` in range.
- **Negative numbers in product subarray** — track BOTH min and max (a very negative min can become the new max).
- **Modular arithmetic** — when problem says "modulo 1e9+7", apply mod after EVERY add/multiply, including in helpers like prefix sums.
- **Matrix bounds** — when the recurrence uses `dp[i+1][j]`, your sentinel row/col must hold a value that doesn't pollute `min`/`max`.
- **`sort.Slice` is unstable** — use `sort.SliceStable` if order of ties matters.

## Suggested practice order

1. **Classics** (must-know): 032, 033, 034, 035, 036
2. **Stack/two-pointer** (Day 2): 037, 042, 045, 048
3. **Prefix-sum & hash** (Day 3): 049, 052
4. **Monotonic deque & merge sort** (Day 4): 043, 044
5. **Matrix DP** (Day 5): 050, 053, 057
6. **Geometry & advanced** (Day 6+): 040, 047, 055, 056

## Problems in this folder

| # | Problem | Difficulty | Key idea |
|---|---|---|---|
| 026 | Text Justification | Hard | Two-pass per line (greedy + padding) |
| 028 | Making File Names Unique | Medium | Map of next suffix to try |
| 032 | Median of Two Sorted Arrays | Hard | Binary search partition on shorter |
| 033 | 3Sum | Medium | Sort + two pointers + dedup |
| 034 | Jump Game II | Medium | BFS layer (greedy range) |
| 035 | Maximum Product Subarray | Medium | Track min AND max ending here |
| 036 | Next Permutation | Medium | Pivot, swap, reverse tail |
| 037 | 132 Pattern | Medium | Right-to-left monotonic stack |
| 038 | Maximum Number of Events | Medium | Day-sweep + min-heap by end |
| 039 | Surrounded Regions | Medium | Border DFS marker pass |
| 040 | Max Points on a Line | Hard | GCD-reduced slope hashing |
| 041 | Next Greater Element III | Medium | Next permutation on digits |
| 042 | Shortest Unsorted Continuous Subarray | Medium | Single-pass min/max window |
| 043 | Reverse Pairs | Hard | Merge sort with cross-pair count |
| 044 | Shortest Subarray with Sum at Least K | Hard | Monotonic deque on prefix sums |
| 045 | Shortest Subarray to be Removed | Medium | Two pointers from each end |
| 046 | Water and Jug Problem | Medium | Bezout's identity (GCD) |
| 047 | Self Crossing | Hard | Three geometric cases |
| 048 | Non-decreasing Array | Medium | Single repair branch |
| 049 | Continuous Subarray Sum | Medium | Prefix sum mod k + first index map |
| 050 | Dungeon Game | Hard | Bottom-up DP from destination |
| 051 | Wiggle Sort II | Medium | Sort + index rewiring (reversed halves) |
| 052 | Ways to Split Array Into Three | Medium | Prefix sums + two-pointer bounds |
| 053 | Maximum Non Negative Product in a Matrix | Medium | Track min and max products |
| 054 | Frequency of the Most Frequent Element | Medium | Sort + sliding window |
| 055 | Maximum Number of Visible Points | Hard | Polar angles + sliding window |
| 056 | Closest Room | Hard | Offline sort by size + sorted set |
| 057 | Minimize Difference Target | Hard | Subset-sum DP across rows |
