package domain

import (
	"strconv"
	"time"
)

type Tweet interface {
	PrintableTweet() string

	GetUser() string
	GetText() string
	GetTweetId() int
	GetDate() *time.Time

	SetUser(string)
	SetText(string)
	SetTweetId(int)
	SetDate(time.Time)
}

type TextTweet struct {
	TweetId int
	User string
	Text string
	Date *time.Time
}

type ImageTweet struct {
	TextTweet
	Image string
}

type QuoteTweet struct {
	TextTweet
	QuotedTweet Tweet
}

func NewTextTweet(user string, text string) *TextTweet {
	date := time.Now()
	return &TextTweet{ 0, user, text,  &date}
}

func NewImageTweet(user string, text string, image string) *ImageTweet {
	textTweet := NewTextTweet(user, text)
	return &ImageTweet{ *textTweet, image}
}

func NewQuoteTweet(user string, text string, quotedTweet Tweet) *QuoteTweet {
	textTweet := NewTextTweet(user, text)
	return &QuoteTweet{ *textTweet, quotedTweet}
}

func (t TextTweet) PrintableTweet() string {
	return "TweetID " + strconv.Itoa(t.TweetId) +  " -> @" + t.User + ": " + t.Text
}

func (t ImageTweet) PrintableTweet() string {
	return "TweetID " + strconv.Itoa(t.TweetId) +  " -> @" + t.User + ": " + t.Text + " " + t.Image
}

func (t QuoteTweet) PrintableTweet() string {
	return "TweetID " + strconv.Itoa(t.TweetId) +  " -> @" + t.User + ": " + t.Text + " \"" + t.QuotedTweet.PrintableTweet() + "\""
}

func (t TextTweet) GetUser() string {
	return t.User
}

func (t TextTweet) GetText() string {
	return t.Text
}

func (t TextTweet) GetTweetId() int {
	return t.TweetId
}

func (t TextTweet) GetDate() *time.Time {
	return t.Date
}

func (t *TextTweet) SetUser(user string) {
	t.User = user
}

func (t *TextTweet) SetText(text string) {
	t.Text = text
}

func (t *TextTweet) SetTweetId(id int) {
	t.TweetId = id
}

func (t *TextTweet) SetDate(date time.Time) {
	t.Date = &date
}