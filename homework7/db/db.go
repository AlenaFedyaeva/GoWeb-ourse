package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	// SelectPost(id int) (Post, error)
	SelectAll() (map[int]*Post, error)
	UpdatePostsMap()
	DBInit()
	Disconnect()
}

func (db *DBMongo) DBInit() {
	dbCli, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	db.Client = dbCli
	err = db.Client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	err = db.Client.Ping(context.Background(), nil)
	if err != nil {
		log.Println(err)
	}
	db.Collection = db.Client.Database(db.Name).Collection(db.CollectionName)
}

func (db *DBMongo) Disconnect(){
	if err := db.Client.Disconnect(context.Background()); err != nil {
		log.Println(err)
	}
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
	mongo.DBInit()
	mongo.UpdatePostsMap()
	fmt.Println(posts,mongo)
}
