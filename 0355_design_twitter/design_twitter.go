package leetcode0355

import (
	"sort"
	"time"
)

type feed struct {
	id   int
	time int64
}

func sortByTime(feeds []feed) {
	sort.Slice(feeds, func(i, j int) bool {
		return feeds[i].time > feeds[j].time
	})
}

type Twitter struct {
	users  map[int][]feed
	follow map[int]map[int]struct{}
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{
		users:  make(map[int][]feed),
		follow: make(map[int]map[int]struct{}),
	}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.users[userId] = append(
		this.users[userId],
		feed{
			id:   tweetId,
			time: time.Now().UnixNano(),
		},
	)
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	tmp := make([]feed, len(this.users[userId]))
	copy(tmp, this.users[userId])
	for id := range this.follow[userId] {
		tmp = append(tmp, this.users[id]...)
	}
	sortByTime(tmp)

	n := 10
	if len(tmp) < n {
		n = len(tmp)
	}
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = tmp[i].id
	}
	return res
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}
	if _, ok := this.follow[followerId]; !ok {
		this.follow[followerId] = map[int]struct{}{
			followeeId: struct{}{},
		}
	} else {
		this.follow[followerId][followeeId] = struct{}{}
	}

}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	delete(this.follow[followerId], followeeId)
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
