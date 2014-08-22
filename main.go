package main

import (
	"os"
	"strings"
	"time"
)

func main() {
	Log("main.start")

	accounts := strings.Split(os.Getenv("TWITTER_ACCOUNTS"), ",")
  startReporter()
	go monitorStart(accounts, FetchInterval)

	<-make(chan bool)
}

func monitorStart(accounts []string, interval int) {
	Log("monitor.start")
	twitter := BuildTwitter()
	collector, _ := NewMetrics()

	for {
		Log("monitor.tick accounts=%s", accounts)
		for _, account := range accounts {
			Log("account.fetch account=%s", account)
			go twitter.fetchUser(account, collector)
		}

		<-time.After(time.Duration(interval) * time.Second)
	}
}
