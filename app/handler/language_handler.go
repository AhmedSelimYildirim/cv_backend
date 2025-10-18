package handler

import (
	"cv_backend/app/service"
	"cv_backend/viewmodel"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type LanguageHandler struct {
	Service *service.LanguageService
}

func NewLanguageHandler(s *service.LanguageService) *LanguageHandler {
	return &LanguageHandler{Service: s}
}

func (h *LanguageHandler) GetAllLanguages(c *fiber.Ctx) error {
	languages, err := h.Service.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	var dtos []viewmodel.LanguageDTO
	for _, l := range languages {
		dtos = append(dtos, *viewmodel.ToLanguageDTO(&l))
	}

	return c.JSON(dtos)
}

func (h *LanguageHandler) GetLanguageByID(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	language, err := h.Service.GetByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Language not found"})
	}
	return c.JSON(viewmodel.ToLanguageDTO(language))
}
