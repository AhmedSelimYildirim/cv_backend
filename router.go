package main

import (
	"cv_backend/app/handler"
	"cv_backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(
	app *fiber.App,
	userHandler *handler.UserHandler,
	personHandler *handler.PersonHandler,
	positionHandler *handler.PositionHandler,
	referenceHandler *handler.ReferenceHandler,
	languageHandler *handler.LanguageHandler,
) {
	api := app.Group("/api")

	// Test endpoint
	api.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Pong! ✅ Server & DB are working."})
	})

	// Public → CV doldurabilir veya kayıt/login olabilir
	api.Post("/register", userHandler.Register)
	api.Post("/login", userHandler.Login)

	// Protected → login olmuş herkes erişebilir
	auth := api.Group("/auth", middleware.JWTMiddleware())
	auth.Get("/me", userHandler.GetProfile)

	// Başvuruları yönetmek için role kontrolü yok
	auth.Get("/persons", personHandler.GetAllPersons)
	auth.Put("/persons/:id/status", personHandler.UpdatePersonStatus)
	auth.Get("/persons/:id", personHandler.GetPersonByID)
	auth.Delete("/persons/:id", personHandler.DeletePerson)

	// Language / Position / Reference → login olmuş herkes ekleyebilir
	auth.Get("/languages", languageHandler.GetAllLanguages)
	auth.Get("/positions", positionHandler.GetAllPositions)
	auth.Get("/references", referenceHandler.GetAllReferences)
}
