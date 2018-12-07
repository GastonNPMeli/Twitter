package service

import (
	"errors"
	"github.com/GastonNPMeli/Twitter/src/domain"
)

type tweetManager struct{
	Tweets map[string][]domain.Tweet
	TweetCount *int
}

func NewTweetManager() tweetManager {
	tweetCount := 0
	newTweetManager := tweetManager{
		make(map[string][]domain.Tweet),
		&tweetCount,
	}
	return newTweetManager
}

func (tm tweetManager) PublishTweet(newTweet domain.Tweet) (tweetID int, err error) {

	if newTweet.GetUser() == "" {
		return -1, errors.New("user is required")
	}

	if newTweet.GetText() == "" {
		return -1, errors.New("text is required")
	}

	if len(newTweet.GetText()) > 140 {
		return -1, errors.New("len can't be more than 140 chars")
	}

	if tm.Tweets[newTweet.GetUser()] == nil {
		tm.Tweets[newTweet.GetUser()] = make([]domain.Tweet,0)
	}

	*tm.TweetCount++
	newTweet.SetTweetId(*tm.TweetCount)
	tm.Tweets[newTweet.GetUser()] = append(tm.Tweets[newTweet.GetUser()], newTweet)

	return newTweet.GetTweetId(), err
}

func (tm tweetManager) GetTweets() map[string][]domain.Tweet {
	return tm.Tweets
}

func (tm tweetManager) GetTweetById(id int) (tweet *domain.Tweet, err error) {

	for _, value := range tm.Tweets {
		for _, tweet := range value {
			if tweet.GetTweetId() == id {
				return &tweet, nil
			}
		}
	}

	return nil, errors.New("invalid tweetID")
}

func (tm tweetManager) CountTweetsByUser(user string) (tweetCount int, err error) {
	tweets, err := tm.GetTweetsByUser(user)
	return len(tweets), err
}

func (tm tweetManager) GetTweetsByUser(user string) (userTweets []domain.Tweet, err error) {
	if _, exists := tm.Tweets[user]; !exists {
		return nil, errors.New("invalid username")
	}

	return tm.Tweets[user], nil
}



