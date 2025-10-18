package main

import (
	"cv_backend/app/handler"

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

	api.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Pong! ✅ Server & DB are working."})
	})
	
	api.Post("/persons", personHandler.CreatePerson)

	// Person endpoints
	api.Get("/persons", personHandler.GetAllPersons)
	api.Get("/persons/:id", personHandler.GetPersonByID)
	api.Put("/persons/:id/status", personHandler.UpdatePersonStatus)
	api.Delete("/persons/:id", personHandler.DeletePerson)

	// Language / Position / Reference endpoints (public)
	api.Get("/languages", languageHandler.GetAllLanguages)
	api.Get("/languages/:id", languageHandler.GetLanguageByID)
	api.Delete("/languages/:id", languageHandler.DeleteLanguage)

	api.Get("/positions", positionHandler.GetAllPositions)
	api.Get("/positions/:id", positionHandler.GetPositionByID)
	api.Delete("/positions/:id", positionHandler.DeletePosition)

	api.Get("/references", referenceHandler.GetAllReferences)
	api.Get("/references/:id", referenceHandler.GetReferenceByID)
	api.Delete("/references/:id", referenceHandler.DeleteReference)
}
