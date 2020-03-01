package lib

// dir.go 维护文件夹中的文件

import (
	"fmt"
	"io/ioutil"
	"path"
	"time"
)

type Blog struct {
	Filename   string
	LastModify time.Time
}

var List []Blog
var UpdateChan = make(chan bool, 1)

func init() {
	go loop()
}

func Update() {

	UpdateChan <- true

	time.AfterFunc(time.Duration(RootConf.Update)*time.Second, func() {
		Update()
	})
}

func initDir() {

	List = []Blog{}

	dir, err := ioutil.ReadDir(RootConf.Root)
	if err != nil {
		fmt.Println(err)
	}

	for _, fi := range dir {
		name := fi.Name()
		ext := path.Ext(name)
		if ext == ".md" {
			List = append(List, Blog{Filename: name, LastModify: fi.ModTime()})
		}
	}
}

func loop() {
	for {
		single := <-UpdateChan
		if single {
			initDir()
		}
	}
}
