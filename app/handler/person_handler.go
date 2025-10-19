package handler

import (
	"cv_backend/app/service"
	"cv_backend/model"
	"cv_backend/viewmodel"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type PersonHandler struct {
	Service *service.PersonService
}

func NewPersonHandler(s *service.PersonService) *PersonHandler {
	return &PersonHandler{Service: s}
}

func (h *PersonHandler) GetAllPersons(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	status := c.Query("status")

	persons, total, err := h.Service.GetPersonsPaginated(status, page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	var dtoList []viewmodel.PersonDTO
	for _, p := range persons {
		dtoList = append(dtoList, *viewmodel.ToPersonDTO(&p))
	}

	return c.JSON(fiber.Map{
		"page":   page,
		"limit":  limit,
		"total":  total,
		"data":   dtoList,
		"status": status,
	})
}

func (h *PersonHandler) GetPersonByID(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	person, err := h.Service.GetPersonByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if person == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Person not found"})
	}
	return c.JSON(viewmodel.ToPersonDTO(person))
}

func (h *PersonHandler) CreatePerson(c *fiber.Ctx) error {
	var person model.Person
	if err := c.BodyParser(&person); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	createdPerson, err := h.Service.CreatePerson(&person)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(viewmodel.ToPersonDTO(createdPerson))
}

func (h *PersonHandler) DeletePerson(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)

	person, err := h.Service.GetPersonByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if person == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Person not found"})
	}

	if err := h.Service.DeletePerson(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *PersonHandler) UpdatePersonStatus(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var body struct {
		StatusType                string `json:"status_type"`
		ReasonForRejection        string `json:"reason_for_rejection"`
		ReasonForRejectionSummary string `json:"reason_for_rejection_summary"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	person, err := h.Service.GetPersonByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if person == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Person not found"})
	}

	switch body.StatusType {
	case "beklemede":
		person.StatusType = model.PersonDurumBeklemede
	case "onaylandi":
		person.StatusType = model.PersonDurumOnaylandi
	case "reddedildi":
		person.StatusType = model.PersonDurumReddedildi
	case "ilgileniliyor":
		person.StatusType = model.PersonDurumIlgileniliyor
	default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid status_type"})
	}

	person.ReasonForRejection = body.ReasonForRejection
	person.ReasonForRejectionSummary = body.ReasonForRejectionSummary

	if err := h.Service.UpdatePersonStatus(person); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(viewmodel.ToPersonDTO(person))
}
