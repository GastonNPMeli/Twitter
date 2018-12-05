package domain

import "time"

type Tweet struct {
	User string
	Text string
	Date time.Time
}

func NewTweet(user string, text string) *Tweet {
	return &Tweet{user, text, time.Now()}
}
