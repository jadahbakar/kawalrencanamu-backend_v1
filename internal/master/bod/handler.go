package bod

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type bodHandler struct {
	bodService BodService
}

func NewBodHandler(f *fiber.App, bs BodService) {
	handler := &bodHandler{bodService: bs}
	f.Get("/bod/all", handler.GetAll)
	f.Get("/bod/:id", handler.GetById)

}

func (bh *bodHandler) GetAll(c *fiber.Ctx) error {
	return c.SendString("GetAll...!")
}

func (bh *bodHandler) GetById(c *fiber.Ctx) error {
	param := c.Params("id")
	result := fmt.Sprintf("GetById : %s", param)
	return c.SendString(result)
}
