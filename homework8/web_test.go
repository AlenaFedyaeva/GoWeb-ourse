package main

import (
	"GoWebCourse/homework8/db"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Client struct {
	*http.Client
	URL string
}

type mockDB struct {
}


func (dbm *mockDB) SelectAll() (map[int]*db.Post, error) {
	rez:= db.PostsDefault
	 return rez,nil
}
func (dbm *mockDB) SelectPost(id int) (db.Post, error){
	rez:=db.PostsDefault[2]
	return *rez, nil
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
	// cli.Get("/?name=Test")
	resp1, _ := cli.Get(fmt.Sprintf("%s/", cli.URL))
	if resp1.StatusCode != http.StatusOK  {
	   t.Errorf("/ endpoint returned wrong response; expected %d; actual %d", http.StatusOK , resp1.StatusCode)
	}
	resp2, _ := cli.Get(fmt.Sprintf("%s/create", cli.URL))
	if resp1.StatusCode != http.StatusOK  {
	   t.Errorf("/create endpoint returned wrong response; expected %d; actual %d", http.StatusOK , resp2.StatusCode)
	}
	resp3, _ := cli.Get(fmt.Sprintf("%s/edit/:%s", cli.URL,"2"))
	if resp1.StatusCode != http.StatusOK  {
	   t.Errorf("/edit/:id endpoint returned wrong response; expected %d; actual %d", http.StatusOK , resp3.StatusCode)
	}
	resp4, _ := cli.Get(fmt.Sprintf("%s/post/:%s", cli.URL,"2"))
	if resp1.StatusCode != http.StatusOK  {
	   t.Errorf("/post/:id endpoint returned wrong response; expected %d; actual %d", http.StatusOK , resp4.StatusCode)
	}
	// json.NewDecoder(resp1.Body)
 }