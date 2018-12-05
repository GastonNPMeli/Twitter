package service

import "github.com/GastonNPMeli/Twitter/src/domain"

var Tweet domain.Tweet

func PublishTweet(newTweet *domain.Tweet) {
	Tweet = *newTweet
}

func GetTweet() domain.Tweet {
	return Tweet
}
