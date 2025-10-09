package handler

import (
	"cv_backend/app/service"
	"cv_backend/viewmodel"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type PersonHandler struct {
	Service *service.PersonService
}

func NewPersonHandler(service *service.PersonService) *PersonHandler {
	return &PersonHandler{Service: service}
}

func (h *PersonHandler) CreatePerson(c *fiber.Ctx) error {
	var dto viewmodel.PersonDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	person := dto.ToModel()
	if err := h.Service.CreatePerson(person); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(person)
}

func (h *PersonHandler) GetAllPersons(c *fiber.Ctx) error {
	persons, err := h.Service.GetAllPersons()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(persons)
}

func (h *PersonHandler) GetPersonByID(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	person, err := h.Service.GetPersonByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Person not found"})
	}
	return c.JSON(person)
}

func (h *PersonHandler) UpdatePerson(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	var dto viewmodel.PersonDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	person := dto.ToModel()
	person.ID = id
	if err := h.Service.UpdatePerson(person); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(person)
}

func (h *PersonHandler) DeletePerson(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	if err := h.Service.DeletePerson(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
