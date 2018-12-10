package service_test

import (
	"github.com/GastonNPMeli/Twitter/src/domain"
	"github.com/GastonNPMeli/Twitter/src/service"
	"io"
	"os"
	"strings"
	"testing"
)

func isValidTweet( t *testing.T, tweet domain.Tweet, id int, user string, text string) bool {
	return tweet != nil && user != "" && text != "" && id != -1 && tweet.GetText() == text && tweet.GetUser() == user
}

func TestPublishedTweetIsSaved(t *testing.T) {

	//Initialization
	tweetWriter := service.NewFileTweetWriter()
	tweetManager := service.NewTweetManager(tweetWriter)
	var tweet *domain.TextTweet
	user := "grupoesfera"
	text := "This is my first tweet"
	tweet = domain.NewTextTweet(user, text)

	//Operation
	tweetManager.PublishTweet(tweet)

	//Validation
	publishedTweet := tweetManager.GetTweets()["grupoesfera"][0]
	if publishedTweet.GetUser() != user &&
		publishedTweet.GetText() != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s",
			user, text, publishedTweet.GetUser(), publishedTweet.GetText())
	}

	if publishedTweet.GetDate() == nil {
		t.Error("Expected date can't be nil")
	}

}

func TestTweetWithoutUserIsNotPublished( t *testing.T) {
	//Initialization
	tweetWriter := service.NewFileTweetWriter()
	tweetManager := service.NewTweetManager(tweetWriter)
	var tweet *domain.TextTweet

	var user string
	text := "This is my first tweet"

	tweet = domain.NewTextTweet(user, text)

	//Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	//Validation
	if err != nil && err.Error() != "user is required" {
		t.Error("Expected error is user is required")
	}
}

func TestTweetWithoutTextIsNotPublished( t *testing.T) {
	//Initialization
	tweetWriter := service.NewFileTweetWriter()
	tweetManager := service.NewTweetManager(tweetWriter)
	var tweet *domain.TextTweet

	user := "gponce"
	var text string

	tweet = domain.NewTextTweet(user, text)

	//Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	//Validation
	if err != nil && err.Error() != "text is required" {
		t.Error("Expected error is text is required")
	}
}

func TestTweetWhichExceeding140CharactersIsNotPublished( t *testing.T) {
	//Initialization
	tweetWriter := service.NewFileTweetWriter()
	tweetManager := service.NewTweetManager(tweetWriter)
	var tweet *domain.TextTweet

	user := "gponce"
	text := "1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890"

	tweet = domain.NewTextTweet(user, text)

	//Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	//Validation
	if err != nil && err.Error() != "len can't be more than 140 chars" {
		t.Error("Expected error is len can't be more than 140 chars")
	}
}

func TestCanPublishAndRetrieveMoreThanOneTweet( t *testing.T) {
	//Initialization
	tweetWriter := service.NewFileTweetWriter()
	tweetManager := service.NewTweetManager(tweetWriter)
	var id1, id2 int
	user1 := "Juan"
	text1 := "Hola soy Juan"
	user2 := "Pedro"
	text2 := "Te faltó una coma, Juan."

	tweet := domain.NewTextTweet(user1, text1)
	secondTweet := domain.NewTextTweet(user2, text2)

	//Operation
	id1, _ = tweetManager.PublishTweet(tweet)
	id2, _ = tweetManager.PublishTweet(secondTweet)

	//Validation
	publishedTweets := tweetManager.GetTweets()
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
	tweetWriter := service.NewFileTweetWriter()
	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet *domain.TextTweet
	var id int
	var err error

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTextTweet(user, text)

	//Operation
	id, _ = tweetManager.PublishTweet(tweet)

	//Validation
	publishedTweet, err := tweetManager.GetTweetById(id)

	if err != nil {
		t.Errorf("Did not expect error, but Error: %s at tweetID %d", err, id)
	}

	if *publishedTweet != tweet {
		t.Errorf("Expected tweets are not equal")
	}

	if !isValidTweet(t, *publishedTweet, id, user, text) {
		return
	}

	publishedTweet, _ = tweetManager.GetTweetById(-1)

	if publishedTweet != nil {
		t.Errorf("Expected tweet should be nil")
		return
	}
}

