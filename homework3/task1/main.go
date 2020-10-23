package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/gorilla/mux"
)

//var tmpl = template.Must(template.New("MyTemplate").ParseFiles("static/tmpl.html"))

type Post struct{
	Id int
	Title string
	Text string
	Author string
	CreatedAt time.Time
	UpdatedAt time.Time
}

var posts = map [int]*Post{
	1:&(Post{
		Id:1,
		Title:"some text1",
		Text: "some text1",
		Author: "some author1",
		CreatedAt: time.Now().Add(-time.Hour),
	}),
	2:&(Post{
		Id:2,
		Title:"some text2",
		Text: "some text2",
		Author: "some author2",
		CreatedAt: time.Now().Add(-time.Hour),
	}),
	3:&(Post{
		Id:3,
		Title:"some text3",
		Text: "some text3",
		Author: "some author3",
		CreatedAt: time.Now().Add(-time.Hour),
	}),
}

func main()  {
	fmt.Println("as")
	router := mux.NewRouter()
	//Шаблон со списком всех постов / короткие без Text
	//router.HandleFunc("/", listPostHandler).Methods("GET")

	//Шаблон ст текстовыми полями для задания  Title Text Author
	//router.HandleFunc("/", createPostHandler).Methods("POST")

	//Шаблон со страницей одного поста / полгого с отображением Text
	router.HandleFunc("/{id}", getPostHandlerID).Methods("GET")

	//Шаблон с текстовыми полями для обновления Title Text Author
	//router.HandleFunc("/{id}", updatePostHandleID).Methods("PUT")
    //r.HandleFunc("/articles", ArticlesHandler)
    //http.Handle("/", r)
	
	for key, value := range posts {
		fmt.Println("Key:", key, "Value:", value)
	}
	port:=":8099"
	fmt.Printf("Start server : port = %s",port)
	log.Fatal(http.ListenAndServe(port,router))
}

//GET
func getPostHandlerID(w http.ResponseWriter, r *http.Request){
	vars:=mux.Vars(r)
	postIDRaw,ok:= vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return	
	}

	postID,err:=strconv.Atoi(postIDRaw)
	if err!=nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	post,ok:=posts[postID]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	tmpl:=template.Must(template.New("first").Parse(`
	{{define "T"}}
	<html>
		<head>
			<title> {{.Title}} </title>
		</head>

		<body> 
			<h1> {{.Title}} </h1>
			<p> {{.Text}} </p>
		</body>
	</html>
	{{end}}
	`))


	if err:= tmpl.ExecuteTemplate(w,"T",post); err!= nil{
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
	
}
//POST
func createPostHandler(w http.ResponseWriter, r *http.Request){

}
//PUT
func updatePostHandleID(w http.ResponseWriter, r *http.Request){

}

//GET
func  listPostHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("HomeHandler")
	vars := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Category: %v\n", vars["category"])
}