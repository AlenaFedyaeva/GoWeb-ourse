package main

import (
	"log"
	"net/http"
	"text/template"
)

//GET
//2. роут и шаблон для просмотра конкретного поста в блоге.
func getPostHandlerID(w http.ResponseWriter, r *http.Request) {
	
	// renderTemplate(w, "onePost", post)

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

//GET
//1. роут и шаблон для отображения всех постов в блоге.
func listPostHandler(w http.ResponseWriter, r *http.Request) {
	// renderTemplate(w, "AllPosts", AllPostsStruct{Title: "Список всех постов", Data: posts})
}
