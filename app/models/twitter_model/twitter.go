package twitter_model

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
)

var appClient *twitter.Client
var userClient *twitter.Client

func init() {
	appClient = GetAuthAppClient()
	userClient = GetAuthUserClient()
}

// search tweets
func SearchTweets(q string) []twitter.Tweet {
	searchTweetParams := &twitter.SearchTweetParams{
		Query:     q + " filter:images",
		ResultType: "mixed",
		TweetMode: "extended",
		Count:     10,
	}
	search, _, _ := userClient.Search.Tweets(searchTweetParams)
	fmt.Println("ツイート数:", len(search.Statuses))
	var tweets []twitter.Tweet
	tweets = append(tweets, search.Statuses...)
	return tweets
}