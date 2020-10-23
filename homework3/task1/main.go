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

type Post struct {
	Id        int
	Title     string
	Text      string
	Author    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

var posts = map[int]*Post{
	1: &(Post{
		Id:        1,
		Title:     "some text1",
		Text:      "some text1",
		Author:    "some author1",
		CreatedAt: time.Now().Add(-time.Hour),
	}),
	2: &(Post{
		Id:        2,
		Title:     "some text2",
		Text:      "some text2",
		Author:    "some author2",
		CreatedAt: time.Now().Add(-time.Hour),
	}),
	3: &(Post{
		Id:        3,
		Title:     "some text3",
		Text:      "some text3",
		Author:    "some author3",
		CreatedAt: time.Now().Add(-time.Hour),
	}),
}

func main() {
	fmt.Println("as")
	router := mux.NewRouter()
	//Шаблон со списком всех постов / короткие без Text
	router.HandleFunc("/", listPostHandler).Methods("GET")
	
	//Exapmple! delete
	router.HandleFunc("/list", viewList).Methods("GET")

	//Шаблон с текстовыми полями для задания  Title Text Author
	router.HandleFunc("/create", createPostHandler).Methods("GET")
	//router.HandleFunc("/create", createPostHandler).Methods("POST")

	//Шаблон со страницей одного поста / полгого с отображением Text
	router.HandleFunc("/{id}", getPostHandlerID).Methods("GET")

	//Шаблон с текстовыми полями для обновления Title Text Author
	//router.HandleFunc("/{id}", updatePostHandleID).Methods("PUT")
	//r.HandleFunc("/articles", ArticlesHandler)
	//http.Handle("/", r)


	port := ":8096"
	fmt.Printf("Start server : port = %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}

// TaskList - список задач
type TaskList struct {
	Name string
	Description string
	List []Task
  }
  
  // Task - задача и ее статус
  type Task struct {
	ID string
	Text string
	Complete bool
  }
var simpleList = TaskList{
	Name: "Название листа",
	Description: "Описание листа с задачами",
	List: []Task{
	Task{"first", "Первая задача", false},
	Task{"second", "Вторая задача", false},
	Task{"thrid", "Третья задача", true},
	},
  }
  

func viewList(w http.ResponseWriter, r *http.Request) {
	tmpl:= template.Must(template.New("Ttt").ParseFiles("./static/tmpl.html"))

	if err := tmpl.ExecuteTemplate(w, "list", simpleList); err != nil {
		log.Println(err)
	}
}

//GET
//2. роут и шаблон для просмотра конкретного поста в блоге.
func getPostHandlerID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postIDRaw, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	postID, err := strconv.Atoi(postIDRaw)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	post, ok := posts[postID]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}


	tmpl:= template.Must(template.New("tmpl_1page").ParseFiles("./static/tmpl_1page.html"))

	if err := tmpl.ExecuteTemplate(w, "list", simpleList); err != nil {
		log.Println(err)
	}

	if err := tmpl.ExecuteTemplate(w, "T", post); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}

}

//POST
//3. Создайте роут и шаблон для создания материала
func createPostHandler(w http.ResponseWriter, r *http.Request) {
	tmpl:= template.Must(template.New("tmpl_create").ParseFiles("./static/tmpl_create.html"))

	if err := tmpl.ExecuteTemplate(w, "Create", simpleList); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

}

//PUT
//3. Создайте роут и шаблон для редактированияматериала.
func updatePostHandleID(w http.ResponseWriter, r *http.Request) {

}

//GET
//1. роут и шаблон для отображения всех постов в блоге.
func listPostHandler(w http.ResponseWriter, r *http.Request) {

	for key, value := range posts {
		fmt.Println("Key:", key, "Value:", value)
	}

	tmpl:= template.Must(template.New("tmpl_1page").ParseFiles("./static/tmpl_allPosts.html"))

	if err := tmpl.ExecuteTemplate(w, "AllPosts", posts); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

}
