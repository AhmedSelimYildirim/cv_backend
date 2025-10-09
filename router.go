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

	// Public
	api.Post("/register", userHandler.Register)
	api.Post("/login", userHandler.Login)

	// Protected
	auth := api.Group("/auth", middleware.JWTMiddleware())
	auth.Get("/me", userHandler.GetProfile)

	// Language
	auth.Post("/languages", languageHandler.CreateLanguage)
	auth.Get("/languages", languageHandler.GetAllLanguages)
	auth.Get("/languages/:id", languageHandler.GetLanguageByID)
	auth.Put("/languages/:id", languageHandler.UpdateLanguage)
	auth.Delete("/languages/:id", languageHandler.DeleteLanguage)

	// Person
	auth.Post("/persons", personHandler.CreatePerson)
	auth.Get("/persons", middleware.RequireRole("admin"), personHandler.GetAllPersons)
	auth.Get("/persons/:id", personHandler.GetPersonByID)
	auth.Put("/persons/:id", personHandler.UpdatePerson)
	auth.Delete("/persons/:id", personHandler.DeletePerson)

	// Position
	auth.Post("/positions", positionHandler.CreatePosition)
	auth.Get("/positions", positionHandler.GetAllPositions)
	auth.Get("/positions/:id", positionHandler.GetPositionByID)
	auth.Put("/positions/:id", positionHandler.UpdatePosition)
	auth.Delete("/positions/:id", positionHandler.DeletePosition)

	// Reference
	auth.Post("/references", referenceHandler.CreateReference)
	auth.Get("/references", referenceHandler.GetAllReferences)
	auth.Get("/references/:id", referenceHandler.GetReferenceByID)
	auth.Put("/references/:id", referenceHandler.UpdateReference)
	auth.Delete("/references/:id", referenceHandler.DeleteReference)
}
