package main

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	DB_NAME       = "posts"
	DB_COLLECTION = "posts"
)

//insertPosts - вставляем пост в БД

func updatePostsMap() {
	postsNew, err := selectAll()
	if err != nil {
		fmt.Println(err)
	}
	posts = postsNew
}


func selectPost(id int) (Post, error) {
	post := Post{}
	// var postBson bson.D
	// hex, err := bson.ObjectIdHex(3)
	// if err != nil {
	// 	return Post{}, err
	// }
	filter := bson.M{"_id": id}

	collection := client.Database(DB_NAME).Collection(DB_COLLECTION)

	if err := collection.FindOne(context.Background(), filter).Decode(&post); err != nil {
		log.Println(err)
		return post, err
	}
	fmt.Println("selectPost: ", post)

	return post, nil
}

// primitive.NewObjectID()
func insertPost(post Post) error {

	collection := client.Database(DB_NAME).Collection(DB_COLLECTION)

	insertResult, err := collection.InsertOne(context.TODO(), post)
	if err != nil {
		log.Println("InsertOne ERROR:", err)
		return err
	}
	// log.Println(insertResult)

	// get the inserted ID string
	newID := insertResult.InsertedID
	fmt.Println("InsertOne() newID:", newID)
	fmt.Println("InsertOne() newID type:", reflect.TypeOf(newID))
	//2) update map values
	updatePostsMap()

	return nil
}

func selectAll() (map[int]*Post, error) {
	resMap := make(map[int]*Post)

	filter := bson.D{{}}
	collection := client.Database(DB_NAME).Collection(DB_COLLECTION)

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Println("selectAll find", err)
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var post Post
		if err = cursor.Decode(&post); err != nil {
			log.Println("selectAll ", err)
		}
		// log.Println("episode: ",post)
		resMap[post.Id] = &post
	}
	return resMap, err
}

func updateRow(id int, p Post) error {
	//1) update bd
	p.UpdatedAt = time.Now()
	colQuerier := bson.M{"_id": id}

	change := bson.M{"$set": bson.M{"title": p.Title,
		"text":      p.Text,
		"author":    p.Author,
		"updatedat": time.Now()}}

	collection := client.Database(DB_NAME).Collection(DB_COLLECTION)
	_, err := collection.UpdateOne(context.Background(), colQuerier, change)
	if err != nil {
		log.Println(err)
		return err
	}

	//2) update map values
	updatePostsMap()
	return nil
}

func deleteRow(id int) error {
	//1) update bd
	colQuerier := bson.M{"_id": id}
	collection := client.Database(DB_NAME).Collection(DB_COLLECTION)

	res, err := collection.DeleteOne(context.Background(), colQuerier)
	fmt.Println("DeleteOne Result TYPE:", reflect.TypeOf(res))

	if err != nil {
		log.Println("DeleteOne() ERROR:", err)
		return err
	}
	// Check if the response is 'nil'
	if res.DeletedCount == 0 {
		log.Println("DeleteOne() document not found:", res)
	} else {
		// Print the results of the DeleteOne() method
		log.Println("DeleteOne Result:", res)
		// *mongo.DeleteResult object returned by API call
		log.Println("DeleteOne TYPE:", reflect.TypeOf(res))
	}
	//2) update map values
	updatePostsMap()
	return nil
}

func tryMongoDB() {
	// 1) insert
	post, _ := posts[1]
	insertPost(*post)
	// 2) select & update map
	updatePostsMap()
	// 3) select one post
	postSelect, _ := selectPost(2)
	fmt.Println("6 lab posts ", posts, "selectPost", postSelect)
	//4) update one post
	postU := Post{Title: "Title updated", Text: "3 updated text", Author: "me"}
	updateRow(3, postU)
	// 5) delete one post
	deleteRow(1)
}
