package main

import (
	"GoWebCourse/homework7/db"
	_ "GoWebCourse/homework7/docs"
	"flag"
	"fmt"
	"os"
)

// @title Posts / my blog
// @version 1.0
// @description This is blog

// @contact.name Alena Fedyaeva
// @contact.email  placeholder@gmail.com

// @host localhost
// @BasePath /



type fileSave struct{
	DB db.DBInfo 
	CollectionName string
	port string
}

func main() {
	fmt.Print(os.Args)
	argsWithoutProg := os.Args[1:]
	fmt.Println(argsWithoutProg)
	portF := flag.String("port","9090","a string")
	configF := flag.String("config","/","a string")
	flag.Parse()
	fmt.Println("port", *portF,portF)
	fmt.Println("config", *configF)	
	
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
	
	port := ":"+*portF
	fmt.Printf(" start server: %s", port)
	c := &db.Controller{
		ControllerDB: mongo,
	}

	srv := db.SetupServer(c)
	srv.Run(port)
}
