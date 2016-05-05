package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/wercker/werckerclient"
)

func main() {
	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CONSUMER_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")
	applicationID := os.Getenv("APPLICATION_ID")
	targetID := os.Getenv("TARGET_ID")
	werckerEndpoint := os.Getenv("WERCKER_ENDPOINT")

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	anaconda.SetConsumerKey(consumerKey)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	clientConfig := werckerclient.Config{
		Endpoint: werckerEndpoint,
	}
	client := werckerclient.NewClient(clientConfig)
	opts := werckerclient.CreateRunOptions{
		ApplicationID: applicationID,
		TargetID:      targetID,
	}

	v := url.Values{}
	v.Set("screen_name", "termie")

	favs, err := api.GetFavorites(v)
	if err != nil {
		panic(err)
	}

	for _, f := range favs {
		user := f.User.ScreenName
		fmt.Printf("User: %s\n", user)

		if user == "tweegeemee" {
			fmt.Printf("%s\n", f.Text)
			for _, e := range f.Entities.Urls {
				fmt.Printf("  %s\n", e.Expanded_url)
			}
			return
		}
	}

}
