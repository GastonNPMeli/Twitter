package service

import (
	"github.com/GastonNPMeli/Twitter/src/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type Router struct {
	Engine *gin.Engine
	Tw TweetManager
}

type PublishTextTweetCommand struct {
	User string `json:"user"`
	Text string `json:"text"`
}

type PublishImageTweetCommand struct {
	User string `json:"user"`
	Text string `json:"text"`
	Image string `json:"image"`
}

type PublishQuoteTweetCommand struct {
	User string `json:"user"`
	Text string `json:"text"`
	QuotedTweetId string `json:"quoteid"`
}

var imagePath = "/Users/gponce/go/src/github.com/GastonNPMeli/Twitter/src/service/david.jpg"

func NewGinRouter(tweetManager TweetManager) Router {
	return Router{gin.Default(), tweetManager}
}

func (r *Router) StartGinRouter() {
	go func() {
		router := gin.Default()
		router.GET("tweets", r.ListTweets)
		router.GET("tweets/@:user", r.ListUserTweets)
		router.GET("tweets/@:user/count", r.GetUserTweetCount)
		router.GET("tweets/id=:tweetId", r.GetTweetById)
		router.GET("tweets/search=:query", r.ListTweetsByQuery)


		router.POST("tweets/newTextTweet/", r.PublishTextTweet)
		router.POST("tweets/newImageTweet/", r.PublishImageTweet)
		router.POST("tweets/newQuoteTweet/", r.PublishQuoteTweet)
		router.Run()
	}()
}

func (r *Router) ListTweets(c * gin.Context) {
	for _, tweets := range r.Tw.GetTweets() {
		for _, tweet := range tweets {
			c.String(http.StatusOK, tweet.PrintableTweet() + "\n")
		}
	}
}

func (r *Router) ListTweetsByQuery(c * gin.Context) {
	ch := make(chan domain.Tweet, 100)
	r.Tw.SearchTweetsContaining(c.Param("query"), ch)

	time.Sleep(100 + time.Millisecond)

	for tweet := range ch {
		c.String(http.StatusOK, tweet.PrintableTweet() + "\n")
	}
}

func (r *Router) GetTweetById(c * gin.Context) {
	tweetId, _ := strconv.Atoi(c.Param("tweetId"))
	tweet, _ := r.Tw.GetTweetById(tweetId)

	if tweet == nil {
		c.String(http.StatusOK, "No existe ningún Tweet con el ID proporcionado!\n")
		return
	}

	c.String(http.StatusOK, (*tweet).PrintableTweet() + "\n")
}

func (r *Router) ListUserTweets(c * gin.Context) {
	tweets, _ := r.Tw.GetTweetsByUser(c.Param("user"))

	for _, tweet := range tweets {
		c.String(http.StatusOK, tweet.PrintableTweet() + "\n")
	}
}

func (r *Router) GetUserTweetCount(c * gin.Context) {
	count, _ := r.Tw.CountTweetsByUser(c.Param("user"))
	countStr := strconv.Itoa(count)
	c.String(http.StatusOK, "@"+ c.Param("user") + " tiene " + countStr + " Tweets.\n")
}

func (r *Router) PublishTextTweet(c * gin.Context) {
	var publishTextTweetCmd PublishTextTweetCommand
	c.BindJSON(&publishTextTweetCmd)

	newTweet := domain.NewTextTweet(publishTextTweetCmd.User, publishTextTweetCmd.Text)
	r.Tw.PublishTweet(newTweet)
	c.JSON(http.StatusOK, "Posted " + newTweet.PrintableTweet())
}

func (r *Router) PublishImageTweet(c * gin.Context) {
	var publishImageTweetCmd PublishImageTweetCommand
	c.BindJSON(&publishImageTweetCmd)

	newTweet := domain.NewImageTweet(publishImageTweetCmd.User, publishImageTweetCmd.Text, publishImageTweetCmd.Image)
	r.Tw.PublishTweet(newTweet)
	c.JSON(http.StatusOK, "Posted " + newTweet.PrintableTweet())
}

func (r *Router) PublishQuoteTweet(c * gin.Context) {
	var publishQuoteTweetCmd PublishQuoteTweetCommand
	c.BindJSON(&publishQuoteTweetCmd)

	quotedTweetId, _ := strconv.Atoi(publishQuoteTweetCmd.QuotedTweetId)
	quotedTweet, _ := r.Tw.GetTweetById(quotedTweetId)

	newTweet := domain.NewQuoteTweet(publishQuoteTweetCmd.User, publishQuoteTweetCmd.Text, *quotedTweet)
	r.Tw.PublishTweet(newTweet)
	c.JSON(http.StatusOK, "Posted " + newTweet.PrintableTweet())
}


