# Design / System Design

4 design problems where you implement a data structure with specific operation guarantees.

## Core patterns

| Pattern | When to use | Problems |
|---|---|---|
| **DLL of buckets** | O(1) insert/remove/min/max on count groups | 060 |
| **Sentinel-headed singly LL** | Avoid head/tail edge cases | 088 |
| **Binary lifting precompute** | k-th ancestor in O(log n) per query | 082 |
| **Per-user list + heap merge** | Timeline / feed across users | 085 |

## Common interview pitfalls

- **Confirm requirements** — "do you want O(1) get-max?", "is there a memory cap?", "is order of inserts guaranteed?". Designing without confirming costs you points.
- **Bucket invariants** — for All O`one, when a key's count changes, you may need to:
  (a) create a new bucket if no adjacent one has the target count, OR
  (b) move to the adjacent bucket, OR
  (c) delete the now-empty bucket. Handle all three.
- **Don't roll your own heap when stdlib works** — `container/heap` is fine; demonstrate fluency with it rather than reimplementing.
- **State the complexity per op** — "Inc is O(1), GetMaxKey is O(1)". Make it explicit; don't make the interviewer ask.
- **Twitter feed limits** — return only top 10; if you pop more, that's wasted work.
- **Binary lifting build cost** — `O(n log n)` time and space. Mention this when describing the data structure.

## Suggested practice order

1. **Day 1**: 088 (Design Linked List) — fundamental data structure, sentinel discipline
2. **Day 2**: 060 (All O`one) — DLL of buckets
3. **Day 3**: 085 (Design Twitter) — per-user list + heap merge
4. **Day 4**: 082 (Kth Ancestor) — binary lifting

## Problems in this folder

| # | Problem | Difficulty | Key idea |
|---|---|---|---|
| 060 | All O`one Data Structure | Hard | DLL of count-buckets |
| 082 | Kth Ancestor of a Tree Node | Hard | Binary lifting precompute |
| 085 | Design Twitter | Medium | Per-user list + heap merge |
| 088 | Design Linked List | Medium | Sentinel-headed singly LL |
