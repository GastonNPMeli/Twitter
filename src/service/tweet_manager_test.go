package service_test

import (
	"github.com/GastonNPMeli/Twitter/src/domain"
	"github.com/GastonNPMeli/Twitter/src/service"
	"testing"
)

func isValidTweet( t *testing.T, tweet *domain.Tweet, id int, user string, text string) bool {
	return tweet != nil && user != "" && text != "" && id != -1 && tweet.Text == text && tweet.User == user
}

func TestPublishedTweetIsSaved(t *testing.T) {

	//Initialization
	var tweet *domain.Tweet
	user := "grupoesfera"
	text := "This is my first tweet"
	tweet = domain.NewTweet(user, text)

	//Operation
	service.PublishTweet(tweet)

	//Validation
	publishedTweet := service.GetTweets()[0]
	if publishedTweet.User != user &&
		publishedTweet.Text != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s",
			user, text, publishedTweet.User, publishedTweet.Text)
	}

	if publishedTweet.Date == nil {
		t.Error("Expected date can't be nil")
	}

}

func TestTweetWithoutUserIsNotPublished( t *testing.T) {
	//Initialization
	var tweet *domain.Tweet

	var user string
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	//Operation
	var err error
	_, err = service.PublishTweet(tweet)

	//Validation
	if err != nil && err.Error() != "user is required" {
		t.Error("Expected error is user is required")
	}
}

func TestTweetWithoutTextIsNotPublished( t *testing.T) {
	//Initialization
	var tweet *domain.Tweet

	user := "gponce"
	var text string

	tweet = domain.NewTweet(user, text)

	//Operation
	var err error
	_, err = service.PublishTweet(tweet)

	//Validation
	if err != nil && err.Error() != "text is required" {
		t.Error("Expected error is text is required")
	}
}

func TestTweetWhichExceeding140CharactersIsNotPublished( t *testing.T) {
	//Initialization
	var tweet *domain.Tweet

	user := "gponce"
	text := "1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890"

	tweet = domain.NewTweet(user, text)

	//Operation
	var err error
	_, err = service.PublishTweet(tweet)

	//Validation
	if err != nil && err.Error() != "len can't be more than 140 chars" {
		t.Error("Expected error is len can't be more than 140 chars")
	}
}

func TestCanPublishAndRetrieveMoreThanOneTweet( t *testing.T) {
	//Initialization
	service.InitializeService()
	var id1, id2 int
	user1 := "Juan"
	text1 := "Hola soy Juan"
	user2 := "Pedro"
	text2 := "Te faltó una coma, Juan."

	tweet := domain.NewTweet(user1, text1)
	secondTweet := domain.NewTweet(user2, text2)

	//Operation
	id1, _ = service.PublishTweet(tweet)
	id2, _ = service.PublishTweet(secondTweet)

	//Validation
	publishedTweets := service.GetTweets()
	if len(publishedTweets) != 2 {
		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
		return
	}

	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]


	if !isValidTweet(t, &firstPublishedTweet, id1, user1, text1) {
		return
	}

	if !isValidTweet(t, &secondPublishedTweet, id2, user2, text2) {
		return
	}

}

func TestCanRetrieveTweetById( t *testing.T) {
	//Initialization
	service.InitializeService()

	var tweet *domain.Tweet
	var id int

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	//Operation
	id, _ = service.PublishTweet(tweet)

	//Validation
	publishedTweet, _ := service.GetTweetById(id)

	isValidTweet(t, publishedTweet, id, user, text)
}