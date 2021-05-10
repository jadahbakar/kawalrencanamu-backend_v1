package routes

import (
	"github.com/gofiber/fiber/v2"
	// "github.com/koddr/tutorial-go-fiber-rest-api/app/controllers"
	"github.com/jadahbakar/kawalrencanamu-backend/pkg/config"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App, conf config.Config) {
	// Create routes group.
	route := a.Group(conf.ApiURLGroup)

	route.Get("/me", me.)
	// Routes for GET method:
	// route.Get("/books", controllers.GetBooks)              // get list of all books
	// route.Get("/book/:id", controllers.GetBook)            // get one book by ID
	// route.Get("/token/new", controllers.GetNewAccessToken) // create a new access tokens
}
