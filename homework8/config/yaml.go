package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-yaml/yaml"
)

// Read config from yaml
func (conf *FileConfig) NewCongigYaml(fname string) (*FileConfig,error){
	data,err:=ioutil.ReadFile(fname)
	if err!=nil {
		return nil,err
	}
	var fc FileConfig
	err=yaml.Unmarshal(data, &fc)
	return &fc,err
}

func WriteDefaultConfigYaml(fname string){
	DefaultConf.WriteConfigJson(fname)
}

func (conf *FileConfig) WriteConfigYaml(fname string) {
	file, _ := os.OpenFile(fname,  os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer file.Close()
	encoder := yaml.NewEncoder(file)
	encoder.Encode(conf)
	fmt.Println("write in ", fname,*conf,file)

}
