package lib

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Conf struct {
	Root   string `json:"root"`
	Update int    `json:"update"`
}

var RootConf *Conf

func InitConf() {
	b, err := ioutil.ReadFile("./conf.json")
	if err != nil {
		log.Println(err)
		return
	}

	json.Unmarshal(b, &RootConf)

	InitDir()
}
