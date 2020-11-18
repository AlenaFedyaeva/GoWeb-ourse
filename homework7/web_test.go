package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type Client struct {
	*http.Client
	URL string
}

type mockDB struct {
}


func (db *mockDB) SelectAll() (map[int]*Post, error) {
	rez:= map[int]*Post{
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
	 return rez,nil
}
func (db *mockDB) SelectPost(id int) (Post, error){
	rez:=Post{
		Id:        3,
		Title:     "some text3",
		Text:      "some text3",
		Author:    "some author3",
		CreatedAt: time.Now().Add(-time.Hour),
	}
	return rez, nil
}
func (db *mockDB) InsertPost(post Post) (int,error){
	return 0,nil
}
func (db *mockDB) UpdateRow(id int, p Post) error {
	return nil
}
func (db *mockDB) DeleteRow(id int) error{
	return nil
}
func (db *mockDB) UpdatePostsMap(){
}
func (db *mockDB) DBInit() error{
	return nil
}
func (db *mockDB) Disconnect(){
	return nil
}


func TestCreteListRequest(t *testing.T) {
	testController := &controller{
		db: &mockDB{},
	}
	srv := httptest.NewServer(setupServer(testController))
	defer srv.Close()

	cli := &Client{
		Client: srv.Client(),
		URL: srv.URL,
	}
	resp, _ := cli.Get(fmt.Sprintf("%s/lists", cli.URL))
	if resp.StatusCode != 200 {
	   t.Errorf("/lists endpoint returned wrong response; expected %d; actual %d", 200, resp.StatusCode)
	}
  }