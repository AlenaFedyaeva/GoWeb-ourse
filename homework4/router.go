package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"

	"github.com/gorilla/mux"
)

func getPostId() int {
	postID++
	return postID
}

type Post struct {
	Id        int
	Title     string `schema:"title,title2example" json:"title" xml:"title`
	Text      string `schema:"text" json:"text" xml:"text`
	Author    string `schema:"author" json:"author" xml:"author`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AllPostsStruct struct {
	Title string
	Data  map[int]*Post
}

type OnePostsStruct struct {
	Title string
	Data  Post
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

	renderTemplate(w, "onePost", post)

}

func renderTemplate(w http.ResponseWriter, tmplName string, data interface{}) {
	templates := template.Must(template.ParseFiles("./static/tmpl_1page.html", "./static/tmpl.html",
		"./static/tmpl_allPosts.html", "./static/tmpl_edit.html", "./static/tmpl_create.html")) //Template caching

	err := templates.ExecuteTemplate(w, tmplName, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
	}
}

//Get
//3. Создайте роут и шаблон для создания материала
func createPostHandlerGet(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "CreatePosts", struct{ Title string }{"Новый пост"})
}

//POST
//3. Создайте роут и шаблон для создания материала
func createPostHandlerPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err) // Handle error
	}

	newID := getPostId()
	post := Post{
		Id: newID,
	}
	// r.PostForm is a map of our POST form values
	err = decoder.Decode(&post, r.PostForm)
	if err != nil {
		log.Println(err) // Handle error
	}
	fmt.Println(post)

	//Update BD
	insertPost(post)
	
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

//3. Создайте роут и шаблон для редактированияматериала.
func updatePostHandleGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(os.Stdout, r.Form)
	if r.Method == "GET" {
		fmt.Fprintf(os.Stdout, "GET, %q", html.EscapeString(r.URL.Path))
	} else if r.Method == "POST" {
		fmt.Fprintf(os.Stdout, "POST, %q", html.EscapeString(r.URL.Path))
	} else if r.Method == "PUT" {
		fmt.Fprintf(os.Stdout, "PUT, %q", html.EscapeString(r.URL.Path))
	} else {
		fmt.Fprintf(os.Stdout, "Invalid request method.", 405)
	}

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
	renderTemplate(w, "EditPost", post)
}

//PUT
//Проверяла через свой браузер: работает только c POST.
//C Postman PUT работает:x-www-form-urlncoded заполняем ключи author, text, title и значения  и отправляем POST
func updatePostHandlePut(w http.ResponseWriter, r *http.Request) {
	//renderTemplate(w, "AllPosts", posts)
	vars := mux.Vars(r)
	postIDRaw, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(postIDRaw)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = r.ParseForm()
	if err != nil {
		log.Println(err) // Handle error
	}

	post := Post{Id: id}

	// r.PostForm is a map of our POST form values
	err = decoder.Decode(&post, r.PostForm)
	if err != nil {
		log.Println(err) // Handle error
	}
	fmt.Println(post)
	
	//posts[id] = &post
	updateRow(id, post)
	// fmt.Println(posts[id])
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

//Delete post
func deletePostHandlerPost(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	postIDRaw, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(postIDRaw)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = r.ParseForm()
	if err != nil {
		log.Println(err) // Handle error
	}

	post := Post{Id: id}

	// r.PostForm is a map of our POST form values
	err = decoder.Decode(&post, r.PostForm)
	if err != nil {
		log.Println(err) // Handle error
	}
	fmt.Println(post)
	
	deleteRow(id) 
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

//GET
//1. роут и шаблон для отображения всех постов в блоге.
func listPostHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "AllPosts", AllPostsStruct{Title: "Список всех постов", Data: posts})
}
