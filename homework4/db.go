package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func insertPost(p Post) error {
	//1) insert into bd
	_, err := database.Exec("INSERT INTO posts(author, postText,title) VALUES (?,?,?)", p.Author, p.Text, p.Title)

	if err != nil {
		log.Println(err)
		return err
	}
	//2) update map values
	updatePostsMap()
	return nil
}

func selectAll() (map[int]*Post, error) {
	res := make(map[int]*Post)
	rows, err := database.Query("SELECT id,author,title,postText,created_at,updated_at FROM posts.posts")

	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		post := Post{}
		err := rows.Scan(&post.Id, &post.Author, &post.Title, &post.Text, &post.CreatedAt, &post.UpdatedAt)

		if err != nil {
			log.Println(err)
			continue
		}
		res[post.Id] = &post
	}
	return res, nil
}

func selectPost(id int) (Post, error) {
	post := Post{}

	row := database.QueryRow(fmt.Sprintf("SELECT * FROM posts.posts where id = %d", id))
	err := row.Scan(&post.Id, &post.Author, &post.Title, &post.Text, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		return post, err
	}
	return post, err
}

func updatePostsMap() {
	postsNew, err := selectAll()
	if err != nil {
		fmt.Println(err)
	}
	posts = postsNew
}

func updateRow(id int, p Post) error {
	//1) update bd
	_, err := database.Exec("update posts set author=?,title=?,postText=?,updated_at=?) where id =?",
		p.Author, p.Title, p.Text, p.UpdatedAt, id)
	if err != nil {
		log.Println(err)
		return err
	}
	//2) update map values
	updatePostsMap()
	return nil
}

func deleteRow(id int) (int64, error) {
	//del row
	result, err := database.Exec("delete from posts.posts where id = ?", id)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	//2) update map values
	updatePostsMap()
	return result.RowsAffected()
}

func dbConnect() {
}

func tryBD(){
	sel1,err:=selectAll() 
	if err != nil {
		fmt.Println(err)
	}
	sel2,err:=selectPost(13)  
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sel1)
	fmt.Println(sel2)
	deleteRow(12)

	insertPost(Post{Text: "wwwww",Title: "insert",} )

	fmt.Println("post rez")
	fmt.Println(posts)
}