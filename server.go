package main

import (
	"test-echo/db"
	route "test-echo/routes"
)

func main() {
	db.Init()
	e := route.Init()

	e.Logger.Fatal(e.Start(":80"))
}
