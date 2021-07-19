package main

import (
	"container/heap"
)

type TweetHeap []*Tweet

func (f TweetHeap) Len() int           { return len(f) }
func (f TweetHeap) Less(i, j int) bool { return f[i].PostTime > f[j].PostTime }
func (f TweetHeap) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
func (f *TweetHeap) Push(x interface{}) {
	e := x.(*Tweet)
	*f = append(*f, e)
	heap.Fix(f, f.Len()-1)
}

func (f *TweetHeap) Pop() interface{} {
	e := (*f)[0]
	*f = (*f)[1:len(*f)]
	heap.Fix(f, 0)
	return e
}

type User struct {
	ID        int
	Followers map[int]struct{} // 关注人列表
	Tweets    *Tweet
}

func NewUser(userId int) *User {
	return &User{ID: userId, Followers: map[int]struct{}{userId: {}}}
}

type Tweet struct {
	ID       int
	PostTime int
	Next     *Tweet
}

type Twitter struct {
	users map[int]*User
	clock int
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{users: make(map[int]*User)}
}

/** Compose a new tweet. */
func (t *Twitter) PostTweet(userId int, tweetId int) {
	user, ok := t.users[userId]
	if !ok {
		user = NewUser(userId)
		t.users[userId] = user
	}
	t.clock++
	tweet := &Tweet{
		ID:       tweetId,
		PostTime: t.clock,
		Next:     user.Tweets,
	}
	user.Tweets = tweet
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user Followers or by the user herself. Tweets must be ordered from most recent to least recent. */
func (t *Twitter) GetNewsFeed(userId int) []int {
	queue := TweetHeap{}
	user, ok := t.users[userId]
	if !ok {
		return nil
	}
	for followerID := range user.Followers {
		if followerUser, ok := t.users[followerID]; ok {
			if tweet := followerUser.Tweets; tweet != nil {
				queue.Push(tweet)
			}
		}
	}

	var res []int
	// 取前10条
	for i := 0; i < 10 && queue.Len() != 0; i++ {
		tweet := queue.Pop().(*Tweet)
		res = append(res, tweet.ID)
		if tweet.Next != nil {
			queue.Push(tweet.Next)
		}
	}
	return res
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (t *Twitter) Follow(followerId int, followeeId int) {
	user, ok := t.users[followerId]
	if !ok {
		user = NewUser(followerId)
		t.users[followerId] = user
	}
	user.Followers[followeeId] = struct{}{}
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (t *Twitter) Unfollow(followerId int, followeeId int) {
	if user, ok := t.users[followerId]; ok {
		if _, ok := user.Followers[followeeId]; ok && followeeId != user.ID {
			delete(user.Followers, followeeId)
		}
	}
}
