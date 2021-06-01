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

	master.AddRoutes(router)
}
