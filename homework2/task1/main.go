package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type FindInfo struct {
	Search    string   `json:"search" xml:"search"`
	Sites     []string   `json:"sites" xml:"sites"`
}

// Поиск строки на веб странице
func findStrInURL(subsrt string, url string) bool {
	rez := false
	b, body := readURL(url)

	if b {
		//fmt.Println(body)
		rez = find(subsrt, body)
	}
	return rez
}

// Поиск подстроки в строке (* переписать функцию на поиск в массиве)
func find(subsrt string, str string) bool {
	res := strings.Contains(str, subsrt)
	return res
}

//Чтение body в string
func readURL(url string) (bool, string) {
	var rezStr string
	var b bool
	httpCli := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest(http.MethodOptions, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := httpCli.Do(req)
	if err != nil {
		fmt.Println(err)
		return b, rezStr
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		//log.Fatal(errors.New("readURL: status code not StatusOK "))
		return b, rezStr
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return b, rezStr
	}
	return true, string(bodyBytes)
}
// Поиск строки в списке веб страниц
func findInURLArray(arr []string, substr string) string{
	rez:="Not found"
	for _, url := range arr {
		b := findStrInURL(substr, url)
		if b {
			fmt.Printf("Find in page :  %s", url)
			rez=url
		}
	}
	return rez
}

func main() {

	router := http.NewServeMux()
	router.HandleFunc("/setInfo", SetInfoHandler)
	router.HandleFunc("/getInfo", getInfoHandler)

	log.Println("Starting server at :8091")
	log.Fatal(http.ListenAndServe(":8091", router))

	
	fmt.Println("\nbye")
}

func getInfoHandler(wr http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		wr.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	contentTypeHeader := req.Header.Get("Content-Type")

	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
		wr.WriteHeader(http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	info := &FindInfo{}
	switch contentTypeHeader {
	case "application/xml":
		if err = xml.Unmarshal(data, info); err != nil {
			log.Println(err)
			wr.WriteHeader(http.StatusInternalServerError)
			return
		}
	default:
		if err = json.Unmarshal(data, info); err != nil {
			log.Println(err)
			wr.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	log.Printf("Search: %s\nSites: %v\n",
		info.Search, info.Sites,
	)
	str:=findInURLArray(info.Sites,info.Search)
	fmt.Fprintf(wr,"Нашли %s на сайте %s", info.Search, str)
	wr.WriteHeader(http.StatusOK)
}


func SetInfoHandler(wr http.ResponseWriter, req *http.Request)  {
	acceptHeader := req.Header.Get("Accept")

	info:=&FindInfo{
		Search: "Go is an open source programming",
		Sites: []string{
			"https://mail.ru",
			"https://golang.org/",
			"https://google.com",
		},
	}
	var respBody []byte
	var err error
	switch acceptHeader {
	case "application/xml":
		respBody, err = xml.Marshal(info)
		if err != nil {
			log.Println(err)
			wr.WriteHeader(http.StatusInternalServerError)
			return
		}
		wr.Header().Set("Content-Type", "application/xml")
	default:
		respBody, err = json.Marshal(info)
		if err != nil {
			log.Println(err)
			wr.WriteHeader(http.StatusInternalServerError)
			return
		}
		wr.Header().Set("Content-Type", "application/json")
	}
	wr.Write(respBody)
}