package service

import (
	"errors"
	"github.com/GastonNPMeli/Twitter/src/domain"
)

var Tweet domain.Tweet

func PublishTweet(newTweet *domain.Tweet) (err error) {

	if newTweet.User == "" {
		return errors.New("user is required")
	}

	if newTweet.Text == "" {
		return errors.New("text is required")
	}

	if len(newTweet.Text) > 140 {
		return errors.New("len can't be more than 140 chars")
	}

	Tweet = *newTweet

	return err
}

func GetTweet() domain.Tweet {
	return Tweet
}
