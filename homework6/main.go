package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var client *mongo.Client

func init() {
	db, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	client = db
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Println(err)
	}
	collection = client.Database("posts").Collection("postlist")
}

func main() {
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Println(err)
		}
	}()

	router := gin.Default()

	router.SetHTMLTemplate(template.Must(template.New("MyTemplate").ParseFiles("static/tmpl.html")))

	router.GET("/lists", listTaskLists)
	// router.GET("/list", getTaskList)
	// router.POST("/lists/add", createTaskList)
	// router.GET("/lists/add", createTaskListForm)
	// router.POST("/lists/edit", updateTaskList)
	// router.GET("/lists/edit", updateTaskListForm)
	port := ":8092"
	fmt.Printf(" start server: %s", port)

	InsertPost("title1", "body1")

	fmt.Println("6 lab")

	
	router.Run(port)

}

func listTaskLists(c *gin.Context) {
	// lists, err := GetLists()
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
	c.HTML(http.StatusOK, "alllists", nil)
}
