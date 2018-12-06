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
	service.InitializeService()
	var tweet *domain.Tweet
	user := "grupoesfera"
	text := "This is my first tweet"
	tweet = domain.NewTweet(user, text)

	//Operation
	service.PublishTweet(tweet)

	//Validation
	publishedTweet := service.GetTweets()["grupoesfera"][0]
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
	service.InitializeService()
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
	service.InitializeService()
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
	service.InitializeService()
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

	firstPublishedTweet := publishedTweets["Juan"][0]
	secondPublishedTweet := publishedTweets["Juan"][0]


	if !isValidTweet(t, firstPublishedTweet, id1, user1, text1) {
		return
	}

	if !isValidTweet(t, secondPublishedTweet, id2, user2, text2) {
		return
	}

}

func TestCanRetrieveTweetById( t *testing.T) {
	//Initialization
	service.InitializeService()

	var tweet *domain.Tweet
	var id int
	var err error

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	//Operation
	id, _ = service.PublishTweet(tweet)

	//Validation
	publishedTweet, err := service.GetTweetById(id)

	if err != nil {
		t.Errorf("Did not expect error, but Error: %s at tweetID %d", err, id)
	}

	if publishedTweet != tweet {
		t.Errorf("Expected tweets are not equal")
	}

	if !isValidTweet(t, publishedTweet, id, user, text) {
		return
	}

	publishedTweet, _ = service.GetTweetById(-1)

	if publishedTweet != nil {
		t.Errorf("Expected tweet should be nil")
		return
	}
}

func TestCanCountTheTweetsSentByAnUser( t *testing.T) {
	//Initialization
	service.InitializeService()
	var tweet, secondTweet, thirdTweet *domain.Tweet

	user := "grupoesfera"
	anotherUser := "nick"

	text := "this is my first tweet"
	secondText := "this is my second tweet"

	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)
	thirdTweet = domain.NewTweet(anotherUser, text)

	_, _ = service.PublishTweet(tweet)
	_, _ = service.PublishTweet(secondTweet)
	_, _ = service.PublishTweet(thirdTweet)

	//Operation
	count, _ := service.CountTweetsByUser(user)

	//Validation
	if count != 2 {
		t.Errorf("Expected count is 2 but was %d", count)
	}
}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {
	//Initialization
	service.InitializeService()

	var tweet, secondTweet, thirdTweet *domain.Tweet
	user := "grupoesfera"
	anotherUser := "nick"

	text := "This is my first tweet"
	secondText := "this is my second tweet"

	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)
	thirdTweet = domain.NewTweet(anotherUser, text)

	_, _ = service.PublishTweet(tweet)
	_, _ = service.PublishTweet(secondTweet)
	_, _ = service.PublishTweet(thirdTweet)

	//operation
	tweets, _ := service.GetTweetsByUser(user)

	//Validation
	if len(tweets) != 2 {
		t.Errorf("Expected count is 2 but was %d", len(tweets))
	}

	firstPublishedTweet := tweets[0]
	secondPublishedTweet := tweets[1]

	if tweet != firstPublishedTweet || secondTweet != secondPublishedTweet {
		t.Errorf("The tweets don't match")
	}
}

func TestRetrieveUserTweetsShouldErrorWhenNoUserFound(t *testing.T) {
	//Initialization
	service.InitializeService()

	if _, err := service.GetTweetsByUser("prueba"); err == nil {
		t.Errorf("Expected username error")
	}

}