func TestCanCountTheTweetsSentByAnUser( t *testing.T) {
	//Initialization
	tweetWriter := service.NewFileTweetWriter()
	tweetManager := service.NewTweetManager(tweetWriter)
	var tweet, secondTweet, thirdTweet *domain.TextTweet

	user := "grupoesfera"
	anotherUser := "nick"

	text := "this is my first tweet"
	secondText := "this is my second tweet"

	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)
	thirdTweet = domain.NewTextTweet(anotherUser, text)

	_, _ = tweetManager.PublishTweet(tweet)
	_, _ = tweetManager.PublishTweet(secondTweet)
	_, _ = tweetManager.PublishTweet(thirdTweet)

	//Operation
	count, _ := tweetManager.CountTweetsByUser(user)

	//Validation
	if count != 2 {
		t.Errorf("Expected count is 2 but was %d", count)
	}
}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {
	//Initialization
	tweetWriter := service.NewFileTweetWriter()
	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet, secondTweet, thirdTweet *domain.TextTweet
	user := "grupoesfera"
	anotherUser := "nick"

	text := "This is my first tweet"
	secondText := "this is my second tweet"

	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)
	thirdTweet = domain.NewTextTweet(anotherUser, text)

	_, _ = tweetManager.PublishTweet(tweet)
	_, _ = tweetManager.PublishTweet(secondTweet)
	_, _ = tweetManager.PublishTweet(thirdTweet)

	//operation
	tweets, _ := tweetManager.GetTweetsByUser(user)

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
	tweetWriter := service.NewFileTweetWriter()
	tweetManager := service.NewTweetManager(tweetWriter)

	if _, err := tweetManager.GetTweetsByUser("prueba"); err == nil {
		t.Errorf("Expected username error")
	}
}

//TODO: Hacer tests de Quoted con referencias y Quoted con imagen

func TestPublishedTweetIsSavedToExternalResource(t *testing.T) {

	// Initialization
	var tweetWriter service.TweetWriter
	tweetWriter = service.NewFileTweetWriter()
	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet domain.Tweet // Fill the tweet with data
	tweet = domain.NewTextTweet("Juan", "Hola")
	// Operation
	_, _ = tweetManager.PublishTweet(tweet)

	//Validation
	var fi, _ = os.OpenFile("tweets.txt", os.O_RDWR, 0644)

	var line = make([]byte, 1024)
	lines := 0
	for {
		_, err := fi.Read(line)

		if err == io.EOF {
			break
		}

		lines++
	}

	if lines != 1 {
		t.Errorf("Expected 1 line but got %d", lines)
	}
}

func TestCanSearchForTweetContainingText(t *testing.T) {
	// Initialization
	var tweetWriter service.TweetWriter
	tweetWriter = service.NewFileTweetWriter()
	tweetManager := service.NewTweetManager(tweetWriter)

	tweet := domain.NewTextTweet("Juan", "Esto es muy raro")
	_, _ = tweetManager.PublishTweet(tweet)

	// Operation
	searchResult := make(chan domain.Tweet)
	query := "raro"
	tweetManager.SearchTweetsContaining(query, searchResult)

	// Validation
	foundTweet := <-searchResult

	if foundTweet == nil {
		t.Errorf("Expected to find tweet")
	}

	if !strings.Contains(foundTweet.GetText(), query) {
		t.Errorf("Expected to find %s, found %s", query, foundTweet.GetText())
	}
}

func TestChannelIsNULLIfThereAreNoMatchingTweets(t *testing.T) {
	// Initialization
	var tweetWriter service.TweetWriter
	tweetWriter = service.NewFileTweetWriter()
	tweetManager := service.NewTweetManager(tweetWriter)

	// Operation
	searchResult := make(chan domain.Tweet)
	query := ""
	tweetManager.SearchTweetsContaining(query, searchResult)

	go func() {
		if searchResult != nil {
			t.Errorf("Expected to find no tweets")
		}
	}()
}