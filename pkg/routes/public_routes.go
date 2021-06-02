package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jadahbakar/kawalrencanamu-backend/internal/master/entities"
	master_interfaces_http "github.com/jadahbakar/kawalrencanamu-backend/internal/master/handler/http"
	master_repository "github.com/jadahbakar/kawalrencanamu-backend/internal/master/repository"
	"github.com/jadahbakar/kawalrencanamu-backend/internal/me"
	"github.com/jadahbakar/kawalrencanamu-backend/pkg/config"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App, conf config.Config, pgxPool *pgxpool.Pool) {
	// Create routes group.
	router := a.Group(fmt.Sprintf("%s%s", conf.ApiURLGroup, conf.ApiURLVersion))
	me.AddRoutes(router)

	//--- Master
	// master.AddRoutes(router)
	// masterRepo := shop_infra_product.NewMemoryRepository()

	// master_interfaces_http.NewMasterHandler(router)
	masterRepo := master_repository.NewMasterRepository(pgxPool)
	master_interfaces_http.AddRoutes(router, entities.AssesmentEnvironmentEntity)

	//-----------

}
