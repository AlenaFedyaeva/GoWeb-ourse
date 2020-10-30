package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)


func insert(p Post){

}


func selectAll() ([]Post, error){
	res := []Post{}
	rows, err := database.Query("SELECT * FROM posts.posts")
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		post := Post{}

		err := rows.Scan(&post.Id, &post.Author,&post.Title,&post.Text,&post.CreatedAt,&post.UpdatedAt)
		if err != nil {
			log.Println(err)
			continue
		}

		res = append(res, post)
	}

	return res, nil
}

func selectRow(id int) (Post,error){
	post:= Post{}
	
	row := database.QueryRow(fmt.Sprintf("SELECT * FROM posts.posts where id = %s", id))
	err := row.Scan(&post.Id, &post.Author,&post.Title,&post.Text,&post.CreatedAt,&post.UpdatedAt)
	if err != nil {
		return post, err
	}
	return post, err
}

func update(id int,p Post){

}

func dbConnect()  {
}