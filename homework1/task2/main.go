package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os"
)

type PublicURLBody struct {
	Href      string `json:"href"`
	Method    string `json:"method"`
	Templated bool `json:"templated"`
}

func unmarshallPublicURLBody(bodyBytes []byte) string{

	//str:=string(bodyBytes)
	//fmt.Println(str)
	body := PublicURLBody{}

	if err := json.Unmarshal(bodyBytes, &body); err != nil {
		log.Panic(err)
	}

	//fmt.Println("\n----------------------")
	//fmt.Printf(body.Href)
	return body.Href
}

func createGetURL(publicLink string) string {
	return fmt.Sprintf("https://cloud-api.yandex.net/v1/disk/public/resources/download?public_key=%s", publicLink)
}

func getPublicURL(getURL string) (bool, string) {
	resp, err := http.Get(getURL)
	if err != nil {
		fmt.Println(err)
		return false, ""
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return false, ""
	}

	//fmt.Println(string(bodyBytes))
	//fmt.Println("\n\n\n")

	return true, unmarshallPublicURLBody(bodyBytes) 
}

func downloadFile(getURL string ) (int64,string){
	resp, err := http.Get(getURL)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	contentDisposition := resp.Header.Get("Content-Disposition")

	_, params, err := mime.ParseMediaType(contentDisposition)
	if err != nil {
		fmt.Println(err)
	}
	filename := params["filename"] // set to "foo.png"


	//fmt.Println("contentDisposition")
	//fmt.Println(filename)


	file, err := os.Create(filename)
    if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	size, err := io.Copy(file, resp.Body) 
	if err != nil {
		fmt.Println(err)
	}
	return size,filename
}

func main() {
	fmt.Println("task2")

	// 1)
	fmt.Println("Public link")
	publicLink := "https://yadi.sk/i/TmWwkH32foKjVA"
	getURL := createGetURL(publicLink)
	fmt.Println(getURL)
	// 2)
	fmt.Println("Download link")
	b, downloadLink := getPublicURL(getURL)
	if b {
		fmt.Println(downloadLink)
	}

	size,fname:=downloadFile(downloadLink)
	fmt.Println(size)
	fmt.Println(fname)
	fmt.Println("/nbye")
}
