package main

import (
	"./handler"
	"github.com/teambition/gear"
)

func main() {
	app := gear.New()
	router := gear.NewRouter()

	router.Get("/blog/list", handler.List)
	router.Get("/blog/detail/:title", handler.Detail)

	app.Use(func(ctx *gear.Context) error {
		return nil
	})

	app.UseHandler(router)

	app.Listen(":8080")
}
