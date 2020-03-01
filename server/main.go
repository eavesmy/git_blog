package main

import (
	"./handler"
	"./lib"
	"github.com/teambition/gear"
)

const Root = "/data/workspace/blogs"
const Update = 60

func init() {
	lib.InitConf(Root, Update)
}

func main() {
	app := gear.New()
	router := gear.NewRouter()

	router.Get("/blog/list", handler.List)
	router.Get("/blog/detail/:title", handler.Detail)
	router.Get("/blog/index", handler.Index)

	app.Use(func(ctx *gear.Context) error {
		ctx.SetHeader("Access-Control-Allow-Origin", "*")
		return nil
	})

	app.UseHandler(router)

	app.Listen(":9093") // 9093
}
