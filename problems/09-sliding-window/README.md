# Sliding Window

6 problems on fixed and variable-size sliding windows over arrays and strings.

## Core patterns

| Pattern | When to use | Problems |
|---|---|---|
| **Variable window + last-seen map** | "Longest substring with property X" | 004 |
| **Fixed window of `k` items** | Anagrams, fixed-length checks | 023 (window of words) |
| **Bucketing within window** | Numeric proximity (Δ-value, k-index) | 061 |
| **Window with rolling hash** | Substring matching of length L | 006 |
| **Sort + sliding window** | "Make k changes to maximize ..." | 054 |
| **Polar-angle window (wraparound)** | Visibility cone | 055 |

## Sliding window template

```go
// Variable window: shrink l while property is violated, track best.
l, best := 0, 0
state := newState()
for r := 0; r < len(arr); r++ {
    state.add(arr[r])
    for !state.valid() {
        state.remove(arr[l])
        l++
    }
    best = max(best, r-l+1)
}
```

For fixed window: add `arr[r]`, when `r-l+1 == k` record best and `state.remove(arr[l]); l++`.

## Common interview pitfalls

- **Map invalidation on slide** — when removing `arr[l]`, decrement the count; only `delete` from the map when count reaches 0, or use `count > 0` checks.
- **Substring-with-Concat overlap** — slide one WORD at a time (not one char). Try each of L starting offsets.
- **Wraparound for angle problems** — duplicate angles with `+2π` to handle the seam between 360° and 0°.
- **Off-by-one on window size** — window length is `r - l + 1`, not `r - l`.
- **`Frequency of the Most Frequent Element` cost formula** — `nums[r]*(r-l+1) - sum`; precompute the running `sum` to keep this O(1) per slide.
- **Hash collisions on rolling hash** — verify a hit by actual substring equality (or use double hashing).

## Suggested practice order

1. **Day 1**: 004 — variable window + last-seen map (classic warmup)
2. **Day 2**: 061 — bucketing within k-index window
3. **Day 3**: 054 — sort + sliding window (rare combo)
4. **Day 4**: 023 — window of words (fixed multiset)
5. **Day 5**: 006 — rolling hash + binary search outer
6. **Day 6**: 055 — angles and wraparound

## Problems in this folder

| # | Problem | Difficulty | Key idea |
|---|---|---|---|
| 004 | Longest Substring Without Repeating Characters | Medium | Variable window + last-seen map |
| 006 | Longest Duplicate Substring | Hard | BS length + Rabin-Karp window |
| 023 | Substring with Concatenation of All Words | Hard | Fixed window of words (multiset) |
| 054 | Frequency of the Most Frequent Element | Medium | Sort + variable window with cost |
| 055 | Maximum Number of Visible Points | Hard | Polar angles + wraparound window |
| 061 | Contains Duplicate III | Hard | Bucket within k-window |
