package db

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// UpdatePostsMap - обновляем map со значениями
func (db *DBMongo) UpdatePostsMap() {
	postsNew, err := db.SelectAll()
	if err != nil {
		log.Println(err)
	}
	Posts = postsNew
}

// SelectAll - выбираем все значения из коллекции
func (db *DBMongo) SelectAll() (map[int]*Post, error) {
	resMap := make(map[int]*Post)
	filter := bson.D{{}}

	cursor, err := db.Collection.Find(context.Background(), filter)
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

// SelectPost - Выбираем один пост из коллекции
func (db *DBMongo) SelectPost(id int) (Post, error) {
	post := Post{}
	filter := bson.M{"_id": id}

	if err := db.Collection.FindOne(context.Background(), filter).Decode(&post); err != nil {
		log.Println(err)
		return post, err
	}
	fmt.Println("selectPost: ", post)

	return post, nil
}

// InsertPost - вставка одного поста в коллекцию
func (db *DBMongo) InsertPost(post Post) error {
	_, err := db.Collection.InsertOne(context.TODO(), post)
	if err != nil {
		log.Println("InsertOne ERROR:", err)
		return err
	}
	
	// get the inserted ID string
	// newID := insertResult.InsertedID
	// fmt.Println("InsertOne() newID:", newID)
	// fmt.Println("InsertOne() newID type:", reflect.TypeOf(newID))
	//2) update map values
	db.UpdatePostsMap()

	return nil
}

// UpdateRow - обновляем пост в коллекции
func (db *DBMongo)  UpdateRow(id int, p Post) error {
	//1) update bd
	p.UpdatedAt = time.Now()
	colQuerier := bson.M{"_id": id}

	change := bson.M{"$set": bson.M{"title": p.Title,
		"text":      p.Text,
		"author":    p.Author,
		"updatedat": time.Now()}}

	_, err := db.Collection.UpdateOne(context.Background(), colQuerier, change)
	if err != nil {
		log.Println(err)
		return err
	}

	//2) update map values
	db.UpdatePostsMap()
	return nil
}

// DeleteRow - удаляем один пост в коллекции
func  (db *DBMongo) DeleteRow(id int) error {
	//1) update bd
	colQuerier := bson.M{"_id": id}
	res, err := db.Collection.DeleteOne(context.Background(), colQuerier)
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
	db.UpdatePostsMap()
	return nil
}

// DBInit - инициализируем БД
func (db *DBMongo) DBInit() error {
	dbCli, err := mongo.NewClient(options.Client().ApplyURI(db.URI))
	if err != nil {
		log.Println(err)
		return err
	}
	db.Client = dbCli
	err = db.Client.Connect(context.Background())
	if err != nil {
		log.Println(err)
		return err
	}
	err = db.Client.Ping(context.Background(), nil)
	if err != nil {
		log.Println(err)
		return err
	}
	db.Collection = db.Client.Database(db.Name).Collection(db.CollectionName)
	return nil
}

// Disconnect - close DB connect 
func (db *DBMongo) Disconnect(){
	if err := db.Client.Disconnect(context.Background()); err != nil {
		log.Println(err)
	}
}
