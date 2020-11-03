package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)


var (
	// Set a Decoder instance as a package global, because it caches
	// meta-data about structs, and an instance can be shared safely.
	decoder   = schema.NewDecoder()
	postID    = 3
	database  *sql.DB
	// templates = template.Must(template.ParseFiles("./static/tmpl_1page.html", "./static/tmpl.html",
		// "./static/tmpl_allPosts.html", "./static/tmpl_edit.html", "./static/tmpl_create.html")) //Template caching
)

func main() {

	fmt.Println("task 4 starting")
	db, err := sql.Open("mysql", "root:my-secret-pw@/posts?parseTime=true")

	if err != nil {
		log.Println("bd conn")
		log.Println(err)
	}
	database = db

	if err := db.Ping(); err != nil {
		log.Println(err)
	}
	defer database.Close()

	//tryBD()

	router := mux.NewRouter()//.StrictSlash(true)
	//Шаблон со списком всех постов / короткие без Text
	router.HandleFunc("/", listPostHandler).Methods("GET")

	//Шаблон с текстовыми полями для задания  Title Text Author
	//router.HandleFunc("/create", createPostHandlerGet).Methods("GET")
	//router.HandleFunc("/create", createPostHandlerPost).Methods("POST")

	//Шаблон со страницей одного поста / полгого с отображением Text
	//router.HandleFunc("/{id}", getPostHandlerID).Methods("GET")

	//Шаблон с текстовыми полями для обновления Title Text Author
	// router.HandleFunc("/edit/{id}", updatePostHandleGet).Methods("GET")
	// router.HandleFunc("/edit/{id}", updatePostHandlePut).Methods("POST")
	// router.HandleFunc("/edit/{id}", updatePostHandlePut).Methods("PUT")
	port := ":9090"
	fmt.Printf("Start server : port = %s", port)
	// http.Handle("/", router)
	log.Fatal(http.ListenAndServe(port, router))


}
