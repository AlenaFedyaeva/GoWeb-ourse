package main

import (
	"GoWebCourse/homework7/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary PageGetPost
// @Description render one post with id
// @Produce  text/html
// @Param id query string false "task list id"
// @Router /post/:id [get]
func getPostHandlerID(c *gin.Context) {
	postIDRaw := c.Param("id")

	postID, err := strconv.Atoi(postIDRaw)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	post, ok := db.Posts[postID]
	if !ok {
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}

	renderTemplate(c, "onePost", post)
}

func renderTemplate(c *gin.Context, tmplName string, data interface{}) {
	c.HTML(http.StatusOK, tmplName, data)
}

// //Get
// //3. Создайте роут и шаблон для создания материала
func createPostHandlerGet(c *gin.Context) {
	renderTemplate(c, "CreatePosts", struct{ Title string }{"Новый пост"})
}

// //POST
// //3. Создайте роут и шаблон для создания материала
func createPostHandlerPost(c *gin.Context) {
	c.Request.ParseForm()

	newID := db.GetPostId()
	post := db.Post{
		Id:     newID,
		Title:  c.Request.FormValue("title"),
		Text:   c.Request.FormValue("text"),
		Author: c.Request.FormValue("author"),
	}
	// fmt.Println(post)

	//Update BD
	mongo.InsertPost(post)

	c.Redirect(http.StatusSeeOther, "/")
}

// //3. Создайте роут и шаблон для редактированияматериала.
func updatePostHandleGet(c *gin.Context) {
	postIDRaw := c.Param("id")
	if postIDRaw == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "empty id value"})
		return
	}

	postID, err := strconv.Atoi(postIDRaw)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	post := db.Posts[postID]

	renderTemplate(c, "EditPost", post)
}

//PUT
//C Postman PUT работает:x-www-form-urlncoded заполняем ключи author, text, title и значения  и отправляем POST
func updatePostHandlePut(c *gin.Context) {
	postIDRaw := c.Param("id")
	if postIDRaw == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "empty id value"})
		return
	}

	postID, err := strconv.Atoi(postIDRaw)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	post := db.Post{
		Title:  c.Request.FormValue("title"),
		Text:   c.Request.FormValue("text"),
		Author: c.Request.FormValue("author"),
	}
	mongo.UpdateRow(postID, post)

	c.Redirect(http.StatusSeeOther, "/")
}

//Delete post
func deletePostHandlerPost(c *gin.Context)  {
	postIDRaw := c.Param("id")
	if postIDRaw == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "empty id value"})
		return
	}

	postID, err := strconv.Atoi(postIDRaw)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	mongo.DeleteRow(postID)
	c.Redirect(http.StatusSeeOther, "/")
}

// //GET
// //1. роут и шаблон для отображения всех постов в блоге.
func listPostHandler(c *gin.Context) {
	renderTemplate(c, "AllPosts", db.AllPostsStruct{Title: "Список всех постов", Data: db.Posts})
}
