package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jadahbakar/kawalrencanamu-backend/internal/me"
	"github.com/jadahbakar/kawalrencanamu-backend/pkg/config"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App, conf config.Config) {
	// Create routes group.
	router := a.Group(fmt.Sprintf("%s%s", conf.ApiURLGroup, conf.ApiURLVersion))
	me.AddRoutes(router)
	// Routes for GET method:
	// router.Get("/me", me.GetMe)

	// route.Get("/books", controllers.GetBooks)              // get list of all books
	// route.Get("/book/:id", controllers.GetBook)            // get one book by ID
	// route.Get("/token/new", controllers.GetNewAccessToken) // create a new access tokens
}
