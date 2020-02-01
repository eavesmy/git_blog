package handler

import (
	"github.com/teambition/gear"
	"io/ioutil"
	"path"
	"time"
)

var Root = "/data/workspace/blogs"

type Blog struct {
	Filename   string
	LastModify time.Time
}

func List(ctx *gear.Context) error {

	dir, err := ioutil.ReadDir(Root)
	if err != nil {
		return err
	}

	list := []Blog{}

	for _, fi := range dir {
		name := fi.Name()
		ext := path.Ext(name)
		if ext == ".md" {
			list = append(list, Blog{Filename: name, LastModify: fi.ModTime()})
		}
	}

	return ctx.JSON(200, list)
}

func Detail(ctx *gear.Context) error {

	title := ctx.Param("title")

	b, err := ioutil.ReadFile(Root + "/" + title)
	if err != nil {
		return ctx.HTML(500, "Something error")
	}

	return ctx.HTML(200, string(b))
}
