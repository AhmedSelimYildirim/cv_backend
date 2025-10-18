package handler

import (
	"cv_backend/app/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ReferenceHandler struct {
	service *service.ReferenceService
}

func NewReferenceHandler(s *service.ReferenceService) *ReferenceHandler {
	return &ReferenceHandler{service: s}
}

func (h *ReferenceHandler) GetAllReferences(c *fiber.Ctx) error {
	refs, err := h.service.GetAllReferences()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(refs)
}

func (h *ReferenceHandler) GetReferenceByID(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	ref, err := h.service.GetReferenceByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Reference not found"})
	}
	return c.JSON(ref)
}

func (h *ReferenceHandler) DeleteReference(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	if err := h.service.DeleteReference(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
