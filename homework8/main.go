package main

import (
	"GoWebCourse/homework8/config"
	"GoWebCourse/homework8/log"
	"fmt"
	"log"

	"github.com/rs/zerolog/log"

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
	var cfg config.FileConfig

	debug := flag.Bool("debug", false, "sets log level to debug")
	portF := flag.String("port","9090","a string")

	flag.Parse()

	// 1) Read port from args 
	if *portF==""{
		log.Fatal("Empty value err")
	}
	fmt.Println("debug: ", *debug,"port: ",*portF)

    //NOTE: Использую сразу все варианты загрузки конфигурации, в учебных целях. В таком простом примере можно было бы оставить только Env например)
	// 2) Read config file name from Environment Variable
	fname,okEnv:=config.LookupEnv("CONF_HW8") 
	if !okEnv{
		log.Fatal("LookupEnv err")
	}
	
	//3) Read config file 
	errFile:=cfg.ReadConfigJson(fname)
	if errFile!=nil{
		log.Fatal("Parse args :",errFile)
	}
	cfg.Port=*portF
	// 4) INIT mongo db
	mongo := &db.DBMongo{
		DBInfo: db.DBInfo{URI: cfg.DB.URI,
			Name: cfg.DB.Name,
		},
		CollectionName: cfg.CollectionName,
	}
	mongo.DBInit()
	defer func() {
		mongo.Disconnect()
	}()
	//  5) INIT logger
	logger,errLogger:=log.NewFileLogger(cfg.LogName)
	if errLogger!=nil{
		log.Fatal("NewFileLogger:",errLogger)
	}

	port := ":"+*portF
	fmt.Printf(" start server: %s", port)
	c := &db.Controller{
		ControllerDB: mongo,
	}

	srv := db.SetupServer(c)
	srv.Run(port)
}
