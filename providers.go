package main

import (
  medium "github.com/sanketsaurav/medium-sdk-go"
  "github.com/spf13/viper"
  "log"
  "fmt"
  "database/sql"
	_ "github.com/go-sql-driver/mysql"
)


func GetPostsFromGhost() []medium.CreatePostOptions {

  // Read Ghost config
  err := viper.ReadInConfig()
  if err != nil {
    log.Fatal(err)
  }

  ghostDB := viper.Get("GHOST_DB_TYPE").(string)
  ghostDSN := fmt.Sprintf(
                "%s:%s@tcp(%s:3306)/%s",
                viper.Get("GHOST_DB_USERNAME"),
                viper.Get("GHOST_DB_PASSWORD"),
                viper.Get("GHOST_DB_HOST"),
                viper.Get("GHOST_DB_NAME"),
              )

  db, err := sql.Open(ghostDB, ghostDSN)
  if err != nil {
    panic(err.Error())
  }
  defer db.Close()

  // Fetch all posts
  var (
    title string
    published_at string
    html string
  )
  rows, err := db.Query("select title, published_at, html from posts")
  if err != nil {
    log.Fatal(err)
  }
  defer rows.Close()

  var posts []medium.CreatePostOptions
  // Print some data
  for rows.Next() {
    err := rows.Scan(&title, &published_at, &html)
    if err != nil {
      log.Fatal(err)
    }

    // Create a medium.CreatePostOptions struct for each entity
    posts = append(posts, medium.CreatePostOptions{
      Title: title,
      PublishedAt: published_at,
      Content: html,
      ContentFormat: medium.ContentFormatHTML,
      PublishStatus: medium.PublishStatusDraft,
    })

  }
  return posts
}
