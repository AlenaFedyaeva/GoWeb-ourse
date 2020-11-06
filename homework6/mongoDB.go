package main

import (
	"context"
	"log"
)

const dbName = "posts"

// type Post struct {
// 	Title string `json:”title,omitempty”`
// 	Body string `json:”body,omitempty”`
// }

func InsertPost(p Post) error {{

	post:= posts[0]

	collection := client.Database("posts").Collection("posts")

	insertResult, err := collection.InsertOne(context.TODO(), post)

	log.Println(insertResult)
	if err != nil {
		log.Println(err)
	}
	//2) update map values
	updatePostsMap()

	return nil
}
