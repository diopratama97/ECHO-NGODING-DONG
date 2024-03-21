package main

import route "test-echo/routes"

func main() {
	e := route.Init()

	e.Logger.Fatal(e.Start(":80"))
}
