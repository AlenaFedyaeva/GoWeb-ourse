package main

import (
	"context"
	"log"
)

const dbName = "posts"

type Post struct {
	Title string `json:”title,omitempty”`

	Body string `json:”body,omitempty”`
}

func InsertPost(title string, body string) {

	post := Post{title, body}

	collection := client.Database("posts").Collection("posts")

	insertResult, err := collection.InsertOne(context.TODO(), post)

	log.Println(insertResult)
	if err != nil {
		log.Println(err)
	}
}

//************************

// func insert(){

// 	collection := client.Database(dbName).Collection("posts")

// 	//  db.users.insertOne({ name: 'Merrick', email: 'test@gmail.com'})

// 	// db.users.insertOne({ name: 'Merrick', email: 'test@gmail.com'})

// 	// Метод insertOne добавляет к документу поле _id
// 	// если оно отсутствует и добавляет документ в коллекцию.

// }

// func GetLists() ([]TaskList, error) {
// 	var res []TaskList

// 	collection := client.Database(dbName).Collection("lists")

// 	cur, err := collection.Find(context.Background(), bson.D{})
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = cur.All(context.TODO(), &res)
// 	return res, err
// }

// // CreateList — создание листа задач
// func CreateList(list TaskList) (TaskList, error) {
// 	list.ID = primitive.NewObjectID()

// 	collection := client.Database(dbName).Collection("lists")
// 	_, err := collection.InsertOne(context.Background(), &list)
// 	return list, err
// }

// //*---------------------
