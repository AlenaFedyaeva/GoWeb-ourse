package main

import (
	"GoWebCourse/homework8/config"
	newlog "GoWebCourse/homework8/newlog"
	"fmt"
	"log"

	// "github.com/rs/zerolog/log"

	"GoWebCourse/homework8/db"
	_ "GoWebCourse/homework8/docs"
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
	// 1) Read flags
	cfg.ParseFlags()

	//NOTE: Использую сразу все варианты загрузки конфигурации, в учебных целях.
	// 2) Read config file name from Environment Variable
	fname,err:=config.GetFileNameFromEnv("CONF_HW8") 
	if err!=nil {
		log.Fatal(err)
	}
	
	//3) Read config file 
	errFile:=cfg.ReadConfigJson(fname)
	if errFile!=nil{
		log.Fatal("Parse args :",errFile)
	}

	//  4) INIT logger
	logger,errLogger:=newlog.NewFileLogger(cfg.LogName)
	if errLogger!=nil{
		log.Fatal("NewFileLogger:",errLogger,logger)
	}
	logger.Info().Str("port", cfg.Port).Msg("start server")

	// 5) INIT mongo db
	mongo := &db.DBMongo{
		DBInfo: db.DBInfo{URI: cfg.DB.URI,
			Name: cfg.DB.Name,
		},
		CollectionName: cfg.CollectionName,
		Logger: logger,
	}
	mongo.DBInit()
	defer func() {
		mongo.Disconnect()
	}()


	fmt.Printf(" start server: %s", cfg.Port)
	c := &db.Controller{
		ControllerDB: mongo,
		Logger: logger,
	}

	srv := db.SetupServer(c)
	srv.Run(fmt.Sprintf(":%s",cfg.Port))
}
