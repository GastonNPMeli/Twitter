package main

import (
	"github.com/GastonNPMeli/Twitter/src/domain"
	"github.com/GastonNPMeli/Twitter/src/service"
	"github.com/abiosoft/ishell"
	"strconv"
)

func main() {


	tweetWriter := service.NewFileTweetWriter()
	tweetManager := service.NewTweetManager(tweetWriter)

	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTextTweet",
		Help: "Publishes a text tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write your user: ")

			user := c.ReadLine()

			c.Print("Write your tweet: ")

			text := c.ReadLine()

			var tweet = domain.NewTextTweet(user, text)

			tweetManager.PublishTweet(tweet)

			c.Print("Tweet sent\n")

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "publishImageTweet",
		Help: "Publishes an image tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write your user: ")

			user := c.ReadLine()

			c.Print("Write your tweet: ")

			text := c.ReadLine()

			c.Print("Write your image url: ")

			image := c.ReadLine()

			var tweet = domain.NewImageTweet(user, text, image)

			tweetManager.PublishTweet(tweet)

			c.Print("Tweet sent\n")

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "publishQuotedTweet",
		Help: "Publishes a quoted tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write your user: ")

			user := c.ReadLine()

			c.Print("Write your tweet: ")

			text := c.ReadLine()

			c.Print("Write the TweetId of the tweet you want to quote: ")

			id, _ := strconv.Atoi(c.ReadLine())

			quotedTweet, _ := tweetManager.GetTweetById(id)

			var tweet = domain.NewQuoteTweet(user, text, *quotedTweet)

			id, _ = tweetManager.PublishTweet(tweet)

			c.Print("Tweet sent\n")

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "Shows a tweet, given its tweetID",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write your tweetID: ")

			id, _ := strconv.Atoi(c.ReadLine())

			tweet, err := tweetManager.GetTweetById(id)

			if tweet == nil {
				c.Printf("%s", err)
			}

			c.Println((*tweet).PrintableTweet())

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweetCount",
		Help: "Shows how many tweets a user has published",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write your username: ")

			user := c.ReadLine()

			count, _ := tweetManager.CountTweetsByUser(user)

			c.Printf("%d tweets by %s\n", count, user)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweets",
		Help: "Shows a list of the tweets a user has published",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write your username: ")

			user := c.ReadLine()

			tweets, err := tweetManager.GetTweetsByUser(user)

			if err != nil {
				c.Printf("Error: %s\n", err)
				return
			}

			for _, tweet := range tweets {
				c.Printf("Tweet %d: %s\n", tweet.GetTweetId(), tweet.PrintableTweet())
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweetsContainingText",
		Help: "Shows a list of the tweets with a given token",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write your match query: ")

			query := c.ReadLine()

			searchResult := make(chan domain.Tweet, 100)
			tweetManager.SearchTweetsContaining(query, searchResult)

			go func() {
				for tweet := range searchResult {
					c.Println(tweet.PrintableTweet())
				}
			}()

			if searchResult == nil {
				c.Printf("No tweets containing '%s' found\n", query)
				return
			}

			return
		},
	})

	shell.Run()

}
