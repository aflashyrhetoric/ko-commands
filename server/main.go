package main

import (
	"ko-commands/server/router"
	"ko-commands/server/routes"
)

func main() {
	r := router.New()
	routes.Register(r)
	r.Serve(":8080")
}
