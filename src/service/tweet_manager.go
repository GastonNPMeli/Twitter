package service

import (
	"errors"
	"github.com/GastonNPMeli/Twitter/src/domain"
)

var Tweets map[string][]*domain.Tweet
var tweetCount int

func InitializeService() {
	Tweets = make(map[string][]*domain.Tweet)
	tweetCount = 1
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

	if Tweets[newTweet.User] == nil {
		Tweets[newTweet.User] = make([]*domain.Tweet,0)
	}

	newTweet.TweetId = tweetCount
	tweetCount++
	Tweets[newTweet.User] = append(Tweets[newTweet.User], newTweet)

	return len(Tweets), err
}

func GetTweets() map[string][]*domain.Tweet {
	return Tweets
}

func GetTweetById(id int) (tweet *domain.Tweet, err error) {

	for _, value := range Tweets {
		for _, tweet := range value {
			if tweet.TweetId == id {
				return tweet, nil
			}
		}
	}

	return nil, errors.New("invalid tweetID")
}

func CountTweetsByUser(user string) (tweetCount int, err error) {
	tweets, err := GetTweetsByUser(user)
	return len(tweets), err
}

func GetTweetsByUser(user string) (userTweets []*domain.Tweet, err error) {
	if _, exists := Tweets[user]; !exists {
		return nil, errors.New("invalid username")
	}

	return Tweets[user], nil
}



