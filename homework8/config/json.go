package config

import (
	"GoWebCourse/homework8/db"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type FileConfig struct {
	DB             db.DBInfo `yaml: db`
	CollectionName string `yaml: collectionName`
	Port           string `yaml: port`
	LogName 	string `yaml: logName`
}

var DefaultConf FileConfig=FileConfig{
	DB: db.DBInfo{URI: "mongodb://localhost:27017",
			Name: "posts",
		},
		CollectionName: "posts",
		Port: "9090",
		LogName: "log.txt",
}

func (conf *FileConfig) WriteConfigJson(fname string) {
	file, _ := os.OpenFile(fname,  os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.Encode(conf)
	fmt.Println("write in ", fname,*conf,file)

}

func (conf *FileConfig) ReadConfigJson(fname string) error {
	var tmp FileConfig
	data, err := ioutil.ReadFile(fname)
	if err!=nil{
		log.Println(err,data, &data)
		return err
	}
	err=json.Unmarshal(data, &tmp)
	if err!=nil{
		log.Println(err,data, &data,tmp)
		return err
	}
	*conf=tmp
	fmt.Println("read from ", fname,conf)
	return nil
}

func WriteDefaultConfigJson(fname string){
	DefaultConf.WriteConfigJson(fname)
}

// func IfFileExist(fname) bool{
// 	// if _, err := os.Stat("/path/to/whatever"); os.IsNotExist(err) {
// 		// path/to/whatever does not exist
// 	// }
// }