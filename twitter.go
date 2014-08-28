package main

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"net/url"
	"os"
	"time"
)

type TwitterAccount struct {
	username string
}

type Twitter struct {
	OauthKey     string
	OauthSecret  string
	AccessToken  string
	AccessSecret string
	Client       *anaconda.TwitterApi
}

func BuildTwitter() *Twitter {
	client := &Twitter{
		OauthKey:     os.Getenv("TWITTER_OAUTH_KEY"),
		OauthSecret:  os.Getenv("TWITTER_OAUTH_SECRET"),
		AccessToken:  os.Getenv("TWITTER_ACCESS_TOKEN"),
		AccessSecret: os.Getenv("TWITTER_ACCESS_SECRET"),
	}
	anaconda.SetConsumerKey(client.OauthKey)
	anaconda.SetConsumerSecret(client.OauthSecret)
	client.Client = anaconda.NewTwitterApi(client.AccessToken, client.AccessSecret)

	return client
}

func (t *Twitter) fetchUser(account string, collector *Metrics) {
	user, err := t.Client.GetUsersShow(account, url.Values{})
	if err != nil {
		Log("Couldn't fetch twitter data")
	} else {
		Log("account.data name=%s followers=%d following=%d", account, user.FollowersCount, user.FriendsCount)

		go collector.TrackGauge(fmt.Sprintf("twitter.%s.followers", account), int64(user.FollowersCount))
		go collector.TrackGauge(fmt.Sprintf("twitter.%s.following", account), int64(user.FriendsCount))
	}

	searchParams, _ := url.ParseQuery("count=100")
	search, err := t.Client.GetSearch(fmt.Sprintf("@%s", account), searchParams)
	if err != nil {
		Log("Couldn't fetch twitter search data")
	} else {
		yesterday := time.Now().Add(-MentionsKeep)
		searchCount := 0

		for _, result := range search {
			createdAt, _ := result.CreatedAtTime()
			if createdAt.After(yesterday) {
				searchCount += 1
			}
		}

		Log("account.search name=%s results=%d total=%d", account, searchCount, len(search))
		go collector.TrackGauge(fmt.Sprintf("twitter.%s.mentions", account), int64(searchCount))
	}
}
