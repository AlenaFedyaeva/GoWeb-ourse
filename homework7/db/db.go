package db

import (
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type DBInfo struct {
	Name string
	URI  string
}

type DBMongo struct {
	DBInfo
	CollectionName string
	Collection     *mongo.Collection
	Client         *mongo.Client
}

type DB interface {
	SelectAll() (map[int]*Post, error)
	SelectPost(id int) (Post, error)
	InsertPost(post Post) (int,error)
	UpdateRow(id int, p Post) error 
	DeleteRow(id int) error
	UpdatePostsMap()
	DBInit() error
	Disconnect()
}
type Controller struct {
	ControllerDB DB
}


func UseExample() {
	fmt.Println("use example")
	mongo := &DBMongo{
		DBInfo: DBInfo{URI:  "mongodb://localhost:27017",
			Name: "posts",
		},
		CollectionName: "posts",
	}
	defer func() {
		mongo.Disconnect()
	}()

	// 1) init, select all posts
	err:=mongo.DBInit()
	if err!=nil{
		log.Fatal(err)
	}
	// 2) insert one
	post, _ := Posts[1]
	mongo.InsertPost(*post)
	// 3)  select all posts
	mongo.UpdatePostsMap()
	// 4) select 1 post
	postSelect, _ := mongo.SelectPost(3)
	fmt.Println("7 lab posts ", Posts, "selectPost", postSelect)
	// 5) update one post
	postU := Post{Title: "Title updated", Text: "3 updated text", Author: "me"}
	mongo.UpdateRow(4, postU)
	// 6) del row
	mongo.DeleteRow(1) 
	fmt.Println("end",Posts,mongo)
}
