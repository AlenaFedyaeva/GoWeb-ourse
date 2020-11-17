package main

import (
	"GoWebCourse/homework7/db"
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

var mongo db.DBMongo

func init() {

}

// @title Posts / my blog
// @version 1.0
// @description This is blog

// @contact.name Alena Fedyaeva
// @contact.email  placeholder@gmail.com

// @host localhost
// @BasePath /
func main() {
	mongo := &db.DBMongo{
		DBInfo: db.DBInfo{URI: "mongodb://localhost:27017",
			Name: "posts",
		},
		CollectionName: "posts",
	}
	mongo.DBInit()
	defer func() {
		mongo.Disconnect()
	}()

	router := gin.Default()

	router.LoadHTMLGlob("static/*")

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
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := ":8095"
	fmt.Printf(" start server: %s", port)
	// tryMongoDB()
	router.Run(port)

	// db.UseExample()
}
