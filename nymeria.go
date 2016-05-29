package main

import (
	medium "github.com/sanketsaurav/medium-sdk-go"
	"github.com/spf13/viper"
	"log"
)

func main() {

	// Load configuration from config.yml
	viper.AutomaticEnv()
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	// Create a Medium client with a self issued acees-toke.
	// Can be geenrated under `Integration tokens` here - https://medium.com/me/settings
	client := medium.NewClientWithAccessToken("2edb41e95c275983fa2cb833fe5a35dcf60ef2c0725ca58089ccba606720584f4")

	// Display the user
	user, err := client.GetUser()
  if err != nil {
      log.Fatal(err)
  } else {
		log.Printf("Importing stories for %s <%s>", user.Name, user.Username)
	}

	for _, post := range GetPostsFromGhost() {
		post.UserID = user.ID
		log.Printf("Importing %s", post.Title)
		p, err := client.CreatePost(post)

		if err != nil {
			log.Fatal(err)
		} else {
			log.Println(p)
		}
	}

}
