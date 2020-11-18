package db

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @Summary PageGetPost
// @Description render one post with id
// @Produce  text/html
// @Param id query string false "task list id"
// @Router /post/:id [get]
func (cnt *Controller) getPostHandlerID(c *gin.Context) {
	postIDRaw := c.Param("id")

	postID, err := strconv.Atoi(postIDRaw)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	post, ok := Posts[postID]
	if !ok {
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}

	renderTemplate(c, "onePost", post)
}

func renderTemplate(c *gin.Context, tmplName string, data interface{}) {
	c.HTML(http.StatusOK, tmplName, data)
}



// @Summary PageNewPost
// @Description show page for new post
// @Produce  text/html
// @Router /create [get]
func (cnt *Controller) createPostHandlerGet(c *gin.Context) {
	renderTemplate(c, "CreatePosts", struct{ Title string }{"Новый пост"})
}

// @Summary FormPageNewPost
// @Description update posts & redirect to /
// @Produce  text/html
// @Router /create [post]
func (cnt *Controller) createPostHandlerPost(c *gin.Context) {
	c.Request.ParseForm()

	newID := GetPostId()
	post := Post{
		Id:     newID,
		Title:  c.Request.FormValue("title"),
		Text:   c.Request.FormValue("text"),
		Author: c.Request.FormValue("author"),
	}
	// fmt.Println(post)

	//Update BD
	cnt.ControllerDB.InsertPost(post)

	c.Redirect(http.StatusSeeOther, "/")
}

// @Summary PageUpdatePost
// @Description show page for new post
// @Produce  text/html
// @Router /edit/:id [get]
func (cnt *Controller) updatePostHandleGet(c *gin.Context) {
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
	post := Posts[postID]

	renderTemplate(c, "EditPost", post)
}

// C Postman PUT работает:x-www-form-urlncoded заполняем ключи author, text, title и значения  и отправляем POST

// @Summary FormUpdatePost
// @Description show page for new post
// @Produce  text/html
// @Router /edit/:id [put]
func (cnt *Controller) updatePostHandlePut(c *gin.Context) {
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

	post := Post{
		Title:  c.Request.FormValue("title"),
		Text:   c.Request.FormValue("text"),
		Author: c.Request.FormValue("author"),
	}
	cnt.ControllerDB.UpdateRow(postID, post)

	c.Redirect(http.StatusSeeOther, "/")
}

// @Summary ButtonDeletePost
// @Description delete post and redirect to / 
// @Produce  text/html
// @Router /delete/:id [post]
func (cnt *Controller) deletePostHandlerPost(c *gin.Context)  {
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

	cnt.ControllerDB.DeleteRow(postID)
	c.Redirect(http.StatusSeeOther, "/")
}
// @Summary PagePostList
// @Description show all posts in blog
// @Produce  text/html
// @Router / [get]
func (cnt *Controller) listPostHandler(c *gin.Context) {
	renderTemplate(c, "AllPosts", AllPostsStruct{Title: "Список всех постов", Data: Posts})
}



func SetupServer(cnt *Controller) *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("static/*")

	//Шаблон со списком всех постов / короткие без Text
	router.GET("/", cnt.listPostHandler)
	router.POST("/delete/:id", cnt.deletePostHandlerPost)

	//Шаблон с текстовыми полями для задания  Title Text Author
	router.GET("/create", cnt.createPostHandlerGet)
	router.POST("/create", cnt.createPostHandlerPost)

	//Шаблон со страницей одного поста / полгого с отображением Text
	router.GET("/post/:id", cnt.getPostHandlerID)

	//Шаблон с текстовыми полями для обновления Title Text Author
	router.GET("/edit/:id", cnt.updatePostHandleGet)
	router.POST("/edit/:id", cnt.updatePostHandlePut)

	//Docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	

	return router
}