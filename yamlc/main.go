package main

import (
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

func main() {
	var c conf
	conf := c.getConf()
	fmt.Println(conf.Host)
}

//profile variables
type conf struct {
	Host   string `yaml:"host"`
	User   string `yaml:"user"`
	Pwd    string `yaml:"pwd"`
	Dbname string `yaml:"dbname"`
}

func (c *conf) getConf() *conf {
	execpath, err := os.Executable() // 获得程序路径
	if err != nil {
		fmt.Println(err.Error())
	}
	configfile := filepath.Join(filepath.Dir(execpath), "../conf.yml")
	yamlFile, err := ioutil.ReadFile(configfile)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}
