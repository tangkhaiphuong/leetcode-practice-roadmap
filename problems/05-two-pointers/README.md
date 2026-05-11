# Two Pointers Technique

8 problems on opposing-pointer and same-direction-pointer patterns.

## Core patterns

| Pattern | When to use | Problems |
|---|---|---|
| **Opposing pointers** (l, r from each end) | Sorted-array pair sum, palindrome, Sum of Squares | 033, 072 |
| **Same-direction (slow/fast)** | Skip duplicates, in-place compaction, edit-distance scan | 013, 021, 042 |
| **Two-string parallel walk** | Compare with optional skips | 003 |
| **Next-permutation (pivot + reverse)** | Lexicographic next | 036 |
| **Two pointers around prefix sum** | Range bounds for split / window | 052 |

## Common interview pitfalls

- **Overflow in `l*l + r*r`** — for inputs up to 2^31 - 1, use `uint64` or `int64`.
- **`l < r` vs `l <= r`** — if a single element can satisfy (e.g., `c = 0` so `l == r == 0` works), use `<=`.
- **Skipping duplicates** — do it AFTER recording a hit, advancing BOTH pointers past the run AND the outer index in nested cases.
- **Initial `r`** — for sqrt-bounded search, `r = floor(sqrt(c))` (not `c`); cuts iterations dramatically.
- **Next Permutation** — find pivot from the right, NOT from the left; pivot is `nums[i] < nums[i+1]` (strictly less, not <=).
- **Edit-distance "one off"** — when lengths differ, the shorter advances by 1, the longer by 2 at the mismatch. Match the right offsets.

## Suggested practice order

1. **Day 1**: 033 (3Sum), 072 (Sum of Squares) — opposing pointers fluency
2. **Day 2**: 003, 013, 021 — string walks
3. **Day 3**: 036 (Next Permutation) — pivot + reverse pattern
4. **Day 4**: 042, 052 — pointer bounds on prefix sums

## Problems in this folder

| # | Problem | Difficulty | Key idea |
|---|---|---|---|
| 003 | Compare Version Numbers | Medium | Two-pointer parse without splitting |
| 013 | Reverse Words in a String | Medium | `strings.Fields` + reverse |
| 021 | One Edit Distance | Medium | Length-diff branch + scan |
| 033 | 3Sum | Medium | Sort + two pointers + dedup |
| 036 | Next Permutation | Medium | Pivot, swap, reverse tail |
| 042 | Shortest Unsorted Continuous Subarray | Medium | Single-pass min/max from both ends |
| 052 | Ways to Split Array Into Three Subarrays | Medium | Prefix sums + two-pointer bounds |
| 072 | Sum of Square Numbers | Medium | Opposing pointers on [0, sqrt(c)] |
