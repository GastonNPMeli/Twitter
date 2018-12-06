package service

import (
	"errors"
	"github.com/GastonNPMeli/Twitter/src/domain"
)

var Tweets []domain.Tweet

func InitializeService() {
	Tweets = []domain.Tweet {}
}

func PublishTweet(newTweet *domain.Tweet) (tweetID int, err error) {

	if newTweet.User == "" {
		return -1, errors.New("user is required")
	}

	if newTweet.Text == "" {
		return -1, errors.New("text is required")
	}

	if len(newTweet.Text) > 140 {
		return -1, errors.New("len can't be more than 140 chars")
	}

	Tweets = append(Tweets, *newTweet)

	return len(Tweets), err
}

func GetTweets() []domain.Tweet {
	return Tweets
}

func GetTweetById(id int) (tweet *domain.Tweet, err error) {

	if id <= len(Tweets) && id > 0 {
		return &Tweets[id - 1], nil
	}

	return nil, errors.New("Invalid tweetID")
}



