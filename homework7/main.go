package main

import (
	"GoWebCourse/homework7/db"
)

// import (
// "context"
// "log"
// "net/http"

// "github.com/gin-gonic/gin"
// "go.mongodb.org/mongo-driver/mongo"
// "go.mongodb.org/mongo-driver/mongo/options"
// )

// var collection *mongo.Collection
// var client *mongo.Client

// func init() {
// 	db, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	client = db
// 	err = client.Connect(context.Background())
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	err = client.Ping(context.Background(), nil)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	collection = client.Database("posts").Collection("postlist")

// }

func main() {
	// defer func() {
	// 	if err := client.Disconnect(context.Background()); err != nil {
	// 		log.Println(err)
	// 	}
	// }()

	// router := gin.Default()

	// // router.SetHTMLTemplate(template.Must(template.New("MyTemplate").ParseFiles("static/tmpl.html")))
	// router.LoadHTMLGlob("static/*")
	// router.GET("/lists", listTaskLists)

	// //Шаблон со списком всех постов / короткие без Text
	// router.GET("/", listPostHandler)
	// router.POST("/delete/:id", deletePostHandlerPost)

	// //Шаблон с текстовыми полями для задания  Title Text Author
	// router.GET("/create", createPostHandlerGet)
	// router.POST("/create", createPostHandlerPost)

	// //Шаблон со страницей одного поста / полгого с отображением Text
	// router.GET("/post/:id", getPostHandlerID)

	// //Шаблон с текстовыми полями для обновления Title Text Author
	// router.GET("/edit/:id", updatePostHandleGet)
	// router.POST("/edit/:id", updatePostHandlePut)

	// port := ":8094"
	// fmt.Printf(" start server: %s", port)
	// // tryMongoDB()
	// router.Run(port)

	
	db.UseExample()
}
