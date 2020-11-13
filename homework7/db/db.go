package db

type DBInfo struct{
	Name string 
}

type DBMongo{
	DBInfo
	Collection string
	*connect
}

type DB interface{
	SelectPost(id int) (Post, error)
}

func UseExample(){
	mongo:=&DBMongo{Name: "posts", Collection: "posts"}
	mongo. updatePostsMap()
}