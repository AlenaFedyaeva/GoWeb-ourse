package main

import (
	"GoWebCourse/homework8/config"
	"fmt"
	"log"

	"GoWebCourse/homework8/db"
	_ "GoWebCourse/homework8/docs"

	"flag"
)

// @title Posts / my blog
// @version 1.0
// @description This is blog

// @contact.name Alena Fedyaeva
// @contact.email  placeholder@gmail.com

// @host localhost
// @BasePath /
func main() {
	var conf config.FileConfig
    //NOTE: Использую сразу все варианты загрузки конфигурации, в учебных целях. В таком простом примере можно было бы оставить только Env например)
	// 1) Read config file name from Environment Variable
	fname,okEnv:=config.LookupEnv("CONF_HW8") 
	if !okEnv{
		log.Fatal("LookupEnv err")
	}

	// 2) Read port from args 
	portF := flag.String("port","9090","a string")
    flag.Parse()

	if *portF==""{
		log.Fatal("Empty value err")
	}
	//3) Read config file 
	errFile:=conf.ReadConfigJson(fname)
	if errFile!=nil{
		log.Fatal("Parse args :",errFile)
	}
	conf.Port=*portF
	// 4) INIT mongo db
	mongo := &db.DBMongo{
		DBInfo: db.DBInfo{URI: conf.DB.URI,
			Name: conf.DB.Name,
		},
		CollectionName: conf.CollectionName,
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
