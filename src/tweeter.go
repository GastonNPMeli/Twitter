package main

import (
	"github.com/GastonNPMeli/Twitter/src/domain"
	"github.com/GastonNPMeli/Twitter/src/service"
	"github.com/abiosoft/ishell"
	"strconv"
)

func main() {

	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write your user: ")

			user := c.ReadLine()

			c.Print("Write your tweet: ")

			text := c.ReadLine()

			var tweet = domain.NewTweet(user, text)

			service.PublishTweet(tweet)

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

			tweet, err := service.GetTweetById(id)

			if tweet == nil {
				c.Printf("%s", err)
			}

			c.Println(tweet)

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

			count, err := service.CountTweetsByUser(user)

			c.Printf("Error: %s\n", err)
			return

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

			tweets, err := service.GetTweetsByUser(user)

			if err != nil {
				c.Printf("Error: %s\n", err)
				return
			}

			for _, tweet := range tweets {
				c.Printf("User %s tweeted '%s' at %s\n", tweet.User, tweet.Text, tweet.Date)
			}

			return
		},
	})

	shell.Run()

}
