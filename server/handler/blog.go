package handler

import (
	"bufio"
	"fmt"
	gtype "github.com/eavesmy/golang-lib/type"
	"github.com/teambition/gear"
	"io/ioutil"
	"os"
	"path"
	"time"
)

const TODO = "/home/eaves/TODO.md"

var Root = "/data/workspace/blogs"

type Blog struct {
	Filename   string
	LastModify time.Time
}

func List(ctx *gear.Context) error {

	// 添加分页

	dir, err := ioutil.ReadDir(Root)
	if err != nil {
		return err
	}

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

func Index(ctx *gear.Context) error {

	fi, _ := os.Open(TODO)
	defer fi.Close()

	reader := bufio.NewReader(fi)
	line, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}

	return ctx.HTML(200, line)
}
