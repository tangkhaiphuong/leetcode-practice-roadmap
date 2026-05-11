# Sorting & Order

8 problems where sorting (or sort-adjacent techniques like quickselect / merge sort) is the key step.

## Core patterns

| Pattern | When to use | Problems |
|---|---|---|
| **Sort + two pointers** | Pair / triple sum, partitioning | 033 |
| **Sort + sliding window** | "Make k changes to maximize ..." | 054 |
| **Sort + sweep / day-pointer** | Interval scheduling, attendance | 038 |
| **Sort + greedy take from largest** | Diminishing values, top-K | 016 (digits removed) |
| **Quickselect / sort-then-rewire** | Wiggle sort, median placements | 051 |
| **Merge sort counting** | Inversions, cross-pair counts | 043 |
| **Offline queries — sort by both inputs** | Closest room, packaging | 056 |
| **Sort + group / pairwise** | Detect duplicates across constraints | 062 |

## Common interview pitfalls

- **`sort.Slice` is unstable** — use `sort.SliceStable` when tie order matters (e.g., when secondary key from input order matters).
- **Quickselect partition correctness** — pick the right pivot strategy (median-of-three or random) to avoid `O(n²)` on adversarial inputs.
- **Merge sort allocation** — pre-allocate the temporary buffer once and reuse, not re-allocate per recursion level.
- **Sorting a custom struct** — Go's `sort.Slice` takes `Less(i, j) bool`. Avoid reading the elements multiple times; capture them once for readability.
- **Sort vs partial sort** — for top-K, `O(n log k)` heap or `O(n)` quickselect beats `O(n log n)` full sort.

## Suggested practice order

1. **Day 1**: 033 — sort + two pointers (classic dedup)
2. **Day 2**: 016, 062 — sort + greedy / sort + group
3. **Day 3**: 038, 056 — interval and offline-query problems
4. **Day 4**: 054 — sort + sliding window (rare combo)
5. **Day 5**: 043 — merge sort counting
6. **Day 6**: 051 — sort-then-index-rewire trick

## Problems in this folder

| # | Problem | Difficulty | Key idea |
|---|---|---|---|
| 016 | Remove K Digits | Medium | Monotonic increasing stack |
| 033 | 3Sum | Medium | Sort + two pointers + dedup |
| 038 | Maximum Number of Events | Medium | Sort by start + min-heap by end |
| 043 | Reverse Pairs | Hard | Merge sort with cross-pair count |
| 051 | Wiggle Sort II | Medium | Sort + reversed-half index rewiring |
| 054 | Frequency of the Most Frequent Element | Medium | Sort + sliding window |
| 056 | Closest Room | Hard | Offline sort + sorted set neighbor BS |
| 062 | Invalid Transactions | Medium | Group by name + pairwise check |
