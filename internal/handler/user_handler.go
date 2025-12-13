package handler

import (
	"strconv"
	"time"

	"user-age-api/internal/models"
	"user-age-api/internal/service"

	"github.com/gofiber/fiber/v2"
)

// UserHandler holds user-related handlers
type UserHandler struct {
	userService *service.UserService
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// -------------------- CREATE USER --------------------
// POST /users
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {

	var req models.CreateUserRequest

	// Parse request body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "invalid request body"})
	}

	// Validate input
	if err := models.Validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": err.Error()})
	}

	// Convert DOB string to time.Time
	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "invalid date format"})
	}

	// Call service layer
	user, err := h.userService.CreateUser(c.Context(), req.Name, dob)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

// -------------------- GET USER BY ID --------------------
// GET /users/:id
func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {

	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "invalid user id"})
	}

	user, err := h.userService.GetUserByID(c.Context(), int32(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).
			JSON(fiber.Map{"error": "user not found"})
	}

	return c.JSON(user)
}

// -------------------- LIST USERS --------------------
// GET /users
func (h *UserHandler) ListUsers(c *fiber.Ctx) error {

	users, err := h.userService.ListUsers(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(users)
}

// -------------------- UPDATE USER --------------------
// PUT /users/:id
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {

	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "invalid user id"})
	}

	var req models.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "invalid request body"})
	}

	if err := models.Validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": err.Error()})
	}

	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "invalid date format"})
	}

	user, err := h.userService.UpdateUser(c.Context(), int32(id), req.Name, dob)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(user)
}

// -------------------- DELETE USER --------------------
// DELETE /users/:id
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {

	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "invalid user id"})
	}

	if err := h.userService.DeleteUser(c.Context(), int32(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
