package main

import (
	"GoWebCourse/homework7/db"
	_ "GoWebCourse/homework7/docs"
	"fmt"
)

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
	
	port := ":8095"
	fmt.Printf(" start server: %s", port)
	c := &db.Controller{
		ControllerDB: mongo,
	}

	srv := db.SetupServer(c)
	srv.Run(port)
}
