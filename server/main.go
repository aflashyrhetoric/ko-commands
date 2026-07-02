package main

import (
	"ko-commands/server/router"
	"ko-commands/server/routes"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5050"
	}
	r := router.New()
	routes.Register(r)
	r.Serve(":" + port)
}
