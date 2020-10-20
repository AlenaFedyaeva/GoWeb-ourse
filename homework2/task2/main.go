package main

import (
	"fmt"
	"log"
	"net/http"
)


func main() {
	router := http.NewServeMux()

	router.HandleFunc("/user", userHandler)
	router.HandleFunc("/", firstHandlerRoot)
	router.HandleFunc("/username", userNameHandler)

	log.Println("Starting server at :8094")
	log.Fatal(http.ListenAndServe(":8094", router))
}

// В корне выдаем cookie firstCookie
func firstHandlerRoot(wr http.ResponseWriter, req *http.Request) {
	cookie := &http.Cookie{
		Name:  "firstCookie",
		Value: "firstCookieValue",
	}
	http.SetCookie(wr, cookie)
	wr.Write([]byte("hello in root!. Get your firstcookie) "))
}

// Читаем cookie userName
func userNameHandler(wr http.ResponseWriter, req *http.Request) {
	wr.Write([]byte("hello! (secondHandle)"))
	cookie, err := req.Cookie("userName")
	if err != nil {
		fmt.Fprintf(wr, "Чтобы получить cookie go  http://localhost:8092/user?name=Ivan", req.URL.Query().Get("name"))
	}
	fmt.Fprintf(wr, "Hello, %s   Your cookie userName- %s !", req.URL.Query().Get("name"),cookie.Value)
}
// Выводим имя пользователя например http://localhost:8092/user?name=Ivan
// выдаем cookie userName
func userHandler(wr http.ResponseWriter, req *http.Request) {
	name:=req.URL.Query().Get("name")

	cookie := &http.Cookie{
		Name:  "userName",
		Value: name,
	}
	http.SetCookie(wr, cookie)
	wr.Write([]byte("Get your cookie! Example: http://localhost:8092/user?name=Ivan "))			
}
