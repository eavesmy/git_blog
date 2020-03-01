package lib

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"
)

// dir.go 维护文件夹中的文件
type Blog struct {
	Filename   string
	LastModify time.Time
	UnixT      int64
}

var Map = map[String]Blog{}

func InitDir() {

	dir, err := ioutil.ReadDir(RootConf.Root)
	if err != nil {
		fmt.Println(err)
	}

	for i, fi := range dir {
		name := fi.Name()
		fmt.Println(i, name)
		ext := path.Ext(name)
		if ext == ".md" {
			// 做个桶排序，每个桶五个
		}
	}
}

func (b *Blogs) addBlog(fi *os.File) {
}
