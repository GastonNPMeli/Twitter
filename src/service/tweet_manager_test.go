package service_test

import (
	"github.com/GastonNPMeli/Twitter/src/domain"
	"github.com/GastonNPMeli/Twitter/src/service"
	"testing"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	//Initialization
	var tweet *domain.Tweet
	user := "grupoesfera"
	text := "This is my first tweet"
	tweet = domain.NewTweet(user, text)

	//Operation
	service.PublishTweet(tweet)

	//Validation
	publishedTweet := service.GetTweet()
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
	err = service.PublishTweet(tweet)

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
	err = service.PublishTweet(tweet)

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
	err = service.PublishTweet(tweet)

	//Validation
	if err != nil && err.Error() != "len can't be more than 140 chars" {
		t.Error("Expected error is len can't be more than 140 chars")
	}
}
