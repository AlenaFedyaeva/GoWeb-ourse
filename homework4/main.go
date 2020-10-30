package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/schema"
)

func getPostId() int {
	postID++
	return postID
}
var (
	// Set a Decoder instance as a package global, because it caches
	// meta-data about structs, and an instance can be shared safely.
		decoder   = schema.NewDecoder()
		postID    = 3
		database *sql.DB
		// templates = template.Must(template.ParseFiles("./static/tmpl_1page.html", "./static/tmpl.html",
		// 	"./static/tmpl_allPosts.html", "./static/tmpl_edit.html", "./static/tmpl_create.html")) //Template caching
	)
	
type Post struct {
	Id        int
	Title     string `schema:"title,title2example" json:"title" xml:"title`
	Text      string `schema:"text" json:"text" xml:"text`
	Author    string `schema:"author" json:"author" xml:"author`
	CreatedAt time.Time
	UpdatedAt time.Time
}

var posts = map[int]*Post{
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


func main(){

	fmt.Println("task 4 starting")

	dbConnect()

	fmt.Println("task 4 bye")
}