package main

import (
	"GoWebCourse/homework7/db"
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


func (dbm *mockDB) SelectAll() (map[int]*db.Post, error) {
	rez:= map[int]*db.Post{
		1: &(db.Post{
			Id:        1,
			Title:     "some text1",
			Text:      "some text1",
			Author:    "some author1",
			CreatedAt: time.Now().Add(-time.Hour),
		}),
		2: &(db.Post{
			Id:        2,
			Title:     "some text2",
			Text:      "some text2",
			Author:    "some author2",
			CreatedAt: time.Now().Add(-time.Hour),
		}),
		3: &(db.Post{
			Id:        3,
			Title:     "some text3",
			Text:      "some text3",
			Author:    "some author3",
			CreatedAt: time.Now().Add(-time.Hour),
		}),
	}
	 return rez,nil
}
func (dbm *mockDB) SelectPost(id int) (db.Post, error){
	rez:=db.Post{
		Id:        3,
		Title:     "some text3",
		Text:      "some text3",
		Author:    "some author3",
		CreatedAt: time.Now().Add(-time.Hour),
	}
	return rez, nil
}
func (dbm *mockDB) InsertPost(post db.Post) (int,error){
	return 0,nil
}
func (dbm *mockDB) UpdateRow(id int, p db.Post) error {
	return nil
}
func (dbm *mockDB) DeleteRow(id int) error{
	return nil
}
func (dbm *mockDB) UpdatePostsMap(){
}
func (dbm *mockDB) DBInit() error{
	return nil
}
func (dbm *mockDB) Disconnect(){
}


func TestListPostRequest(t *testing.T) {
	testController := &db.Controller{
		ControllerDB: &mockDB{},
	}
	srv := httptest.NewServer(db.SetupServer(testController))
	defer srv.Close()

	cli := &Client{
		Client: srv.Client(),
		URL: srv.URL,
	}
	resp, _ := cli.Get(fmt.Sprintf("%s/", cli.URL))
	if resp.StatusCode != 200 {
	   t.Errorf("/ endpoint returned wrong response; expected %d; actual %d", 200, resp.StatusCode)
	}
  }