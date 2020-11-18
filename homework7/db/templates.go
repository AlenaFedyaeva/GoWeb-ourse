package db

import (
	"time"
)

type Post struct {
	Id        int	 `bson:"_id" json:"id"`
	Title     string `schema:"title,title2example" json:"title" xml:"title`
	Text      string `schema:"text" json:"text" xml:"text`
	Author    string `schema:"author" json:"author" xml:"author`
	CreatedAt time.Time `json:"cratedAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type AllPostsStruct struct {
	Title string
	Data  map[int]*Post
}

type OnePostsStruct struct {
	Title string
	Data  Post
}

var Posts  map[int]*Post
var PostsDefault = map[int]*Post{
	1: &(Post{
		Id:        1,
		Title:     "some text1",
		Text:      "some text1",
		Author:    "some author1",
		CreatedAt: time.Now().Add(-time.Hour),
	}),
	2: &(Post{
		Id:        2,
		Title:     "some text2",
		Text:      "some text2",
		Author:    "some author2",
		CreatedAt: time.Now().Add(-time.Hour),
	}),
	3: &(Post{
		Id:        3,
		Title:     "some text3",
		Text:      "some text3",
		Author:    "some author3",
		CreatedAt: time.Now().Add(-time.Hour),
	}),
}
// temporary ! need - primitive.NewObjectID().
var (
	postID = 10
)

func GetPostId() int {
	postID++
	return postID
}
