package handler

import (
	"cv_backend/app/service"
	"cv_backend/model"
	"cv_backend/utils"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

// POST /api/register
func (h *UserHandler) Register(c *fiber.Ctx) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if err := h.service.Register(&user); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.SuccessResponse(c, "User registered successfully", nil)
}

// POST /api/login
func (h *UserHandler) Login(c *fiber.Ctx) error {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	user, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, err.Error())
	}

	role := "user"
	if user.Email == "admin@site.com" {
		role = "admin"
	}

	// Token üret
	token, err := utils.GenerateJWT(uint(user.ID), user.Email, role)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Token creation failed")
	}

	return utils.SuccessResponse(c, "Login successful", fiber.Map{
		"token": token,
		"user":  user,
		"role":  role,
	})
}

func (h *UserHandler) GetProfile(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"message": "Profil bilgisi burada gösterilecek.",
		"user_id": c.Locals("user_id"),
		"email":   c.Locals("email"),
		"role":    c.Locals("role"),
	})
}
