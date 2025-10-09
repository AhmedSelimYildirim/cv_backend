package handler

import (
	"cv_backend/app/service"
	"cv_backend/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type PositionHandler struct {
	service *service.PositionService
}

func NewPositionHandler(service *service.PositionService) *PositionHandler {
	return &PositionHandler{service: service}
}

func (h *PositionHandler) CreatePosition(c *fiber.Ctx) error {
	var position model.Position
	if err := c.BodyParser(&position); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.service.CreatePosition(&position); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(position)
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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if position == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Position not found"})
	}
	return c.JSON(position)
}

func (h *PositionHandler) UpdatePosition(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	var position model.Position
	if err := c.BodyParser(&position); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	position.ID = id

	if err := h.service.UpdatePosition(&position); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
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
