package domain

import "time"

type Tweet struct {
	TweetId int
	User string
	Text string
	Date *time.Time
}

func NewTweet(user string, text string) *Tweet {
	date := time.Now()
	return &Tweet{0, user, text,  &date}
}
