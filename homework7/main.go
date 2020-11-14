package main

import (
	"GoWebCourse/homework7/db"
	"fmt"

	"github.com/gin-gonic/gin"
)


var mongo DBMongo

func init() {
	mongo := &DBMongo{
		DBInfo: DBInfo{URI:  "mongodb://localhost:27017",
			Name: "posts",
		},
		CollectionName: "posts",
	}
}

func main() {
	defer func() {
		mongo.Disconnect()
	}()

	router := gin.Default()

	// router.SetHTMLTemplate(template.Must(template.New("MyTemplate").ParseFiles("static/tmpl.html")))
	router.LoadHTMLGlob("static/*")
	router.GET("/lists", listTaskLists)

	//Шаблон со списком всех постов / короткие без Text
	router.GET("/", listPostHandler)
	router.POST("/delete/:id", deletePostHandlerPost)

	//Шаблон с текстовыми полями для задания  Title Text Author
	router.GET("/create", createPostHandlerGet)
	router.POST("/create", createPostHandlerPost)

	//Шаблон со страницей одного поста / полгого с отображением Text
	router.GET("/post/:id", getPostHandlerID)

	//Шаблон с текстовыми полями для обновления Title Text Author
	router.GET("/edit/:id", updatePostHandleGet)
	router.POST("/edit/:id", updatePostHandlePut)

	port := ":8094"
	fmt.Printf(" start server: %s", port)
	// tryMongoDB()
	router.Run(port)


	db.UseExample()
}

