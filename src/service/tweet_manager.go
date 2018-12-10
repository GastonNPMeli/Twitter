package service

import (
	"errors"
	"github.com/GastonNPMeli/Twitter/src/domain"
	"os"
	"strings"
)

const path = "/Users/gponce/go/src/github.com/GastonNPMeli/Twitter/src/service/tweets.txt"

type TweetWriter interface {
	WriteTweet(domain.Tweet)
}

type FileTweetWriter struct {
	TweetList *os.File
}

type TweetManager struct{
	Tweets map[string][]domain.Tweet
	TweetCount *int
	TweetWriter
}

func NewTweetManager(tWritter TweetWriter) TweetManager {
	tweetCount := 0
	newTweetManager := TweetManager{
		make(map[string][]domain.Tweet),
		&tweetCount,
		tWritter,
	}
	return newTweetManager
}

func NewFileTweetWriter() *FileTweetWriter {
	fi, _ := os.OpenFile(path, os.O_RDWR, 0644)
	return &FileTweetWriter{fi}
}

func (tm TweetManager) PublishTweet(newTweet domain.Tweet) (tweetID int, err error) {

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

	tm.WriteTweet(newTweet)

	return newTweet.GetTweetId(), err
}

func (tm TweetManager) GetTweets() map[string][]domain.Tweet {
	return tm.Tweets
}

func (tm TweetManager) GetTweetById(id int) (tweet *domain.Tweet, err error) {

	for _, value := range tm.Tweets {
		for _, tweet := range value {
			if tweet.GetTweetId() == id {
				return &tweet, nil
			}
		}
	}

	return nil, errors.New("invalid tweetID")
}

func (tm TweetManager) CountTweetsByUser(user string) (tweetCount int, err error) {
	tweets, err := tm.GetTweetsByUser(user)
	return len(tweets), err
}

func (tm TweetManager) GetTweetsByUser(user string) (userTweets []domain.Tweet, err error) {
	if _, exists := tm.Tweets[user]; !exists {
		return nil, errors.New("invalid username")
	}

	return tm.Tweets[user], nil
}

func (tm TweetManager) SearchTweetsContaining (query string, ch chan domain.Tweet) {
	go func() {
		found := false

		for _, value := range tm.Tweets {
			for _, tweet := range value {
				if strings.Contains(tweet.GetText(), query) {
					ch <- tweet
					found = true
				}
			}
		}
		close(ch)

		if !found {
			ch = nil
		}
	}()
}

//FileTweetWriter

func (tw FileTweetWriter) WriteTweet(tweet domain.Tweet) {
	go func() {
		_, _ = tw.TweetList.WriteString(tweet.PrintableTweet()+"\n")
	}()
}
