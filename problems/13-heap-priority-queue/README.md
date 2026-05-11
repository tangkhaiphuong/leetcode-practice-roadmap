# Heap / Priority Queue

4 problems where a heap is the right tool: top-K maintenance, day-sweep scheduling, stream merging, and strata processing.

## Core patterns

| Pattern | When to use | Problems |
|---|---|---|
| **Min-heap top-K** | "Maintain K smallest"; pop smallest, push new | (cross: 024 character-fix priority) |
| **Max-heap (negate or custom Less)** | "Pick the largest currently available" | 083 |
| **Day-sweep + min-heap on end** | Interval scheduling, attendance | 038 |
| **Heap-based merge of K streams** | Feed merge across users | 085 |
| **Lazy deletion** | When you can't easily remove arbitrary elements; mark stale, skip on pop | (cross: 038 expired events) |
| **Heap for repeated character resolution** | Password length/diversity fixes | 024 |

## Go heap recipe

`container/heap` requires a type implementing `sort.Interface` plus `Push(x interface{})` and `Pop() interface{}`. Boilerplate:

```go
type intHeap []int
func (h intHeap) Len() int            { return len(h) }
func (h intHeap) Less(i, j int) bool  { return h[i] < h[j] }   // min-heap
func (h intHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *intHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *intHeap) Pop() interface{}   { old := *h; n := len(old); x := old[n-1]; *h = old[:n-1]; return x }

// Use heap.Push(h, v) / heap.Pop(h) — NEVER call h.Push directly.
```

## Common interview pitfalls

- **`heap.Push(h, v)` vs `h.Push(v)`** — calling `h.Push` directly skips heapify and corrupts the heap. Always use the package-level functions.
- **Pointer receivers on Push/Pop** — required because they mutate the slice header.
- **Comparable types** — when storing structs, define a custom `Less`; don't try to interface-cast to `int`.
- **Don't forget to expire** — for day-sweep heaps (Max Events), pop expired entries before reading the top.
- **Lazy deletion overhead** — if your heap can fill with stale entries, your O(log n) operations become amortized but worst-case can degrade. Bound stale fraction.
- **Diminishing Balls** — solvable with strata math without an explicit heap, but heap-based simulation is fine for smaller cases and a useful baseline.

## Suggested practice order

1. **Day 1**: 038 (Max Events) — fundamental day-sweep + heap
2. **Day 2**: 085 (Twitter) — heap-merge of K sorted streams
3. **Day 3**: 024 (Strong Password) — heap-priority of which runs to break first
4. **Day 4**: 083 (Diminishing Balls) — heap mental model (or strata math)

## Problems in this folder

| # | Problem | Difficulty | Key idea |
|---|---|---|---|
| 024 | Strong Password Checker | Hard | Length-regime + priority of runs |
| 038 | Maximum Number of Events | Medium | Day-sweep + min-heap by end |
| 083 | Sell Diminishing-Valued Colored Balls | Medium | Strata greedy (heap baseline) |
| 085 | Design Twitter | Medium | Per-user list + max-heap merge |
