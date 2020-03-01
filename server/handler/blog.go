package handler

import (
	"../lib"
	"bufio"
	"fmt"
	"github.com/teambition/gear"
	"io/ioutil"
	"os"
)

const TODO = "/home/eaves/TODO.md"

func List(ctx *gear.Context) error {

	// 服务器不分页，一共没多少文章，前端拿目录分页
	return ctx.JSON(200, lib.List)
}

func Detail(ctx *gear.Context) error {

	title := ctx.Param("title")

	b, err := ioutil.ReadFile(lib.RootConf.Root + "/" + title)
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
