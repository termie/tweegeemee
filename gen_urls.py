import os
import pprint

import twitter

def main():

  consumer_key=os.environ.get("TWITTER_CONSUMER_KEY")
  consumer_secret=os.environ.get("TWITTER_CONSUMER_SECRET")
  access_token=os.environ.get("TWITTER_ACCESS_TOKEN")
  access_token_secret=os.environ.get("TWITTER_ACCESS_TOKEN_SECRET")


  api = twitter.Api(consumer_key=consumer_key,
                    consumer_secret=consumer_secret,
                    access_token_key=access_token,
                    access_token_secret=access_token_secret)

  favs = api.GetFavorites(screen_name='termie')
  #pprint.pprint(favs)
  urls = []
  for fav in favs:
    if fav.user.screen_name != 'tweegeemee':
      continue
    urls.append(fav.urls[0].expanded_url)

  for u in urls:
    print u

if __name__ == '__main__':
    main()
