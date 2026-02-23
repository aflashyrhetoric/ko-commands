package routes

import (
	"ko-commands/server/controllers"
	"ko-commands/server/router"
)

func Register(r *router.Router) {
	r.Get("/recycling", controllers.Recycling)
}
