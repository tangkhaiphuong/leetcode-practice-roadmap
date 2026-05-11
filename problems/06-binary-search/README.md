# Binary Search

14 problems including value search, partition search, "search the answer", and fast exponentiation.

## Core patterns

| Pattern | When to use | Examples |
|---|---|---|
| **Search a value in sorted array** | Direct lookup (allow duplicates) | 073 |
| **Search the answer** (parametric search) | Monotonic predicate → binary-search the threshold | 006, 044, 054, 068, 075, 076 |
| **Binary search on prefix-sum / index space** | First/last position with property | 044 |
| **Fast exponentiation** | Pow, modular pow, matrix pow | 077 |
| **Search the partition** | Median of two arrays — partition both, check cross | 032 |
| **Right-to-left monotonic stack + BS** | 132 pattern (find ranges) | 037 |
| **Merge sort + cross-pair count** | Counting inversions / reverse pairs | 043 |
| **Sorted set + neighbor lookup** | Offline queries with insertion | 056 |
| **SQL window function** (conceptual BS) | Nth highest | 090 |

## Two reliable templates

```go
// Template 1: closed range [lo, hi]; loop while lo <= hi.
for lo <= hi {
    mid := (lo + hi) / 2
    if found { return mid }
    if condition { hi = mid - 1 } else { lo = mid + 1 }
}

// Template 2: half-open [lo, hi); loop while lo < hi. Good for "first index with X".
for lo < hi {
    mid := (lo + hi) / 2
    if pred(mid) { hi = mid } else { lo = mid + 1 }
}
// returns lo (== hi)
```

## Common interview pitfalls

- **Overflow in `(lo + hi) / 2`** — use `lo + (hi-lo)/2` as a habit.
- **Predicate monotonicity** — confirm before reaching for binary search. If `pred(k)` holds, must `pred(k+1)` also hold (or vice-versa)?
- **Boundary off-by-one** — write tiny test cases (n=1, n=2) and trace by hand.
- **Pow with negative exponent** — `n = INT_MIN` can't be negated as int32. Convert to int64 first.
- **Duplicates in rotated sorted array** — when `nums[lo] == nums[mid] == nums[hi]`, you can't tell which half is sorted; shrink both ends by 1.
- **`Median of Two Sorted Arrays`** — partition on the SHORTER array so `j = half - i` stays in range.

## Suggested practice order

1. **Day 1**: 077 (Pow), 073 — exponentiation + classic rotated BS
2. **Day 2**: 032 — partition search (classic Hard)
3. **Day 3**: 074, 044 — index-space binary search
4. **Day 4**: 037, 043 — combined with stack / merge sort
5. **Day 5**: 068 — DP dimension flip + BS bound
6. **Day 6**: 054, 056, 075 — "search the answer" with offline queries
7. **Day 7**: 006, 076 — binary search + rolling hash
8. **Day 8**: 090 — SQL Nth Highest (window functions / LIMIT OFFSET)

## Problems in this folder

| # | Problem | Difficulty | Key idea |
|---|---|---|---|
| 006 | Longest Duplicate Substring | Hard | BS length + Rabin-Karp |
| 032 | Median of Two Sorted Arrays | Hard | Partition both arrays |
| 037 | 132 Pattern | Medium | Right-to-left monotonic stack |
| 043 | Reverse Pairs | Hard | Merge sort with count |
| 044 | Shortest Subarray with Sum at Least K | Hard | Monotonic deque on prefix sums |
| 054 | Frequency of the Most Frequent Element | Medium | Sort + sliding window |
| 056 | Closest Room | Hard | Offline sort by size + sorted set |
| 068 | Super Egg Drop | Hard | dp[m][k] = max floors with m moves |
| 073 | Search in Rotated Sorted Array II | Medium | Modified BS w/ duplicate skip |
| 074 | Nth Digit | Medium | Bracket by digit-length bands |
| 075 | Minimum Space Wasted From Packaging | Hard | Sort + prefix sum + BS per box |
| 076 | Longest Common Subpath | Hard | Binary search length + rolling hash |
| 077 | Pow(x, n) | Medium | Iterative fast exponentiation |
| 090 | Nth Highest Salary (SQL) | Medium | LIMIT/OFFSET or DENSE_RANK |
