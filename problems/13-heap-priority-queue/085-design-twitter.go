//go:build ignore

// Problem 085: Design Twitter (LeetCode #355)
// Difficulty: Medium
// Categories: Heap, Hash Table, Linked List, Design
//
// Description:
//   Implement a Twitter-like API:
//     PostTweet(userId, tweetId)
//     GetNewsFeed(userId) -> 10 most recent tweet ids posted by user or
//                            anyone they follow, newest first.
//     Follow(followerId, followeeId)
//     Unfollow(followerId, followeeId)
//
// Approach: Per-User Tweet List + Merge via Heap on Feed Read
//   Tweets stored as (timestamp, tweetId), prepended to user's list.
//   For GetNewsFeed: collect heads of (userId + followees) tweet lists into
//   a max-heap by timestamp. Pop up to 10, advance the popped list and
//   reinsert.
//
// Complexity:
//   Post:    O(1)
//   Follow:  O(1)
//   Feed:    O(F log F + 10 log F) where F = followee count
//   Space:   O(N) total tweets

package main

import (
	"container/heap"
	"fmt"
)

type tweet struct {
	id, ts int
}

type item struct {
	tweet
	userTweets []tweet
	idx        int
}

type maxHeap []*item

func (h maxHeap) Len() int            { return len(h) }
func (h maxHeap) Less(i, j int) bool  { return h[i].ts > h[j].ts }
func (h maxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x interface{}) { *h = append(*h, x.(*item)) }
func (h *maxHeap) Pop() interface{}   { old := *h; n := len(old); x := old[n-1]; *h = old[:n-1]; return x }

type Twitter struct {
	clock   int
	tweets  map[int][]tweet // user -> tweets newest-first
	follows map[int]map[int]bool
}

func ConstructorTwitter() Twitter {
	return Twitter{tweets: map[int][]tweet{}, follows: map[int]map[int]bool{}}
}

func (t *Twitter) PostTweet(userId, tweetId int) {
	t.clock++
	t.tweets[userId] = append([]tweet{{tweetId, t.clock}}, t.tweets[userId]...)
}

func (t *Twitter) GetNewsFeed(userId int) []int {
	h := &maxHeap{}
	add := func(uid int) {
		ts := t.tweets[uid]
		if len(ts) == 0 {
			return
		}
		heap.Push(h, &item{tweet: ts[0], userTweets: ts, idx: 0})
	}
	add(userId)
	for f := range t.follows[userId] {
		add(f)
	}
	var res []int
	for h.Len() > 0 && len(res) < 10 {
		top := heap.Pop(h).(*item)
		res = append(res, top.id)
		if top.idx+1 < len(top.userTweets) {
			top.idx++
			top.tweet = top.userTweets[top.idx]
			heap.Push(h, top)
		}
	}
	return res
}

func (t *Twitter) Follow(followerId, followeeId int) {
	if followerId == followeeId {
		return
	}
	if t.follows[followerId] == nil {
		t.follows[followerId] = map[int]bool{}
	}
	t.follows[followerId][followeeId] = true
}
func (t *Twitter) Unfollow(followerId, followeeId int) {
	if t.follows[followerId] != nil {
		delete(t.follows[followerId], followeeId)
	}
}

func main() {
	tw := ConstructorTwitter()
	tw.PostTweet(1, 5)
	fmt.Println(tw.GetNewsFeed(1))
	tw.Follow(1, 2)
	tw.PostTweet(2, 6)
	fmt.Println(tw.GetNewsFeed(1))
	tw.Unfollow(1, 2)
	fmt.Println(tw.GetNewsFeed(1))
}
