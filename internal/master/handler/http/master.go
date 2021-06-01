package http

import (
	"github.com/gofiber/fiber"
	"github.com/jadahbakar/kawalrencanamu-backend/internal/master/entities"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}
type MasterHandler struct {
	MasterEntity entities.AssesmentEnvironmentEntity
}

func NewMasterHandler(f fiber.Router, us entities.AssesmentEnvironmentEntity) {
	handler := &MasterHandler{
		MasterEntity: us,
	}
	f.Get("/master/", handler.Fetch)

}

func (m *MasterHandler) Fetch(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Fetch"})
}
