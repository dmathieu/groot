# Groot

Regularly fetches twitter stats and sends them to librato.

### Installing

Groot is meant to run on heroku. Installing it is pretty easy.

You will first need to create a [twitter oauth app](https://apps.twitter.com/app/new).  
With the created app, you will also need to create your own access token so you can be authenticated into your account.

```
heroku create

heroku config:set BUILDPACK_URL=https://github.com/kr/heroku-buildpack-go.git#go1.2
heroku config:set LIBRATO_EMAIL=<your librato email>
heroku config:set LIBRATO_TOKEN=<your librato token>
heroku config:set TWITTER_ACCESS_SECRET=<your twitter oauth app access secret>
heroku config:set TWITTER_ACCESS_TOKEN=<your twitter oauth app access token>
heroku config:set TWITTER_ACCOUNTS=<the twitter accounts you want to track, separated by a comma>
heroku config:set TWITTER_OAUTH_KEY=<your twitter oauth key>
heroku config:set TWITTER_OAUTH_SECRET<your twitter oauth secret>

git push heroku
```

That's all!

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request

## License

Groot is released under the MIT license. See [LICENSE](LICENSE)
