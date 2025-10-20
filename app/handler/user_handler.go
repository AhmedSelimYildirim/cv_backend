package handler

import (
	"cv_backend/app/service"
	"cv_backend/model"
	"cv_backend/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var user struct {
		Name     string `json:"name"`
		Surname  string `json:"surname"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&user); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	u := &model.User{
		Name:     user.Name,
		Surname:  user.Surname,
		Email:    user.Email,
		Password: user.Password,
	}

	if err := h.service.Register(u); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.SuccessResponse(c, "User registered successfully", nil)
}

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

	token, err := utils.GenerateJWT(uint(user.ID), user.Email)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Token creation failed")
	}

	return utils.SuccessResponse(c, "Login successful", fiber.Map{
		"token": token,
		"user":  user,
	})
}

func (h *UserHandler) GetProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id")
	return utils.SuccessResponse(c, "Profile fetched successfully", fiber.Map{
		"user_id": userID,
		"email":   c.Locals("email"),
		"role":    c.Locals("role"),
	})
}

func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.service.GetAllUsers()
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	return utils.SuccessResponse(c, "All users fetched successfully", users)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	userID := c.Locals("user_id")
	var id int64
	switch v := userID.(type) {
	case float64:
		id = int64(v)
	case int:
		id = int64(v)
	case int64:
		id = v
	default:
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Invalid user")
	}

	var body struct {
		Name     string `json:"name"`
		Surname  string `json:"surname"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&body); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	user, err := h.service.GetByID(id)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	if user == nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, "User not found")
	}

	if body.Name != "" {
		user.Name = body.Name
	}
	if body.Surname != "" {
		user.Surname = body.Surname
	}
	if body.Email != "" {
		user.Email = body.Email
	}
	if body.Password != "" {
		hashed, err := utils.HashPassword(body.Password)
		if err != nil {
			return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Password hash failed")
		}
		user.Password = hashed
	}

	if err := h.service.UpdateUser(user); err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, "User updated successfully", user)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	userID := c.Locals("user_id")
	tokenUserID := int64(userID.(float64))
	paramID, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid ID")
	}

	if tokenUserID != paramID {
		return utils.ErrorResponse(c, fiber.StatusForbidden, "You can only delete your own account")
	}

	user, err := h.service.GetByID(paramID)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	if user == nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, "User not found")
	}

	if err := h.service.DeleteUser(paramID); err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.SendStatus(fiber.StatusNoContent)
}
