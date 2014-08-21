package main

import (
	"github.com/ChimeraCoder/anaconda"
	"log"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	FetchInterval = 5
)

type TwitterAccount struct {
	username string
}

type Twitter struct {
	OauthKey     string
	OauthSecret  string
	AccessToken  string
	AccessSecret string
}

func (t *Twitter) GetClient() *anaconda.TwitterApi {
	anaconda.SetConsumerKey(t.OauthKey)
	anaconda.SetConsumerSecret(t.OauthSecret)
	return anaconda.NewTwitterApi(t.AccessToken, t.AccessSecret)
}

func BuildTwitter() *Twitter {
	return &Twitter{
		OauthKey:     os.Getenv("TWITTER_OAUTH_KEY"),
		OauthSecret:  os.Getenv("TWITTER_OAUTH_SECRET"),
		AccessToken:  os.Getenv("TWITTER_ACCESS_TOKEN"),
		AccessSecret: os.Getenv("TWITTER_ACCESS_SECRET"),
	}
}

func main() {
	Log("main.start")

	accounts := strings.Split(os.Getenv("TWITTER_ACCOUNTS"), ",")
	go monitorStart(accounts, FetchInterval)

	<-make(chan bool)
}

func monitorStart(accounts []string, interval int) {
	Log("monitor.start")

	for {
		Log("monitor.tick accounts=%s", accounts)
		for _, account := range accounts {
			Log("account.fetch account=%s", account)
			go fetchUser(account)
		}

		<-time.After(time.Duration(interval) * time.Second)
	}
}

func fetchUser(account string) {
	api := BuildTwitter().GetClient()
	user, err := api.GetUsersShow(account, url.Values{})
  if err != nil {
		panic(err)
	}

  Log("account.data name=%s followers=%d following=%d", user.Name, user.FollowersCount, user.FriendsCount)
}

func Log(l string, t ...interface{}) {
	log.Printf(l, t...)
}
