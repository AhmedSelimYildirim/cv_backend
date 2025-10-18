package handler

import (
	"cv_backend/app/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type PositionHandler struct {
	service *service.PositionService
}

func NewPositionHandler(s *service.PositionService) *PositionHandler {
	return &PositionHandler{service: s}
}

func (h *PositionHandler) GetAllPositions(c *fiber.Ctx) error {
	positions, err := h.service.GetAllPositions()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(positions)
}

func (h *PositionHandler) GetPositionByID(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	position, err := h.service.GetPositionByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Position not found"})
	}
	return c.JSON(position)
}

func (h *PositionHandler) DeletePosition(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	if err := h.service.DeletePosition(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
