package bod

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

type BodHandler struct {
	bodService BodService
}

func NewBodHandler(f *fiber.App, bs BodService) {
	handler := &BodHandler{bodService: bs}
	f.Get("/bod/all", handler.GetAll)
	f.Get("/bod/:id", handler.GetById)
}

func (bh *BodHandler) GetAll(c *fiber.Ctx) error {
	listBod, err := bh.bodService.FindAll()
	if err != nil {
		zap.Error(err)
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": listBod,
	})
}

func (bh *BodHandler) GetById(c *fiber.Ctx) error {
	paramId := c.Params("id")
	id, err := strconv.Atoi(paramId)
	// if error in parsing string to int
	if err != nil {
		zap.Error(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse Id",
		})
	}
	data, err := bh.bodService.FindById(id)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": data,
	})
}
