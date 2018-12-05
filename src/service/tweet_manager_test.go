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
