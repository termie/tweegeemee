package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

func main() {
	consumerKey := os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")
	// applicationID := os.Getenv("APPLICATION_ID")
	// pipelineID := os.Getenv("PIPELINE_ID")
	// werckerEndpoint := os.Getenv("WERCKER_ENDPOINT")

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	anaconda.SetConsumerKey(consumerKey)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	// clientConfig := werckerclient.Config{
	//   Endpoint: werckerEndpoint,
	// }
	// client := werckerclient.NewClient(&clientConfig)
	// opts := werckerclient.CreateRunOptions{
	//   // ApplicationID: applicationID,
	//   PipelineID: pipelineID,
	//   Branch:     "werckerize",
	// }

	v := url.Values{}
	v.Set("screen_name", "termie")

	favs, err := api.GetFavorites(v)
	if err != nil {
		panic(err)
	}

	for _, f := range favs {
		user := f.User.ScreenName
		// fmt.Printf("User: %s\n", user)

		if user == "tweegeemee" {
			// fmt.Printf("%s\n", f.Text)
			for _, e := range f.Entities.Urls {
				fmt.Printf("%s\n", e.Expanded_url)
			}
			// url := f.Entities.Urls[0].Expanded_url
			// fmt.Printf(url
			// opts.EnvVars = werckerclient.EnvVarFromMap(map[string]string{
			//   "TGM_URL": url,
			// })
			// fmt.Printf("%+v\n", opts)
			// r, err := client.CreateRun(&opts)
			// if err != nil {
			//   panic(err)
			// }
			// fmt.Printf("%+v\n", r)

			// return
		}
	}

}
